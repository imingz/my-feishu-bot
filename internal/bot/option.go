package bot

import "log/slog"

// WithLogger uses the provided logger.
func WithLogger(logger *slog.Logger) Option {
	return func(b *Bot) {
		b.logger = logger
	}
}

type Option func(*Bot)
