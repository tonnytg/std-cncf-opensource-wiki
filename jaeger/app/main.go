package main

import (
	"context"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	// Configuração do exportador Jaeger
	exp, err := jaeger.NewRawExporter(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "my-service",
			Tags: []jaeger.Tag{
				jaeger.StringTag("jaeger.version", "1.0.0"),
			},
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := exp.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Configuração do provedor de trace
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes()))
	otel.SetTracerProvider(tp)

	// Criação de um span root
	tr := otel.Tracer("example")

	// Criação de um span filho
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tr.Start(r.Context(), "handleRequest")
		defer span.End()

		// Seu código aqui

		w.Write([]byte("Hello World"))
	})

	// Configuração do servidor HTTP com middleware OpenTelemetry
	h := otelhttp.NewHandler(http.DefaultServeMux, "server")
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatal(err)
	}
}
