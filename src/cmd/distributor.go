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
}

var createDistributorSubCmd = &cobra.Command{
	Use:     "create",
	Short:   "create a new distributor",
	Aliases: []string{"new"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			cmd.MarkFlagRequired("name")
			cmd.MarkFlagRequired("owner")
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
	Use:     "list",
	Short:   "list distributors",
	Aliases: []string{"ls"},
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
	Use:     "remove",
	Short:   "remove a distributor",
	Aliases: []string{"rm"},
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

var createDistributorCouponSubCmd = &cobra.Command{
	Use:     "create-distributor-code",
	Short:   "create a new distributor code",
	Aliases: []string{"new-distributor-code"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("distributor-code-name")
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
			if err = tui.Send(cmd, args, action.CreateDistributorCoupon); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateDistributorCouponInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listDistributorCouponsSubCmd = &cobra.Command{
	Use:   "list-distributor-codes",
	Short: "list distributor codes",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.ListDistributorCoupons); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListDistributorCouponsInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var describeDistributorCouponSubCmd = &cobra.Command{
	Use:   "describe-distributor-code",
	Short: "describe a distributor code",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no distributor id/name argument provided")
				cmd.Usage()
				os.Exit(1)
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.DescribeDistributorCoupon); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeDistributorCouponInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editDistributorCouponSubCmd = &cobra.Command{
	Use:   "edit-distributor-code",
	Short: "edit a distributor code",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no distributor code id/name argument provided")
				cmd.Usage()
				os.Exit(1)
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.EditDistributorCoupon); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditDistributorCouponInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var revokeDistributorCouponSubCmd = &cobra.Command{
	Use:   "revoke-distributor-code",
	Short: "revoke distributor code",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no distributor code id/name argument provided")
				cmd.Usage()
				os.Exit(1)
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.RevokeDistributorCoupon); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RevokeDistributorCouponInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeDistributorCouponSubCmd = &cobra.Command{
	Use:   "remove-distributor-code",
	Short: "remove distributor code",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no distributor code id/name argument provided")
				cmd.Usage()
				os.Exit(1)
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.RemoveDistributorCoupon); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveDistributorCouponInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var inviteDistributorCouponSubCmd = &cobra.Command{
	Use:   "invite-distributor-code",
	Short: "invite distributor code",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no distributor code id/name argument provided")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("emails")

			emails, _ := cmd.Flags().GetStringSlice("emails")
			if len(emails) == 0 {
				fmt.Println("Error: no emails provided")
				cmd.Usage()
				os.Exit(1)
			}

			basePolicies, _ := cmd.Flags().GetStringSlice("base-policies")
			if len(basePolicies) > 0 {
				for _, policy := range basePolicies {
					if policy != "create-tenant" && policy != "create-swarm" {
						fmt.Println("invalid value '%s' for --base-policies; allowed values are: create-tenant, create-swarm", policy)
						cmd.Usage()
						os.Exit(1)
					}
				}
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.InviteDistributorCoupon); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.InviteDistributorCouponInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var reportDistributorSubCmd = &cobra.Command{
	Use:   "report",
	Short: "downloads/prints a full report for the distributor",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("from")
			cmd.MarkFlagRequired("to")

			isChanged := cmd.Flags().Changed("output")
			if isChanged {
				output, _ := cmd.Flags().GetString("output")

				if output == "" {
					fmt.Println("Error: output cannot be empty.Use a dot (.) to indicate the current directory.")
					cmd.Usage()
					os.Exit(1)
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.GetDistributorReport); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.GetDistributorReportInteractive(cmd); err != nil {
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

	distributorCmd.AddCommand(reportDistributorSubCmd)
	reportDistributorSubCmd.Flags().String("from", "", "Start date and time in DD/MM/YYYY+HH:mm:ss format")
	reportDistributorSubCmd.Flags().String("to", "", "End date and time in DD/MM/YYYY+HH:mm:ss format")
	reportDistributorSubCmd.Flags().String("distributor-code", "", "The distributor code")
	reportDistributorSubCmd.Flags().String("format", "json", "Formats the result")
	reportDistributorSubCmd.Flags().StringP("output", "o", "", "Specify the output file or directory.Use a dot (.) to indicate the current directory.")

	distributorCmd.AddCommand(createDistributorCouponSubCmd)
	createDistributorCouponSubCmd.Flags().String("distributor-code-name", "", "Name of the distributor code")
	createDistributorCouponSubCmd.Flags().String("description", "", "Description of the distributor code")
	createDistributorCouponSubCmd.Flags().Int("redemption-count", -1, "Max redemptions of the distributor code")
	createDistributorCouponSubCmd.Flags().StringSlice("swarms", []string{}, "List of swarm ids associated to the distributor code")
	createDistributorCouponSubCmd.Flags().String("zone", "", "Zone of the distributor code creation")
	createDistributorCouponSubCmd.Flags().String("external-id", "", "External ID of the distributor code")

	distributorCmd.AddCommand(listDistributorCouponsSubCmd)
	listDistributorCouponsSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for distributor codes")
	listDistributorCouponsSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different distributor codes")

	distributorCmd.AddCommand(describeDistributorCouponSubCmd)
	describeDistributorCouponSubCmd.Flags().String("format", "default", "Formats the output")

	distributorCmd.AddCommand(editDistributorCouponSubCmd)
	editDistributorCouponSubCmd.Flags().String("distributor-code-name", "", "New name of the distributor code")
	editDistributorCouponSubCmd.Flags().String("description", "", "New description of the distributor code")
	editDistributorCouponSubCmd.Flags().Int("redemption-count", 0, "New max redemptions of the distributor code")
	editDistributorCouponSubCmd.Flags().String("external-id", "", "New external ID of the distributor code")

	distributorCmd.AddCommand(revokeDistributorCouponSubCmd)

	distributorCmd.AddCommand(removeDistributorCouponSubCmd)

	distributorCmd.AddCommand(inviteDistributorCouponSubCmd)
	inviteDistributorCouponSubCmd.Flags().StringSlice("emails", []string{}, "list of emails to invite")
	inviteDistributorCouponSubCmd.Flags().StringSlice("base-policies", []string{}, "list of base policies to invite, allowed values are: create-tenant, create-swarm")

	rootCmd.AddCommand(distributorCmd)
	distributorCmd.PersistentFlags().String("name", "", "Name of the distributor")
	distributorCmd.PersistentFlags().String("id", "", "ID of the distributor")
}
