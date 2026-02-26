package llm

import (
	"context"
	"errors"

	domainllm "github.com/x3ns/xtai-cli/internal/domain/llm"
)

// OpenAIProvider is a placeholder implementation of llm.Provider for OpenAI.
// Real HTTP calls and configuration will be added in later phases.
type OpenAIProvider struct{}

// NewOpenAIProvider constructs a new OpenAIProvider.
func NewOpenAIProvider() *OpenAIProvider {
	return &OpenAIProvider{}
}

// Ensure OpenAIProvider satisfies the llm.Provider interface.
var _ domainllm.Provider = (*OpenAIProvider)(nil)

// Generate is a placeholder implementation that will be replaced with a real
// OpenAI API integration in a later phase.
func (p *OpenAIProvider) Generate(ctx context.Context, prompt string) (string, error) {
	_ = ctx
	_ = prompt
	return "", errors.New("OpenAI provider is not implemented yet")
}

