package cmd

import (
	"fmt"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var swarmCmd = &cobra.Command{
	Use:   "swarm",
	Short: "Execute commands in swarm sections",
}

var createSwarmSubCmd = &cobra.Command{
	Use:     "create",
	Short:   "create a new swarm",
	Aliases: []string{"new"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			cmd.MarkFlagRequired("name")
			cmd.MarkFlagRequired("description")
			cmd.MarkFlagRequired("configuration")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.CreateSwarm); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateSwarmInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listSwarmSubCmd = &cobra.Command{
	Use:     "list",
	Short:   "list swarms",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.ListSwarms); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListSwarmsInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var describeSwarmSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe a swarm",
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
			if err = tui.Send(cmd, args, action.DescribeSwarm); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeSwarmInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editSwarmDescriptionSubCmd = &cobra.Command{
	Use:   "edit-description",
	Short: "edit a swarm description",
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
				fmt.Println("Error: no new description argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = action.EditSwarmDescription(cmd, args...); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditSwarmDescriptionInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editSwarmNameSubCmd = &cobra.Command{
	Use:   "edit-name",
	Short: "edit a swarm name",
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
				fmt.Println("Error: no new name argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = action.EditSwarmName(cmd, args...); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditSwarmNameInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeSwarmSubCmd = &cobra.Command{
	Use:     "remove",
	Short:   "remove a swarm",
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
			if err = tui.Send(cmd, args, action.RemoveSwarm); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveSwarmInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var addOperatorToSwarmSubCmd = &cobra.Command{
	Use:   "add-operator",
	Short: "invites an operator",
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
			cmd.MarkFlagRequired("role")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.AddOperatorToSwarm); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.AddOperatorToSwarmInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listSwarmOperatorsSubCmd = &cobra.Command{
	Use:   "list-operators",
	Short: "lists swarm operators",
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
			if err = tui.Send(cmd, args, action.ListSwarmOperators); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListSwarmOperatorsInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeSwarmOperatorSubCmd = &cobra.Command{
	Use:   "remove-operator",
	Short: "removes swarm operator by email or id",
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
				fmt.Println("Error: no operator argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.RemoveSwarmOperator); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveSwarmOperatorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var describeSwarmOperatorsSubCmd = &cobra.Command{
	Use:   "describe-operator",
	Short: "describe swarm operator",
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
				fmt.Println("Error: no operator email or id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.DescribeSwarmOperator); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeSwarmOperatorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var EditSwarmOperatorRoleSubCmd = &cobra.Command{
	Use:   "edit-operator",
	Short: "edit swarm operator role",
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
				fmt.Println("Error: no operator id or email argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.EditSwarmOperatorRole); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditSwarmOperatorRoleInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var createSwarmNexusSubCmd = &cobra.Command{
	Use:   "create-nexus",
	Short: "create a new nexus",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("nexus-name")
			cmd.MarkFlagRequired("location")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.CreateSwarmNexus); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateSwarmNexusInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editSwarmNexusSubCmd = &cobra.Command{
	Use:   "edit-nexus",
	Short: "edit a nexus",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			if len(args) == 0 {
				fmt.Println("Error: no new name argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.EditSwarmNexus); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditSwarmNexusInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeSwarmNexusSubCmd = &cobra.Command{
	Use:   "remove-nexus",
	Short: "remove a nexus",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			if len(args) == 0 {
				fmt.Println("Error: no nexus id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.RemoveSwarmNexus); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveSwarmNexusInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listSwarmNexusesSubCmd = &cobra.Command{
	Use:   "list-nexuses",
	Short: "list nexuses",
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
			if err = tui.Send(cmd, args, action.ListSwarmNexuses); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListSwarmNexusesInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var describeSwarmNexusSubCmd = &cobra.Command{
	Use:   "describe-nexus",
	Short: "describe a nexus",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			if len(args) == 0 {
				fmt.Println("Error: no nexus id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.DescribeSwarmNexus); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeSwarmNexusInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var createSwarmNodeSubCmd = &cobra.Command{
	Use:   "create-node",
	Short: "create a new node",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("node-name")
			cmd.MarkFlagRequired("nexus-id")
			cmd.MarkFlagRequired("provider-id")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.CreateSwarmNode); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateSwarmNodeInteractive(cmd); err != nil {

				utils.PrintError(err)
			}
		}
	},
}

var describeSwarmNodeSubCmd = &cobra.Command{
	Use:   "describe-node",
	Short: "describe a node",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			if len(args) == 0 {
				fmt.Println("Error: no node id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.DescribeSwarmNode); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeSwarmNodeInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editSwarmNodeSubCmd = &cobra.Command{
	Use:   "edit-node",
	Short: "edit a node",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			if len(args) == 0 {
				fmt.Println("Error: no new name argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.EditSwarmNode); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditSwarmNodeInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var deleteSwarmNodeSubCmd = &cobra.Command{
	Use:   "delete-node",
	Short: "delete a node",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			if len(args) == 0 {
				fmt.Println("Error: no node id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.DeleteSwarmNode); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DeleteSwarmNodeInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listSwarmNodesSubCmd = &cobra.Command{
	Use:   "list-nodes",
	Short: "list nodes",
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
			if err = tui.Send(cmd, args, action.ListSwarmNodes); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListSwarmNodesInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var CreateSwarmRedundancyClassSubCmd = &cobra.Command{
	Use:   "create-redundancy-class",
	Short: "create a new redundancy class",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("redundancy-class-name")
			cmd.MarkFlagRequired("inner-k")
			cmd.MarkFlagRequired("inner-n")
			cmd.MarkFlagRequired("outer-k")
			cmd.MarkFlagRequired("outer-n")
			cmd.MarkFlagRequired("anti-affinity-group")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.CreateSwarmRedundancyClass); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateSwarmRedundancyClassInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var DescribeRedundancyClassesInteractiveSubCmd = &cobra.Command{
	Use:   "describe-redundancy-class",
	Short: "describe a redundancy class",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {

			if len(args) == 0 {
				fmt.Println("Error: no redundancy class id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.DescribeSwarmRedundancyClass); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeSwarmRedundancyClassInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listRedundancyClassesSubCmd = &cobra.Command{
	Use:   "list-redundancy-classes",
	Short: "list redundancy classes",
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
			if err = tui.Send(cmd, args, action.ListSwarmRedundancyClasses); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListSwarmRedundancyClassesInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var createSwarmRingSubCmd = &cobra.Command{

	Use:   "create-ring",
	Short: "create a new ring",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("nexus-ids")
			cmd.MarkFlagRequired("redundancy-class-id")

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.SetupSwarmRing); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.SetupSwarmRingInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listSwarmRingsSubCmd = &cobra.Command{
	Use:   "list-rings",
	Short: "list rings",
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
			if err = tui.Send(cmd, args, action.ListSwarmRings); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListSwarmRingInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {
	swarmCmd.AddCommand(createSwarmSubCmd)
	createSwarmSubCmd.Flags().String("name", "", "Name of the swarm")
	createSwarmSubCmd.Flags().String("description", "", "Description of the swarm")
	createSwarmSubCmd.Flags().String("configuration", "", "A Json object containing the swarm configuration")

	swarmCmd.AddCommand(listSwarmSubCmd)
	listSwarmSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for swarms")
	listSwarmSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different swarms")

	swarmCmd.AddCommand(describeSwarmSubCmd)
	describeSwarmSubCmd.Flags().String("format", "default", "Format of the output")

	swarmCmd.AddCommand(editSwarmDescriptionSubCmd)

	swarmCmd.AddCommand(editSwarmNameSubCmd)

	swarmCmd.AddCommand(addOperatorToSwarmSubCmd)
	addOperatorToSwarmSubCmd.Flags().String("email", "", "Email of the operator")
	addOperatorToSwarmSubCmd.Flags().String("role", "", "Role of the operator")
	addOperatorToSwarmSubCmd.Flags().String("first-name", "", "First name od the operator")
	addOperatorToSwarmSubCmd.Flags().String("last-name", "", "Last name of the operator")

	swarmCmd.AddCommand(listSwarmOperatorsSubCmd)
	listSwarmOperatorsSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for operators")
	listSwarmOperatorsSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different operators")

	swarmCmd.AddCommand(removeSwarmOperatorSubCmd)

	swarmCmd.AddCommand(removeSwarmSubCmd)
	removeSwarmSubCmd.Flags().String("email", "", "Email address")
	removeSwarmSubCmd.Flags().String("password", "", "Password")
	removeSwarmSubCmd.Flags().String("code", "", "Two factor authentication code")

	swarmCmd.AddCommand(describeSwarmOperatorsSubCmd)
	describeSwarmOperatorsSubCmd.Flags().String("format", "default", "Format of the output")

	swarmCmd.AddCommand(EditSwarmOperatorRoleSubCmd)
	EditSwarmOperatorRoleSubCmd.Flags().String("role", "", "Role of the operator")

	swarmCmd.AddCommand(createSwarmNexusSubCmd)
	createSwarmNexusSubCmd.Flags().String("nexus-name", "", "Name of the nexus")
	createSwarmNexusSubCmd.Flags().String("location", "", "Location of the nexus")
	createSwarmNexusSubCmd.Flags().String("description", "", "Description of the nexus")

	swarmCmd.AddCommand(editSwarmNexusSubCmd)
	editSwarmNexusSubCmd.Flags().String("nexus-name", "", "Name of the nexus")
	editSwarmNexusSubCmd.Flags().String("description", "", "Description of the nexus")

	swarmCmd.AddCommand(removeSwarmNexusSubCmd)

	swarmCmd.AddCommand(listSwarmNexusesSubCmd)
	listSwarmNexusesSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for nexuses")
	listSwarmNexusesSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different nexuses")

	swarmCmd.AddCommand(describeSwarmNexusSubCmd)
	describeSwarmNexusSubCmd.Flags().String("format", "default", "Format of the output")

	swarmCmd.AddCommand(createSwarmNodeSubCmd)
	createSwarmNodeSubCmd.Flags().String("node-name", "", "Name of the node")
	createSwarmNodeSubCmd.Flags().String("nexus-id", "", "ID of the nexus")
	createSwarmNodeSubCmd.Flags().String("description", "", "Description of the node")
	createSwarmNodeSubCmd.Flags().String("provider-id", "", "ID of the provider")

	swarmCmd.AddCommand(describeSwarmNodeSubCmd)
	describeSwarmNodeSubCmd.Flags().String("format", "default", "Format of the output")

	swarmCmd.AddCommand(editSwarmNodeSubCmd)
	editSwarmNodeSubCmd.Flags().String("node-name", "", "Name of the node")
	editSwarmNodeSubCmd.Flags().String("description", "", "Description of the node")

	swarmCmd.AddCommand(deleteSwarmNodeSubCmd)

	swarmCmd.AddCommand(listSwarmNodesSubCmd)
	listSwarmNodesSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for nodes")
	listSwarmNodesSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different nodes")
	listSwarmNodesSubCmd.Flags().String("nexus-id", "", "ID of the nexus")

	swarmCmd.AddCommand(CreateSwarmRedundancyClassSubCmd)
	CreateSwarmRedundancyClassSubCmd.Flags().String("redundancy-class-name", "", "Name of the redundancy class")
	CreateSwarmRedundancyClassSubCmd.Flags().String("redundancy-class-description", "", "Description of the redundancy class")
	CreateSwarmRedundancyClassSubCmd.Flags().Int("inner-n", 1, "Inner N")
	CreateSwarmRedundancyClassSubCmd.Flags().Int("inner-k", 0, "Inner K")
	CreateSwarmRedundancyClassSubCmd.Flags().Int("outer-n", 1, "Outer N")
	CreateSwarmRedundancyClassSubCmd.Flags().Int("outer-k", 0, "Outer K")
	CreateSwarmRedundancyClassSubCmd.Flags().Int("anti-affinity-group", 1, "Anti affinity group")

	swarmCmd.AddCommand(DescribeRedundancyClassesInteractiveSubCmd)
	DescribeRedundancyClassesInteractiveSubCmd.Flags().String("format", "default", "Format of the output")

	swarmCmd.AddCommand(listRedundancyClassesSubCmd)
	listRedundancyClassesSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for redundancy classes")
	listRedundancyClassesSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different redundancy classes")

	swarmCmd.AddCommand(createSwarmRingSubCmd)
	createSwarmRingSubCmd.Flags().StringSlice("nexus-ids", []string{}, "Comma separated list of nexus ids")
	createSwarmRingSubCmd.Flags().String("redundancy-class-id", "", "ID of the redundancy class")
	createSwarmRingSubCmd.Flags().Int("rings-number", 0, "Number of rings to create")

	swarmCmd.AddCommand(listSwarmRingsSubCmd)
	listSwarmRingsSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for rings")
	listSwarmRingsSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different rings")
	listSwarmRingsSubCmd.Flags().String("redundancy-class-id", "", "ID of the redundancy class")

	rootCmd.AddCommand(swarmCmd)
	swarmCmd.PersistentFlags().String("name", "", "Name of the swarm")
	swarmCmd.PersistentFlags().String("id", "", "ID of the swarm")
}
