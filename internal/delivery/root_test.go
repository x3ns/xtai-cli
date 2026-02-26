package delivery

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootCommand_Help(t *testing.T) {
	cmd := NewRootCommand()
	buf := &bytes.Buffer{}

	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"--help"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("expected no error running --help, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "xtai is a natural-language CLI agent") {
		t.Fatalf("expected help output to mention xtai description, got: %q", output)
	}
}

func TestPlanCommand_PrintsPlaceholder(t *testing.T) {
	cmd := NewRootCommand()
	buf := &bytes.Buffer{}

	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"plan"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("expected no error running plan, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "xtai plan: planning is not implemented yet") {
		t.Fatalf("expected plan placeholder output, got: %q", output)
	}
}

func TestExecuteCommand_PrintsPlaceholder(t *testing.T) {
	cmd := NewRootCommand()
	buf := &bytes.Buffer{}

	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"execute"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("expected no error running execute, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "xtai execute: execution engine is not implemented yet") {
		t.Fatalf("expected execute placeholder output, got: %q", output)
	}
}

func TestReplayCommand_PrintsPlaceholder(t *testing.T) {
	cmd := NewRootCommand()
	buf := &bytes.Buffer{}

	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"replay"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("expected no error running replay, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "xtai replay: session replay is not implemented yet") {
		t.Fatalf("expected replay placeholder output, got: %q", output)
	}
}

func TestSessionsCommand_PrintsPlaceholder(t *testing.T) {
	cmd := NewRootCommand()
	buf := &bytes.Buffer{}

	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"sessions"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("expected no error running sessions, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "xtai sessions: session listing is not implemented yet") {
		t.Fatalf("expected sessions placeholder output, got: %q", output)
	}
}

