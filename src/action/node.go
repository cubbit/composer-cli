// Package action provides CLI actions for managing nodes.
package action

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateNode(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, name, label, privateIP, publicIP string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var nodes *api.NewNodesResponse

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if privateIP, err = cmd.Flags().GetString("private-ip"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if publicIP, err = cmd.Flags().GetString("public-ip"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if label, err = cmd.Flags().GetString("label"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	nodesBodyRequest := api.BulkInsertNewNodeRequestBody{
		Nodes: []api.CreateNewNodeRequestBody{
			{
				Name:      name,
				PrivateIP: privateIP,
				PublicIP:  publicIP,
				Label:     &label,
			},
		},
	}

	if nodes, err = api.CreateNodes(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodesBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		nodes.Nodes,
		func(n *api.NewNodeResponseItem) []string {
			return []string{
				n.ID,
			}
		},
		&utils.SmartOutputConfig[*api.NewNodeResponseItem]{
			SingleResourceCompactOutput: true,
			SingleResource:              true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func CreateNodeBatch(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, filePath string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filePath, err = cmd.Flags().GetString("file"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading batch file: %w", err)
	}

	var batchNodes api.BulkInsertNewNodeRequestBody
	if err = json.Unmarshal(fileData, &batchNodes); err != nil {
		schemaTemplate := map[string]interface{}{
			"nodes": []map[string]interface{}{
				{
					"name":       "",
					"label":      "",
					"config":     map[string]interface{}{},
					"private_ip": "",
					"public_ip":  "",

					"agents": []map[string]interface{}{
						{
							"port":     0,
							"features": map[string]interface{}{},
							"volume": map[string]interface{}{
								"mount_point": "",
								"disk":        "",
							},
						},
					},
				},
			},
		}

		return fmt.Errorf("error parsing batch file: %w\nExpected file schema:\n%s", err, getStructSchema(schemaTemplate))
	}

	if len(batchNodes.Nodes) == 0 {
		return fmt.Errorf("no nodes found in batch file")
	}

	for i, node := range batchNodes.Nodes {
		if node.Name == "" {
			return fmt.Errorf("node at index %d is missing a name", i)
		}
		if node.PrivateIP == "" || !utils.IsValidIP(node.PrivateIP) {
			return fmt.Errorf("node '%s' has an invalid private IP address", node.Name)
		}
		if node.PublicIP == "" || !utils.IsValidIP(node.PublicIP) {
			return fmt.Errorf("node '%s' has an invalid public IP address", node.Name)
		}
	}

	nodes, err := api.CreateNodes(*urls, resolvedProfile.APIKey, swarmID, nexusID, batchNodes)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		nodes.Nodes,
		func(n *api.NewNodeResponseItem) []string {
			return []string{
				n.ID,
			}
		},
		&utils.SmartOutputConfig[*api.NewNodeResponseItem]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func DescribeNode(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, nodeID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var node *api.NewNode

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if node, err = api.GetNode(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.NewNode{node},
		func(n *api.NewNode) []string {
			return []string{
				n.ID,
				n.Name,
				n.PrivateIP,
				n.PublicIP,
			}
		},
		&utils.SmartOutputConfig[*api.NewNode]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		})

}

func EditNode(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, nodeID, name, label, privateIP, publicIP string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var nodeBodyRequest api.UpdateNewNodeRequestBody

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if privateIP, err = cmd.Flags().GetString("private-ip"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if publicIP, err = cmd.Flags().GetString("public-ip"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if label, err = cmd.Flags().GetString("label"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if name != "" {
		nodeBodyRequest.Name = &name
	}

	if privateIP != "" {
		nodeBodyRequest.PrivateIP = &privateIP
	}

	if publicIP != "" {
		nodeBodyRequest.PublicIP = &publicIP
	}

	if label != "" {
		nodeBodyRequest.Label = &label
	}

	if err = api.UpdateNode(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID, nodeBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingNodeRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.UpdateNewNodeRequestBody{&nodeBodyRequest},
		func(n *api.UpdateNewNodeRequestBody) []string {
			return []string{
				nodeID,
			}
		},
		&utils.SmartOutputConfig[*api.UpdateNewNodeRequestBody]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func RemoveNode(cmd *cobra.Command, args []string) error {
	var err error
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var swarmID, nexusID, nodeID string

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.DeleteNode(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingNodeRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{nodeID},
		func(n string) []string {
			return []string{n}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func ListNodes(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, sort, filter string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var nodes *api.GenericPaginatedResponse[*api.NewNode]

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if sort, err = cmd.Flags().GetString("sort"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filter, err = cmd.Flags().GetString("filter"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filter != "" {
		filter = utils.BuildFilterQuery(filter)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if nodes, err = api.ListNodes(*urls, resolvedProfile.APIKey, swarmID, nexusID, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		nodes.Data,
		func(n *api.NewNode) []string {
			return []string{
				n.ID,
			}
		},
		&utils.SmartOutputConfig[*api.NewNode]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func GenerateNodeDeployFiles(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, nodeID, outputDir string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var nodeConfigs []api.NodeConfig
	var node *api.NewNode

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if outputDir, err = cmd.Flags().GetString("output-dir"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if node, err = api.GetNode(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	if node == nil {
		utils.PrintNotFound("Node not found")
		return nil
	}

	var agents *api.GenericPaginatedResponse[*api.NewAgent]
	if agents, err = api.ListAgents(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingAgentsRequest, err)
	}

	if len(agents.Data) == 0 {
		utils.PrintNotFound("No agents found for the specified node")
		return nil
	}

	nodeConfigs = make([]api.NodeConfig, 0, len(agents.Data))
	for _, agent := range agents.Data {
		nodeConfig := api.NodeConfig{
			ID:        node.ID,
			Name:      node.Name,
			PublicIP:  node.PublicIP,
			PrivateIP: node.PrivateIP,
			Agents: []api.AgentConfig{
				{
					ID:         agent.ID,
					MountPoint: agent.Volume.MountPoint,
					Disk:       agent.Volume.Disk,
					Port:       agent.Port,
					Secret:     agent.Secret,
				},
			},
		}
		nodeConfigs = append(nodeConfigs, nodeConfig)
	}

	if err = generateDeployFiles(urls, "ansible", nodeConfigs, outputDir); err != nil {
		return fmt.Errorf("error generating deployment files: %w", err)
	}

	utils.PrintSuccess("Deployment files generated successfully in directory: " + outputDir)

	return nil
}

func getStructSchema(schemaTemplate map[string]interface{}) string {
	bytes, err := json.MarshalIndent(schemaTemplate, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error generating schema: %v", err)
	}

	return string(bytes)
}

func generateInventoryContent(nodeConfigs []api.NodeConfig) string {
	var content strings.Builder
	for _, nodeConfig := range nodeConfigs {
		content.WriteString(fmt.Sprintf("[%s]\n%s\n", nodeConfig.Name, nodeConfig.PublicIP))
	}
	return content.String()
}

func generateHostNamesContent(nodeConfigs []api.NodeConfig) string {
	var content strings.Builder
	for _, nodeConfig := range nodeConfigs {
		content.WriteString(fmt.Sprintf("%s\n", nodeConfig.Name))
	}
	return content.String()
}

func generateAgentSecretsContent(nodeConfig api.NodeConfig) map[string]interface{} {
	agentsMap := make(map[string]interface{})

	for i, agent := range nodeConfig.Agents {
		agentsMap[strconv.Itoa(i)] = map[string]interface{}{
			"agent_secret":         agent.Secret,
			"mount_point":          agent.MountPoint,
			"disk":                 agent.Disk,
			"cccp_server_port":     agent.Port,
			"cccp_server_local_ip": nodeConfig.PrivateIP,
			"machine_id":           nodeConfig.ID,
		}
	}

	return agentsMap
}

func downloadAndGenerateAnsibleTar(tarURL string, nodeConfigs []api.NodeConfig, outputPath string) error {
	resp, err := http.Get(tarURL)
	if err != nil {
		return fmt.Errorf("failed to download tarball: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download tarball: %s", resp.Status)
	}

	tmpDir, err := os.MkdirTemp("", "ansible-tar")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	if err := utils.ExtractTar(resp.Body, tmpDir); err != nil {
		return fmt.Errorf("failed to extract tarball: %w", err)
	}

	if err := injectAnsibleFiles(tmpDir+"/cubbit-agent-playbook", nodeConfigs); err != nil {
		return fmt.Errorf("failed to inject generated files: %w", err)
	}

	if err := utils.CreateTar(outputPath, tmpDir); err != nil {
		return fmt.Errorf("failed to repack tarball: %w", err)
	}

	return nil
}

func injectAnsibleFiles(root string, nodes []api.NodeConfig) error {
	filesDir := filepath.Join(root, "files")
	invDir := filepath.Join(root, "inventory")

	if err := os.MkdirAll(filesDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(invDir, 0755); err != nil {
		return err
	}

	invPath := filepath.Join(invDir, "hosts.ini")
	if err := os.WriteFile(invPath, []byte(generateInventoryContent(nodes)), 0644); err != nil {
		return err
	}

	hostNamesPath := filepath.Join(filesDir, "host-names")
	if err := os.WriteFile(hostNamesPath, []byte(generateHostNamesContent(nodes)), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(filesDir, "ssh-public-keys"), []byte(""), 0644); err != nil {
		return err
	}

	for _, node := range nodes {
		path := filepath.Join(filesDir, fmt.Sprintf("%s-agent-secrets.json", node.Name))
		data, _ := json.MarshalIndent(generateAgentSecretsContent(node), "", "  ")
		if err := os.WriteFile(path, data, 0644); err != nil {
			return err
		}
	}

	return nil
}

func generateDeployFiles(conf *configuration.URLs, deployOption string, nodeConfigs []api.NodeConfig, outputDir string) error {
	info, err := os.Stat(outputDir)
	if os.IsNotExist(err) {
		utils.PrintInfo(fmt.Sprintf("Output directory %s doesn't exist, creating it...", outputDir))
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to check output directory: %w", err)
	} else if !info.IsDir() {
		return fmt.Errorf("output path %s exists but is not a directory", outputDir)
	}

	basePath := filepath.Join(outputDir, "cubbit-agent-playbook.tar")

	switch deployOption {
	case "ansible":
		return downloadAndGenerateAnsibleTar(constants.AnsibleTarURL, nodeConfigs, basePath)
	default:
		return fmt.Errorf("invalid deployment option: %s", deployOption)
	}
}
