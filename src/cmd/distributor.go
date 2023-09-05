package cmd

import (
	"fmt"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var distributorCmd = &cobra.Command{
	Use:   "distributor",
	Short: "Execute commands in distributor sections",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var createDistributorSubCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new distributor",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			cmd.MarkFlagRequired("name")
			cmd.MarkFlagRequired("description")
			cmd.MarkFlagRequired("owner")
			cmd.MarkFlagRequired("first-name")
			cmd.MarkFlagRequired("last-name")
			cmd.MarkFlagRequired("swarms")

			swarms, _ := cmd.Flags().GetStringSlice("swarms")
			if len(swarms) == 0 {
				fmt.Println("Error: no swarms provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.CreateDistributor); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateDistributorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listDistributorSubCmd = &cobra.Command{
	Use:   "list",
	Short: "list distributors",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.ListDistributor); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListDistributorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeDistributorSubCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a distributor",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("password")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.RemoveDistributor); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveDistributorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {
	distributorCmd.AddCommand(createDistributorSubCmd)
	createDistributorSubCmd.Flags().String("name", "", "Name of the distributor")
	createDistributorSubCmd.Flags().String("description", "", "Description of the distributor")
	createDistributorSubCmd.Flags().String("image-url", "", "Image URL of the distributor")
	createDistributorSubCmd.Flags().String("owner", "", "Email of the invited distributor operator")
	createDistributorSubCmd.Flags().String("first-name", "", "First name of the invited distributor operator")
	createDistributorSubCmd.Flags().String("last-name", "", "Last name of the invited distributor operator")
	createDistributorSubCmd.Flags().StringSlice("swarms", []string{}, "List of swarm ids associated to the distributor")

	distributorCmd.AddCommand(listDistributorSubCmd)
	listDistributorSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for distributors")
	listDistributorSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different distributors")

	distributorCmd.AddCommand(removeDistributorSubCmd)
	removeDistributorSubCmd.Flags().String("email", "", "Email address")
	removeDistributorSubCmd.Flags().String("password", "", "Password")
	removeDistributorSubCmd.Flags().String("code", "", "Two factor authentication code")

	rootCmd.AddCommand(distributorCmd)
	distributorCmd.PersistentFlags().String("name", "", "Name of the distributor")
	distributorCmd.PersistentFlags().String("id", "", "ID of the distributor")
}
