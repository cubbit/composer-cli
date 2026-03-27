package service

import (
	"fmt"
	"sort"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/cubbit/composer-cli/utils/printer/tree"
	"github.com/cubbit/composer-cli/utils/printer/utils"
	"github.com/spf13/cobra"
)

func PrintClusterDetails(cmd *cobra.Command, cluster api.InfraAggregateCluster) error {
	printFuncs := []func() error{
		func() error { return printClusterInfo(cmd, cluster) },
		func() error { return printPhysicalNodes(cmd, cluster.Details.Nodes) },
		func() error { return PrintVirtualNodes(cmd, cluster.Details.VirtualNodes) },
	}

	return printer.Compose(cmd, printFuncs...)
}

func printClusterInfo(cmd *cobra.Command, cluster api.InfraAggregateCluster) error {
	rowMapper := func(c api.InfraAggregateCluster) []string {
		return []string{
			c.ClusterID,
			c.Name,
			string(c.Type),
			fmt.Sprintf("%d", len(c.Details.Nodes)),
			fmt.Sprintf("%d", len(c.Details.VirtualNodes)),
			c.Details.LastUpdate.Format("2006-01-02 15:04:05"),
		}
	}

	tableData := []api.InfraAggregateCluster{cluster}
	tableColumns := []table.Column[api.InfraAggregateCluster]{
		{Title: "Cluster ID"},
		{Title: "Name"},
		{Title: "Type"},
		{Title: "Physical Nodes"},
		{Title: "Virtual Nodes"},
		{Title: "Last Update"},
	}

	infoText := "Cluster Information\n" + utils.Separator + "\n"

	return printer.CreateTable(cmd, tableData,
		table.WithColumns[api.InfraAggregateCluster](tableColumns),
		table.WithRowMapper[api.InfraAggregateCluster](rowMapper),
		table.WithShowHeader[api.InfraAggregateCluster](true),
		table.WithPrefix[api.InfraAggregateCluster](infoText),
		table.WithSuffix[api.InfraAggregateCluster]("\n"),
	)
}

func printPhysicalNodes(cmd *cobra.Command, nodes []api.InfraAggregateNodeDetail) error {
	if len(nodes) == 0 {
		return nil
	}

	nodesTitle := "\nPhysical Nodes\n" + utils.Separator + "\n"

	nodeNodes := make([]tree.TreeNode, len(nodes))
	for i, node := range nodes {
		nodeNodes[i] = buildPhysicalNodeTree(node)
	}

	return printer.PrintTree(cmd, nodeNodes,
		tree.WithPrefix(nodesTitle),
		tree.WithSuffix("\n"),
	)
}

func buildPhysicalNodeTree(node api.InfraAggregateNodeDetail) tree.TreeNode {
	nodeTitle := fmt.Sprintf("%s (%s)", node.NodeName, node.NodeID)

	children := []tree.TreeNode{
		buildNodeStatusTree(node.Status),
		buildNodeHardwareTree(node),
		buildNodeNetworkTree(node),
		buildNodeDisksTree(node),
	}

	return tree.TreeNode{
		Value:     nodeTitle,
		Children:  children,
		Formatter: defaultFormatter,
	}
}

func buildNodeStatusTree(status api.InfraAggregateStatus) tree.TreeNode {
	statusStr := fmt.Sprintf("Status: %s", status.Code)
	if status.Details != "" {
		statusStr += fmt.Sprintf(" - %s", status.Details)
	}

	return tree.TreeNode{
		Value:     statusStr,
		Formatter: defaultFormatter,
	}
}

func buildNodeHardwareTree(node api.InfraAggregateNodeDetail) tree.TreeNode {
	hardwareTitle := "Hardware"

	children := []tree.TreeNode{}

	if node.OSName != nil {
		children = append(children, tree.TreeNode{
			Value:     fmt.Sprintf("OS: %s", *node.OSName),
			Formatter: defaultFormatter,
		})
	}

	if node.CPU != nil {
		children = append(children, tree.TreeNode{
			Value:     fmt.Sprintf("CPU Cores: %d", node.CPU.Cores),
			Formatter: defaultFormatter,
		})
	}

	if node.RAM != nil {
		children = append(children, tree.TreeNode{
			Value:     fmt.Sprintf("RAM: %.2f GB", node.RAM.Available),
			Formatter: defaultFormatter,
		})
	}

	return tree.TreeNode{
		Value:     hardwareTitle,
		Children:  children,
		Formatter: defaultFormatter,
	}
}

func buildNodeNetworkTree(node api.InfraAggregateNodeDetail) tree.TreeNode {
	networkTitle := "Network"

	children := []tree.TreeNode{}

	if node.ExternalIP != nil {
		children = append(children, tree.TreeNode{
			Value:     fmt.Sprintf("External IP: %s", *node.ExternalIP),
			Formatter: defaultFormatter,
		})
	}

	if node.InternalIP != nil {
		children = append(children, tree.TreeNode{
			Value:     fmt.Sprintf("Internal IP: %s", *node.InternalIP),
			Formatter: defaultFormatter,
		})
	}

	return tree.TreeNode{
		Value:     networkTitle,
		Children:  children,
		Formatter: defaultFormatter,
	}
}

