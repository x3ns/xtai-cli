package llm

import "context"

// Provider defines the abstraction for all LLM providers xtai can use.
// Implementations live in infrastructure/llm and must be safe to mock for
// tests. The interface is intentionally small for now and can be extended if
// additional capabilities (such as streaming or embeddings) are required.
type Provider interface {
	// Generate produces a response for the given prompt.
	Generate(ctx context.Context, prompt string) (string, error)
}

