package completions

import (
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode/internal/app"
	"github.com/sst/opencode/internal/commands"
)

func TestUserDefinedCommandsInCompletions(t *testing.T) {
	// Create a mock app with user-defined commands
	config := &opencode.Config{
		Commands: map[string]opencode.ConfigCommand{
			"my_custom": {
				Description: "My custom command",
				Trigger:     []string{"custom", "mycmd"},
			},
			"another_cmd": {
				Description: "Another command",
				Trigger:     []string{"another"},
			},
		},
	}

	commandRegistry := commands.LoadFromConfig(config)
	
	// Create a minimal mock app
	mockApp := &app.App{
		Commands: commandRegistry,
	}

	provider := NewCommandCompletionProvider(mockApp)

	// Verify that the provider can handle user commands without crashing
	// and that the command registry contains the expected user commands
	
	// Check that user commands are present in the registry
	userCustom, exists := commandRegistry[commands.CommandName("user_my_custom")]
	if !exists {
		t.Error("Expected user_my_custom to be in command registry")
	}
	if userCustom.Description != "My custom command" {
		t.Errorf("Expected description 'My custom command', got '%s'", userCustom.Description)
	}
	if !userCustom.HasTrigger() || userCustom.PrimaryTrigger() != "custom" {
		t.Errorf("Expected primary trigger 'custom', got '%s'", userCustom.PrimaryTrigger())
	}

	userAnother, exists := commandRegistry[commands.CommandName("user_another_cmd")]
	if !exists {
		t.Error("Expected user_another_cmd to be in command registry")
	}
	if userAnother.Description != "Another command" {
		t.Errorf("Expected description 'Another command', got '%s'", userAnother.Description)
	}
	if !userAnother.HasTrigger() || userAnother.PrimaryTrigger() != "another" {
		t.Errorf("Expected primary trigger 'another', got '%s'", userAnother.PrimaryTrigger())
	}

	// Verify the provider can be created without error
	if provider.GetId() != "commands" {
		t.Errorf("Expected provider ID 'commands', got '%s'", provider.GetId())
	}

	if provider.GetEmptyMessage() != "no matching commands" {
		t.Errorf("Expected empty message 'no matching commands', got '%s'", provider.GetEmptyMessage())
	}
}