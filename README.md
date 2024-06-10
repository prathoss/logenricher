# logextractor

logextractor is tiny package enabling `log/slog` logs enrichment with data
from context.

Use cases:
- correlation id enrichment
- span and trace id enrichment

Example:

```go
package main

import (
    "context"
    "log/slog"

    "github.com/google/uuid"
    "github.com/prathoss/logextractor"
)

func CorrelationIDExtractor(ctx context.Context) []slog.Attr {
    correlationID := ctx.Value("correlation-id").(uuid.UUID)
    return []slog.Attr{slog.String("correlation_id", correlationID.String())}
}

func main() {
    slog.SetDefault(slog.New(&logextractor.SlogHandlerWrapper{
        Handler: slog.NewTextHandler(
            os.Stdout,
            &slog.HandlerOptions{},
        ),
        extractors: []logextractor.Extractor{
            CorrelationIDExtractor,
        },
    }))
}

```