package domain

import "context"

// Provider defines the abstraction for all LLM providers xtai can use.
// Implementations live in infrastructure/llm and must be safe to mock for tests.
type Provider interface {
	// Generate produces a response for the given prompt.
	// The concrete request/response shapes may be refined as requirements evolve.
	Generate(ctx context.Context, prompt string) (string, error)
}

