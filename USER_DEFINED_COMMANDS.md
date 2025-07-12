# User-Defined Commands

OpenCode TUI now supports user-defined commands that can be configured in your `opencode.json` configuration file. This allows you to extend the built-in command set with your own custom commands.

## Configuration Format

Add a `commands` section to your `opencode.json` file:

```json
{
  "commands": {
    "command_name": {
      "description": "Description of what the command does",
      "keybind": "optional keybinding",
      "trigger": ["text", "triggers", "for", "command"]
    }
  }
}
```

## Command Properties

Each user-defined command can have the following properties:

- **description** (required): A human-readable description of what the command does
- **keybind** (optional): Keyboard shortcut(s) to trigger the command
- **trigger** (optional): Text triggers that can be typed to invoke the command (e.g., "/custom")

## Keybinding Format

Keybindings support the following formats:

- Simple keys: `"a"`, `"enter"`, `"esc"`
- Modifier combinations: `"ctrl+c"`, `"shift+enter"`, `"ctrl+shift+alt+m"`
- Leader key combinations: `"<leader>h"`, `"<leader>custom"`
- Multiple keybindings: `"ctrl+c, <leader>q"` (comma-separated)

The leader key is configured in the `keybinds.leader` setting (defaults to `"ctrl+x"`).

## Command Naming

User-defined commands are automatically prefixed with `"user_"` internally to avoid conflicts with built-in commands. For example, a command named `"my_command"` becomes `"user_my_command"` in the system.

## Example Configuration

```json
{
  "$schema": "https://opencode.ai/config.json",
  "commands": {
    "quick_note": {
      "description": "Take a quick note",
      "keybind": "<leader>n",
      "trigger": ["note", "quick"]
    },
    "toggle_feature": {
      "description": "Toggle experimental feature",
      "keybind": "ctrl+shift+t",
      "trigger": ["toggle", "feature"]
    },
    "custom_help": {
      "description": "Show custom help",
      "trigger": ["help", "custom"]
    }
  },
  "keybinds": {
    "leader": "ctrl+x"
  }
}
```

## Usage

Once configured, user-defined commands:

1. **Appear in command completions**: Type `/` and start typing your trigger to see your custom commands
2. **Can be triggered by keybindings**: Press the configured key combination
3. **Can be triggered by text**: Type `/` followed by any of the configured triggers
4. **Show up in help**: User commands are listed alongside built-in commands

## Integration Notes

- User-defined commands integrate seamlessly with the existing command system
- They support all the same features as built-in commands (keybindings, triggers, descriptions)
- Multiple commands can share the same trigger (useful for providing alternatives)
- User commands can have the same triggers as built-in commands (both will be available)
- Custom keybindings for built-in commands can still be configured in the `keybinds` section

## Command Execution

**Note**: Currently, user-defined commands are recognized and displayed in the UI but do not execute custom actions. This feature provides the foundation for future extensibility where users will be able to define custom behavior for their commands.