package main

import (
	"context"
	"log"
	"trasingExample/setup"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	tp, err := setup.NewTracerProvider()
	if err != nil {
		log.Fatal(err)
	}

	tracer := tp.Tracer("span.go")

	_, span := tracer.Start(
		context.Background(),
		"span",
		trace.WithSpanKind(trace.SpanKindServer), // the method is ni trace but it returns the "SPAN" with that kind
	)
	span.SetAttributes(attribute.String("keys", log.Prefix()))

	span.End()

	print(span)
}
