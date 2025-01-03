package main

import (
	"os/exec"
	"testing"
)

const todoCliPath = "../todo-cli"

func Test_CLI_Integration(t *testing.T) {
	tests := []struct {
		name        string
		commandArgs []string
		wantOutput  string
	}{
		// add tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(todoCliPath, tt.commandArgs...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("\nFailed to execute command: %v", err)
			}

			gotMessage := string(output)
			if gotMessage != tt.wantOutput {
				t.Errorf("\nwantOutput: %v\ngotMessage: %v", tt.wantOutput, gotMessage)
			}
		})
	}
}
