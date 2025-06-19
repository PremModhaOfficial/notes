package main

import (
	"context"
	"log"
	"time"
	"trasingExample/setup"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func span_main() {
	tp, err := setup.NewTracerProvider()
	if err != nil {
		log.Fatal(err)
	}

	tracer := tp.Tracer("span.go")

	_, span := tracer.Start(
		context.Background(), "span",
		trace.WithSpanKind(trace.SpanKindServer), // the method is ni trace but it returns the "SPAN" with that kind
	)

	span.SetStatus(codes.Error)
	time.Sleep(time.Millisecond * 300)
	span.End()
}