func buildNodeDisksTree(node api.InfraAggregateNodeDetail) tree.TreeNode {
	disksTitle := "Disks"

	if len(node.Disks) == 0 {
		return tree.TreeNode{
			Value:     disksTitle,
			Children:  []tree.TreeNode{},
			Formatter: defaultFormatter,
		}
	}

	children := make([]tree.TreeNode, len(node.Disks))
	for i, disk := range node.Disks {
		children[i] = buildDiskTree(disk)
	}

	return tree.TreeNode{
		Value:     disksTitle,
		Children:  children,
		Formatter: defaultFormatter,
	}
}

func buildDiskTree(disk api.InfraAggregateDiskDetail) tree.TreeNode {
	diskTitle := fmt.Sprintf("Disk: %s", disk.DiskUUID)

	children := []tree.TreeNode{
		{
			Value:     fmt.Sprintf("Path: %s", disk.Path),
			Formatter: defaultFormatter,
		},
		{
			Value:     fmt.Sprintf("Status: %s", disk.Status.Code),
			Formatter: defaultFormatter,
		},
		{
			Value:     fmt.Sprintf("Total: %s", utils.FormatBytes(disk.TotalStorageSizeBytes)),
			Formatter: defaultFormatter,
		},
		{
			Value:     fmt.Sprintf("Used: %s", utils.FormatBytes(disk.UsedStorageBytes)),
			Formatter: defaultFormatter,
		},
	}

	if disk.PVRef != "" {
		children = append(children, tree.TreeNode{
			Value:     fmt.Sprintf("PV Reference: %s", disk.PVRef),
			Formatter: defaultFormatter,
		})
	}

	return tree.TreeNode{
		Value:     diskTitle,
		Children:  children,
		Formatter: defaultFormatter,
	}
}

func PrintVirtualNodes(cmd *cobra.Command, virtualNodes []api.InfraAggregateVirtualNodeDetail) error {
	if len(virtualNodes) == 0 {
		return nil
	}

	virtualNodesTitle := "\nVirtual Nodes\n" + utils.Separator + "\n"

	nodeNodes := make([]tree.TreeNode, len(virtualNodes))
	for i, node := range virtualNodes {
		nodeNodes[i] = buildVirtualNodeTree(node)
	}

	return printer.PrintTree(cmd, nodeNodes,
		tree.WithPrefix(virtualNodesTitle),
		tree.WithSuffix("\n"),
	)
}

func buildVirtualNodeTree(node api.InfraAggregateVirtualNodeDetail) tree.TreeNode {
	nodeTitle := fmt.Sprintf("%s (%s)", node.NodeName, node.NodeID)

	children := []tree.TreeNode{
		buildVirtualNodeStatusTree(node.Status),
		buildVirtualNodeStorageTree(node),
	}

	return tree.TreeNode{
		Value:     nodeTitle,
		Children:  children,
		Formatter: defaultFormatter,
	}
}

func buildVirtualNodeStatusTree(status api.InfraAggregateStatus) tree.TreeNode {
	statusStr := fmt.Sprintf("Status: %s", status.Code)
	if status.Details != "" {
		statusStr += fmt.Sprintf(" - %s", status.Details)
	}

	return tree.TreeNode{
		Value:     statusStr,
		Formatter: defaultFormatter,
	}
}

func buildVirtualNodeStorageTree(node api.InfraAggregateVirtualNodeDetail) tree.TreeNode {
	storageTitle := fmt.Sprintf("Storage: %s", node.StorageType)

	children := []tree.TreeNode{}

	keys := make([]string, 0, len(node.StorageConfiguration))
	for key := range node.StorageConfiguration {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := node.StorageConfiguration[key]
		children = append(children, tree.TreeNode{
			Value:     fmt.Sprintf("%s: %v", key, value),
			Formatter: defaultFormatter,
		})
	}

	return tree.TreeNode{
		Value:     storageTitle,
		Children:  children,
		Formatter: defaultFormatter,
	}
}

func defaultFormatter(v any) string {
	if str, ok := v.(string); ok {
		return str
	}
	return fmt.Sprintf("%v", v)
}

func PrintClusters(cmd *cobra.Command, clusters []api.InfrastructureCluster) error {
	tableColumns := []table.Column[api.InfrastructureCluster]{
		{Title: "Cluster ID"},
		{Title: "Name"},
		{Title: "Type"},
	}

	rowMapper := func(v api.InfrastructureCluster) []string {
		return []string{
			v.ClusterID,
			v.Name,
			v.Type,
		}
	}

	return printer.CreateTable(
		cmd,
		clusters,
		table.WithColumns[api.InfrastructureCluster](tableColumns),
		table.WithRowMapper[api.InfrastructureCluster](rowMapper),
		table.WithShowHeader[api.InfrastructureCluster](true),
		table.WithSuffix[api.InfrastructureCluster]("\n"),
	)
}
