package commands

import (
	"testing"

	"github.com/sst/opencode-sdk-go"
)

func TestLoadFromConfig_UserDefinedCommands(t *testing.T) {
	// Test configuration with user-defined commands
	config := &opencode.Config{
		Commands: map[string]opencode.ConfigCommand{
			"my_command": {
				Description: "my custom command",
				Keybind:     "<leader>x",
				Trigger:     []string{"mycmd", "custom"},
			},
			"another_command": {
				Description: "another custom command",
				Trigger:     []string{"another"},
			},
		},
	}

	registry := LoadFromConfig(config)

	// Check that user commands are loaded
	userCommand1, exists := registry[CommandName("user_my_command")]
	if !exists {
		t.Error("Expected user_my_command to be loaded")
	}

	if userCommand1.Description != "my custom command" {
		t.Errorf("Expected description 'my custom command', got '%s'", userCommand1.Description)
	}

	if len(userCommand1.Keybindings) != 1 || userCommand1.Keybindings[0].Key != "x" || !userCommand1.Keybindings[0].RequiresLeader {
		t.Errorf("Expected keybinding '<leader>x', got %+v", userCommand1.Keybindings)
	}

	if len(userCommand1.Trigger) != 2 || userCommand1.Trigger[0] != "mycmd" || userCommand1.Trigger[1] != "custom" {
		t.Errorf("Expected triggers ['mycmd', 'custom'], got %v", userCommand1.Trigger)
	}

	userCommand2, exists := registry[CommandName("user_another_command")]
	if !exists {
		t.Error("Expected user_another_command to be loaded")
	}

	if userCommand2.Description != "another custom command" {
		t.Errorf("Expected description 'another custom command', got '%s'", userCommand2.Description)
	}

	if len(userCommand2.Keybindings) != 0 {
		t.Errorf("Expected no keybindings, got %+v", userCommand2.Keybindings)
	}

	if len(userCommand2.Trigger) != 1 || userCommand2.Trigger[0] != "another" {
		t.Errorf("Expected trigger ['another'], got %v", userCommand2.Trigger)
	}

	// Check that default commands are still present
	_, exists = registry[AppHelpCommand]
	if !exists {
		t.Error("Expected default AppHelpCommand to be present")
	}
}

func TestLoadFromConfig_EmptyUserCommands(t *testing.T) {
	// Test configuration with no user-defined commands
	config := &opencode.Config{}

	registry := LoadFromConfig(config)

	// Check that default commands are loaded
	_, exists := registry[AppHelpCommand]
	if !exists {
		t.Error("Expected default AppHelpCommand to be present")
	}

	// Ensure no user commands are present
	for name := range registry {
		if len(string(name)) > 5 && string(name)[:5] == "user_" {
			t.Errorf("Unexpected user command found: %s", name)
		}
	}
}

func TestLoadFromConfig_CustomKeybindings(t *testing.T) {
	// Test that custom keybindings still work with user commands
	config := &opencode.Config{
		Keybinds: opencode.Keybinds{
			AppHelp: "<leader>?",
		},
		Commands: map[string]opencode.ConfigCommand{
			"test_cmd": {
				Description: "test command",
				Keybind:     "ctrl+t",
			},
		},
	}

	registry := LoadFromConfig(config)

	// Check that custom keybinding for default command is applied
	helpCommand, exists := registry[AppHelpCommand]
	if !exists {
		t.Error("Expected AppHelpCommand to be present")
	}

	if len(helpCommand.Keybindings) != 1 || helpCommand.Keybindings[0].Key != "?" || !helpCommand.Keybindings[0].RequiresLeader {
		t.Errorf("Expected custom keybinding '<leader>?', got %+v", helpCommand.Keybindings)
	}

	// Check that user command has correct keybinding
	testCommand, exists := registry[CommandName("user_test_cmd")]
	if !exists {
		t.Error("Expected user_test_cmd to be present")
	}

	if len(testCommand.Keybindings) != 1 || testCommand.Keybindings[0].Key != "ctrl+t" || testCommand.Keybindings[0].RequiresLeader {
		t.Errorf("Expected keybinding 'ctrl+t', got %+v", testCommand.Keybindings)
	}
}