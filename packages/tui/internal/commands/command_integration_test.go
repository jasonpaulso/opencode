package commands

import (
	"testing"

	"github.com/sst/opencode-sdk-go"
)

func TestUserDefinedCommandsIntegration(t *testing.T) {
	// Test that user-defined commands integrate properly with the command system
	config := &opencode.Config{
		Commands: map[string]opencode.ConfigCommand{
			"test_command": {
				Description: "Test command for integration",
				Keybind:     "<leader>t",
				Trigger:     []string{"test", "integration"},
			},
		},
	}

	registry := LoadFromConfig(config)

	// Test that the user command can be found by name
	userCommand := registry[CommandName("user_test_command")]
	if userCommand.Name != CommandName("user_test_command") {
		t.Errorf("Expected command name 'user_test_command', got '%s'", userCommand.Name)
	}

	// Test that the command has the right properties for triggers
	if !userCommand.HasTrigger() {
		t.Error("Expected user command to have triggers")
	}

	if userCommand.PrimaryTrigger() != "test" {
		t.Errorf("Expected primary trigger 'test', got '%s'", userCommand.PrimaryTrigger())
	}

	if !userCommand.MatchesTrigger("integration") {
		t.Error("Expected command to match 'integration' trigger")
	}

	// Test sorting includes user commands
	sorted := registry.Sorted()
	found := false
	for _, cmd := range sorted {
		if cmd.Name == CommandName("user_test_command") {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected user command to be included in sorted commands")
	}
}

func TestUserCommandsWithComplexKeybindings(t *testing.T) {
	// Test complex keybinding scenarios
	config := &opencode.Config{
		Commands: map[string]opencode.ConfigCommand{
			"multi_key": {
				Description: "Multi-key command",
				Keybind:     "ctrl+shift+alt+m, <leader>mm",
				Trigger:     []string{"multi"},
			},
		},
	}

	registry := LoadFromConfig(config)
	userCommand := registry[CommandName("user_multi_key")]

	if len(userCommand.Keybindings) != 2 {
		t.Errorf("Expected 2 keybindings, got %d: %+v", len(userCommand.Keybindings), userCommand.Keybindings)
	}

	// Check first keybinding
	if userCommand.Keybindings[0].Key != "ctrl+shift+alt+m" || userCommand.Keybindings[0].RequiresLeader {
		t.Errorf("Expected first keybinding 'ctrl+shift+alt+m' without leader, got %+v", userCommand.Keybindings[0])
	}

	// Check second keybinding - the parseBindings function should detect <leader> and strip it
	if userCommand.Keybindings[1].Key != "mm" || !userCommand.Keybindings[1].RequiresLeader {
		t.Errorf("Expected second keybinding 'mm' with leader, got %+v", userCommand.Keybindings[1])
	}
}

func TestUserCommandsConflictPrevention(t *testing.T) {
	// Test that user commands don't accidentally override defaults
	config := &opencode.Config{
		Commands: map[string]opencode.ConfigCommand{
			"help": {
				Description: "User help command",
				Keybind:     "<leader>h",
				Trigger:     []string{"help"},
			},
		},
	}

	registry := LoadFromConfig(config)

	// Both default and user commands should exist
	defaultHelp, defaultExists := registry[AppHelpCommand]
	userHelp, userExists := registry[CommandName("user_help")]

	if !defaultExists {
		t.Error("Expected default help command to still exist")
	}

	if !userExists {
		t.Error("Expected user help command to exist")
	}

	// They should have different names but potentially same triggers
	if defaultHelp.Name == userHelp.Name {
		t.Error("Default and user commands should have different names")
	}

	// Both can have the same trigger - this is allowed for flexibility
	if !defaultHelp.MatchesTrigger("help") || !userHelp.MatchesTrigger("help") {
		t.Error("Both commands should match 'help' trigger")
	}
}