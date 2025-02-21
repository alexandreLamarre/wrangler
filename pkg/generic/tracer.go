package generic

import "go.opentelemetry.io/otel"

var (
	genericTracer = otel.Tracer("GenericController")
)
