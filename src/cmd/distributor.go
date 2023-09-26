package cmd

import (
	"fmt"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var zones = []string{"de", "fr"}

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
	Use:     "create-coupon",
	Short:   "create a new distributor coupon",
	Aliases: []string{"new-coupon"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("coupon-name")
			cmd.MarkFlagRequired("swarms")

			swarms, _ := cmd.Flags().GetStringSlice("swarms")
			if len(swarms) == 0 {
				fmt.Println("Error: no swarms provided")
				cmd.Usage()
				os.Exit(1)
			}

			zone, _ := cmd.Flags().GetString("zone")
			if zone != "" {
				for _, z := range zones {
					if zone == z {
						return
					}
				}

				fmt.Println("Error: the provided zone is invalid, please enter <'fr'|'de'> for the French zone or leave it empty.")
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
	Use:   "list-coupons",
	Short: "list distributor coupons",
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
	Use:   "describe-coupon",
	Short: "describe a distributor coupon",
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
	Use:   "edit-coupon",
	Short: "edit a distributor coupon",
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
				fmt.Println("Error: no distributor coupon id/name argument provided")
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
	Use:   "revoke-coupon",
	Short: "revoke distributor coupon code",
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
				fmt.Println("Error: no distributor coupon id/name argument provided")
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
	Use:   "remove-coupon",
	Short: "remove distributor coupon",
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
				fmt.Println("Error: no distributor coupon id/name argument provided")
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

var reportDistributSubCmd = &cobra.Command{
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

	distributorCmd.AddCommand(reportDistributSubCmd)
	reportDistributSubCmd.Flags().String("from", "", "Start date and time in DD/MM/YYYY+HH:mm:ss format")
	reportDistributSubCmd.Flags().String("to", "", "End date and time in DD/MM/YYYY+HH:mm:ss format")
	reportDistributSubCmd.Flags().String("coupon", "", "The distributor coupon id or name")
	reportDistributSubCmd.Flags().String("format", "json", "Formats the result")
	reportDistributSubCmd.Flags().StringP("output", "o", "", "Specify the output file or directory.Use a dot (.) to indicate the current directory.")

	distributorCmd.AddCommand(createDistributorCouponSubCmd)
	createDistributorCouponSubCmd.Flags().String("coupon-name", "", "Name of the distributor coupon")
	createDistributorCouponSubCmd.Flags().String("description", "", "Description of the distributor coupon")
	createDistributorCouponSubCmd.Flags().Int("redemption-count", -1, "Max redemptions of the distributor coupon")
	createDistributorCouponSubCmd.Flags().StringSlice("swarms", []string{}, "List of swarm ids associated to the distributor coupon")
	createDistributorCouponSubCmd.Flags().String("zone", "", "Zone of the distributor coupon creation")

	distributorCmd.AddCommand(listDistributorCouponsSubCmd)
	listDistributorCouponsSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for distributor coupons")
	listDistributorCouponsSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different distributor coupons")

	distributorCmd.AddCommand(describeDistributorCouponSubCmd)
	describeDistributorCouponSubCmd.Flags().String("format", "default", "Formats the output")

	distributorCmd.AddCommand(editDistributorCouponSubCmd)
	editDistributorCouponSubCmd.Flags().String("coupon-name", "", "New name of the distributor coupon")
	editDistributorCouponSubCmd.Flags().String("description", "", "New description of the distributor coupon")
	editDistributorCouponSubCmd.Flags().Int("redemption-count", 0, "New max redemptions of the distributor coupon")

	distributorCmd.AddCommand(revokeDistributorCouponSubCmd)

	distributorCmd.AddCommand(removeDistributorCouponSubCmd)

	rootCmd.AddCommand(distributorCmd)
	distributorCmd.PersistentFlags().String("name", "", "Name of the distributor")
	distributorCmd.PersistentFlags().String("id", "", "ID of the distributor")
}
