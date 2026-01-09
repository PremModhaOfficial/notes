# NATS Micro Part 2: Internals & Patterns

Welcome to Part 2 of our journey into NATS Micro. In Part 1, we learned the basics of creating services, endpoints, and groups. Now, we'll go a layer deeper to understand how the framework operates under the hood. 

Understanding the internals will help you build more robust, scalable, and observable microservices. We'll explore the core `$SRV` protocol, the service lifecycle, and powerful patterns for production-ready applications.

## 1. The $SRV Protocol: The Heartbeat of Micro

At the core of NATS Micro is a simple yet powerful set of conventions built on top of standard NATS messaging. This is known as the `$SRV` protocol. It's not a separate protocol in the way TCP is to HTTP; it's just a structured way of using NATS subjects to manage services.

This protocol is what allows services to be discoverable, observable, and versioned, all without a central service registry.

### Three Levels of Service Subscriptions

When you add an endpoint (a "verb") to your service, the framework doesn't just create one subscription; it creates three. This allows for incredible flexibility in how you can call a service. Let's say we have a `billing` service with a `charge` verb.

1.  **`$SRV.charge`**: This is the most general subject. A request sent here will be received by *any* service that has a `charge` verb, regardless of its name or ID. This is useful for broad, system-wide commands.

2.  **`$SRV.charge.billing`**: This subject targets all instances of a specific service by name. If you have multiple `billing` service instances running for load balancing, a request to this subject will be routed to one of them (thanks to queue groups). This is the most common way to call a service.

3.  **`$SRV.charge.billing.some-unique-id`**: This is the most specific subject. It targets a single, specific instance of a service. This is used by the framework for internal monitoring and can be used for debugging or sticky sessions.

This layered approach gives you fine-grained control over request routing.

> **Feynman Box: The Registry's Internal Filing System**
>
> Imagine a massive library with no central card catalog. How would you find a book?
>
> The `$SRV` protocol is like a clever filing system. Every book (service endpoint) is placed in three different spots.
>
> - On a shelf for its genre (e.g., "Science Fiction"). This is like `$SRV.charge`. Any sci-fi book will do.
> - On a shelf for its author (e.g., "Asimov"). This is like `$SRV.charge.billing`. You want a book by Asimov.
> - A specific copy with a unique serial number. This is like `$SRV.charge.billing.some-unique-id`. You want *that specific copy* of the book.
>
> The `nats micro ls` command is like a librarian who quickly sends out a request to all the "author" and "genre" shelves, asking them to report back what they have. This way, you can build a complete catalog of the entire library on the fly, without ever needing a central, single point of failure.

### How `nats micro ls` Works

The `nats micro ls` command (and its programmatic equivalent) is a perfect example of the `$SRV` protocol in action. It doesn't connect to a registry. Instead, it sends out a broadcast request on the `$SRV.PING` subject.

Every running microservice has a built-in handler that listens on `$SRV.PING`, `$SRV.INFO`, and `$SRV.STATS`.

-   **`$SRV.PING`**: When a service receives a PING request, it responds with its `INFO` data to a private inbox provided by the requester. The `ls` command collects all these responses.
-   **`$SRV.INFO`**: You can also request INFO from a specific service.
-   **`$SRV.STATS`**: This provides runtime statistics for the service and its endpoints.

This discovery mechanism is decentralized, fast, and resilient.

## 2. Service Lifecycle: From Birth to Shutdown

A microservice instance has a well-defined lifecycle managed by the framework. Understanding this lifecycle is key to managing state and ensuring graceful operation.

-   **`AddService()`**: This is the birth of your service. When you call `micro.AddService()`, the framework creates all the necessary `$SRV` subscriptions for your service and its endpoints. At this point, your service is "live" and can begin receiving requests.

-   **Request Handling**: As requests come in, the framework invokes your handler functions. For each request, it also starts a timer to measure processing time and increments the request counter.

-   **`Stop()`**: This initiates a graceful shutdown. The framework "drains" the NATS subscriptions. Draining means the subscriptions remain active, but NATS is instructed not to send any new messages. The service will continue to process any requests that are already in its buffer. This prevents you from dropping in-flight requests. The `Stop()` method returns an error if the service was already stopped.

-   **`Done()`**: The `Done()` method returns a channel that is closed when the service has finished processing all in-flight requests and the `Stop()` process is complete. This is useful for blocking in your `main` function to wait for a clean exit.

-   **`DoneHandler`**: You can register a `DoneHandler` in your service configuration. This is a callback function that the framework will execute *after* the service has fully stopped. It's the perfect place to perform cleanup tasks like closing database connections or releasing other resources.

> **Feynman Box: Opening Shop, Serving Customers, Closing Time**
>
> Think of your service like a small coffee shop.
>
> - **`AddService()`** is unlocking the doors and flipping the "Open" sign. The cash registers (endpoints) are ready to take orders.
> - **Request Handling** is the day-to-day business. Customers come in, you serve them coffee, and you keep a tally of how many coffees you've sold.
> - **`Stop()`** is announcing, "We're closing in 15 minutes!" You don't take any new customers, but you finish making drinks for everyone who is already inside.
> - The **`DoneHandler`** is what you do after the last customer leaves. You clean the counters, turn off the espresso machine, and lock the door. The shop is now fully closed and ready for the next day.

## 3. Stats Collection: Knowing Your Service's Performance

Observability is a first-class citizen in NATS Micro. The framework automatically collects key performance indicators for every service and endpoint. You can query these stats using `nats micro stats` or programmatically.

The core `EndpointStats` struct gives you:

-   `NumRequests`: The total number of requests this endpoint has received since the service started.
-   `NumErrors`: The number of times the handler returned an error using `req.Error()`.
-   `ProcessingTime`: The total time spent in the handler for all requests.
-   `AverageProcessingTime`: The average processing time per request.

This built-in data is invaluable for understanding how your service is performing at a glance.

### Custom `StatsHandler`

Sometimes, the default metrics aren't enough. You might want to track business-specific metrics, like the number of transactions of a certain type or the cache hit/miss ratio.

The framework provides a `StatsHandler` hook in the service configuration. This is a function that gets called for every single request, giving you access to the endpoint stats and the request object itself.

Here’s a simple example of a custom stats handler:

```go
// A custom handler to print stats for each request
func customStatsHandler(e micro.Endpoint, r micro.Request) {
    stats := e.Stats()
    log.Printf(
        "Service: %s, Endpoint: %s, Requests: %d, AvgTime: %s",
        stats.Name,
        e.Name(),
        stats.NumRequests,
        stats.AverageProcessingTime,
    )
}

// Add it to your service config
config := micro.Config{
    Name: "OrderProcessor",
    Version: "1.0.0",
    StatsHandler: customStatsHandler,
}
```

You can use this hook to:

-   Push metrics to a monitoring system like Prometheus or Datadog.
-   Log detailed performance data.
-   Implement custom alerting based on application-level metrics.

> **Feynman Box: The Dashboard Showing How Busy Each Waiter Is**
>
> Imagine our coffee shop again. The built-in stats are like the manager keeping a basic tally: "Waiter-A served 50 customers and dropped 2 trays. Waiter-B served 60 customers and dropped 1 tray."
>
> A `StatsHandler` is like giving each waiter a detailed timesheet. Now the manager can see *not just* how many customers they served, but also *what kind* of orders they took (espressos vs. lattes), how long each type of order took, and maybe even how many times they had to remake a drink.
>
> This gives the manager a much richer view of the shop's performance, allowing them to make smarter decisions, like scheduling more staff during the morning latte rush.


