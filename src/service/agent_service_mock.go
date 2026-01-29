package service

import "github.com/spf13/cobra"

type AgentServiceMock struct {
	CreateAgentFunc      func(cmd *cobra.Command, args []string) error
	CreateAgentBatchFunc func(cmd *cobra.Command, args []string) error
	DescribeAgentFunc    func(cmd *cobra.Command, args []string) error
	EditAgentFunc        func(cmd *cobra.Command, args []string) error
	ListAgentsFunc       func(cmd *cobra.Command, args []string) error
	RemoveAgentFunc      func(cmd *cobra.Command, args []string) error
	CheckAgentStatusFunc func(cmd *cobra.Command, args []string) error
}

// NewAgentServiceMock returns a mock with default outputs for all methods
func NewAgentServiceMock() *AgentServiceMock {
	return &AgentServiceMock{
		CreateAgentFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Agent created successfully")
			return nil
		},
		CreateAgentBatchFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Batch agents created successfully")
			return nil
		},
		DescribeAgentFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Agent description: {id: 1, status: running}")
			return nil
		},
		EditAgentFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Agent edited successfully")
			return nil
		},
		ListAgentsFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Agents list: [agent1, agent2, agent3]")
			return nil
		},
		RemoveAgentFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Agent removed successfully")
			return nil
		},
		CheckAgentStatusFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Agent status: healthy")
			return nil
		},
	}
}

func (m *AgentServiceMock) CreateAgent(cmd *cobra.Command, args []string) error {
	if m.CreateAgentFunc != nil {
		return m.CreateAgentFunc(cmd, args)
	}
	return nil
}

func (m *AgentServiceMock) CreateAgentBatch(cmd *cobra.Command, args []string) error {
	if m.CreateAgentBatchFunc != nil {
		return m.CreateAgentBatchFunc(cmd, args)
	}
	return nil
}

func (m *AgentServiceMock) DescribeAgent(cmd *cobra.Command, args []string) error {
	if m.DescribeAgentFunc != nil {
		return m.DescribeAgentFunc(cmd, args)
	}
	return nil
}

func (m *AgentServiceMock) EditAgent(cmd *cobra.Command, args []string) error {
	if m.EditAgentFunc != nil {
		return m.EditAgentFunc(cmd, args)
	}
	return nil
}

func (m *AgentServiceMock) ListAgents(cmd *cobra.Command, args []string) error {
	if m.ListAgentsFunc != nil {
		return m.ListAgentsFunc(cmd, args)
	}
	return nil
}

func (m *AgentServiceMock) RemoveAgent(cmd *cobra.Command, args []string) error {
	if m.RemoveAgentFunc != nil {
		return m.RemoveAgentFunc(cmd, args)
	}
	return nil
}

func (m *AgentServiceMock) CheckAgentStatus(cmd *cobra.Command, args []string) error {
	if m.CheckAgentStatusFunc != nil {
		return m.CheckAgentStatusFunc(cmd, args)
	}
	return nil
}
