package config

import "time"

// Config represents top-level xtai configuration loaded from file or
// environment. For Phase 1 this is a minimal structure with a simple
// in-memory loader.
type Config struct {
	LLM      LLMConfig     `json:"llm" yaml:"llm"`
	Timeouts TimeoutConfig `json:"timeouts" yaml:"timeouts"`
}

// LLMConfig configures the LLM provider used by xtai.
type LLMConfig struct {
	Provider string `json:"provider" yaml:"provider"`
	Model    string `json:"model" yaml:"model"`
}

// TimeoutConfig configures default operation timeouts.
type TimeoutConfig struct {
	Command time.Duration `json:"command" yaml:"command"`
}

// Load returns a placeholder configuration. File-based loading will be added
// in a later phase when configuration formats are finalized.
func Load() Config {
	return Config{
		LLM: LLMConfig{
			Provider: "openai",
			Model:    "gpt-5",
		},
		Timeouts: TimeoutConfig{
			Command: 60 * time.Second,
		},
	}
}

