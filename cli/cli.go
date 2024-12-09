package cli

import (
	"github.com/spf13/cobra"
)

// CommandConfig holds the configuration for creating a command.
type CommandConfig struct {
	Use         string
	Short       string
	Long        string
	Run         func(cmd *cobra.Command, args []string)
	SubCommands []*cobra.Command // To add subcommands
	Flags       []FlagConfig
}

// FlagConfig defines a flag configuration.
type FlagConfig struct {
	Name         string
	Short        string
	DefaultValue interface{}
	Usage        string
	Required     bool
}

// NewCommand creates a new cobra command based on CommandConfig.
func NewCommand(config CommandConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   config.Use,
		Short: config.Short,
		Long:  config.Long,
		Run:   config.Run,
	}

	// Add flags
	for _, flag := range config.Flags {
		switch v := flag.DefaultValue.(type) {
		case string:
			cmd.Flags().StringP(flag.Name, flag.Short, v, flag.Usage)
		case int:
			cmd.Flags().IntP(flag.Name, flag.Short, v, flag.Usage)
		case bool:
			cmd.Flags().BoolP(flag.Name, flag.Short, v, flag.Usage)
		}
		if flag.Required {
			_ = cmd.MarkFlagRequired(flag.Name)
		}
	}

	// Add subcommands
	for _, sub := range config.SubCommands {
		cmd.AddCommand(sub)
	}

	return cmd
}
