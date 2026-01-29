package cmd_agent

import (
	"fmt"
	"os"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewAgentSubCmdCreate(
	newAgentService service.AgentServiceInterface,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a new agent",
		PreRunE: func(cmd *cobra.Command, args []string) error {

			cmd.MarkFlagRequired("nexus-id")
			cmd.MarkFlagRequired("node-id")

			batch, _ := cmd.Flags().GetBool("batch")
			if batch {
				file, _ := cmd.Flags().GetString("file")
				if file == "" {
					cmd.Usage()
					return fmt.Errorf("--file flag is required when using --batch mode")
				}

				if _, err := os.Stat(file); os.IsNotExist(err) {
					return fmt.Errorf("file does not exist: %s", file)
				}
				return nil
			}

			cmd.MarkFlagRequired("agent-port")
			cmd.MarkFlagRequired("agent-disk")
			cmd.MarkFlagRequired("agent-mount-point")

			agentPort := cmd.Flags().Lookup("agent-port")
			if agentPort != nil && !agentPort.Changed {
				cmd.Usage()
				return fmt.Errorf("required flag(s) \"agent-port\" not set")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			batch, _ := cmd.Flags().GetBool("batch")
			if batch {
				if err := newAgentService.CreateAgentBatch(cmd, args); err != nil {
					utils.PrintError(err)
				}
				return
			}

			if err := newAgentService.CreateAgent(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	cmd.Flags().Int("agent-port", 0, "Port of the agent")
	cmd.Flags().String("agent-disk", "", "Disk of the agent")
	cmd.Flags().String("agent-mount-point", "", "Mount point of the agent")
	cmd.Flags().String("agent-features", "", "Features of the agent")
	cmd.Flags().Bool("batch", false, "Create multiple agents from a batch file")
	cmd.Flags().String("file", "", "Path to the JSON file containing agent definitions")

	return cmd
}
