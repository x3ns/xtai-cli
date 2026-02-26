package llm

import (
	"context"
	"errors"

	domainllm "github.com/x3ns/xtai-cli/internal/domain/llm"
)

// OllamaProvider is a placeholder implementation of llm.Provider for a local
// Ollama backend. Real model invocation will be added separately.
type OllamaProvider struct{}

// NewOllamaProvider constructs a new OllamaProvider.
func NewOllamaProvider() *OllamaProvider {
	return &OllamaProvider{}
}

// Ensure OllamaProvider satisfies the llm.Provider interface.
var _ domainllm.Provider = (*OllamaProvider)(nil)

// Generate is a placeholder implementation that will be replaced with a real
// Ollama integration.
func (p *OllamaProvider) Generate(ctx context.Context, prompt string) (string, error) {
	_ = ctx
	_ = prompt
	return "", errors.New("Ollama provider is not implemented yet")
}

