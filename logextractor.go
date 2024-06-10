// Package logenricher is tiny package enabling `log/slog` logs enrichment with data
// from context. Start with SlogHandlerWrapper.
package logenricher

import (
	"context"
	"log/slog"
)

// Extractor represents a function type that extracts attributes from a context.Context object.
// Used in SlogHandlerWrapper for extracting data from context.Context and including them into all logs lines.
type Extractor func(ctx context.Context) []slog.Attr

// SlogHandlerWrapper is a type that wraps a slog.Handler and includes a list of extractors.
// It is used for extracting data from context.Context objects and including them in all log lines.
//
// Usage:
//
//	correlationIDExtractor := func(ctx context.Context) []slog.Attr {
//		correlationID := ctx.Value("correlation-id").(uuid.UUID)
//		return []slog.Attr{slog.String("correlation_id", correlationID.String())}
//	}
//
//	slog.SetDefault(slog.New(&logenricher.SlogHandlerWrapper{
//		Handler: slog.NewTextHandler(
//			os.Stdout,
//			&slog.HandlerOptions{},
//		),
//		Extractors: []logenricher.Extractor{
//			CorrelationIDExtractor,
//		},
//	}))
type SlogHandlerWrapper struct {
	slog.Handler
	Extractors []Extractor
}

func (s *SlogHandlerWrapper) Handle(ctx context.Context, record slog.Record) error {
	for _, extractor := range s.Extractors {
		record.AddAttrs(extractor(ctx)...)
	}

	return s.Handler.Handle(ctx, record)
}
