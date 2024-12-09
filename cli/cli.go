package cli

import (
	"github.com/spf13/cobra"
)

// CommandConfig holds the configuration for creating a Cobra command.
type CommandConfig struct {
	Use         string                                  // The usage string for the command
	Short       string                                  // A short description of the command
	Long        string                                  // A detailed description of the command
	Run         func(cmd *cobra.Command, args []string) // The function to execute when the command is called
	SubCommands []*cobra.Command                        // Subcommands to add under this command
	Flags       []FlagConfig                            // Flags to configure for this command
}

// FlagConfig defines the configuration for a command-line flag.
type FlagConfig struct {
	Name         string      // The name of the flag
	Short        string      // The short (single-letter) name of the flag
	DefaultValue interface{} // The default value of the flag
	Usage        string      // A description of the flag's usage
	Required     bool        // Whether the flag is required
}

// NewCommand creates a new Cobra command based on the provided CommandConfig.
func NewCommand(config CommandConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   config.Use,
		Short: config.Short,
		Long:  config.Long,
		Run:   config.Run,
	}

	// Add flags to the command
	for _, flag := range config.Flags {
		switch v := flag.DefaultValue.(type) {
		case string:
			cmd.Flags().StringP(flag.Name, flag.Short, v, flag.Usage)
		case int:
			cmd.Flags().IntP(flag.Name, flag.Short, v, flag.Usage)
		case bool:
			cmd.Flags().BoolP(flag.Name, flag.Short, v, flag.Usage)
		default:
			panic("unsupported flag type") // Handle unsupported flag types
		}
		if flag.Required {
			_ = cmd.MarkFlagRequired(flag.Name) // Mark the flag as required
		}
	}

	// Add subcommands to the command
	for _, sub := range config.SubCommands {
		cmd.AddCommand(sub)
	}

	return cmd
}
