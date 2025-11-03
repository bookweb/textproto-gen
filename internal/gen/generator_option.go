package gen

import "log/slog"

// GeneratorOption is option for Generator.
type GeneratorOption func(g *Generator)

// WithTitle sets title.
func WithTitle(title string) GeneratorOption {
	return func(g *Generator) {
		g.SetTitle(title)
	}
}

// WithDescription sets description.
func WithDescription(description string) GeneratorOption {
	return func(g *Generator) {
		g.SetDescription(description)
	}
}

// WithLogger sets logger.
func WithLogger(logger *slog.Logger) GeneratorOption {
	return func(g *Generator) {
		g.logger = logger
	}
}
