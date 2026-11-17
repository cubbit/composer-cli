package gateway

import "github.com/spf13/cobra"

type GatewayServiceMock struct {
	CreateFunc             func(cmd *cobra.Command, args []string) error
	DescribeFunc           func(cmd *cobra.Command, args []string) error
	ListFunc               func(cmd *cobra.Command, args []string) error
	UpdateFunc             func(cmd *cobra.Command, args []string) error
	ListStorageClassesFunc func(cmd *cobra.Command, args []string) error
	CreateStorageClassFunc func(cmd *cobra.Command, args []string) error
	DeleteStorageClassFunc func(cmd *cobra.Command, args []string) error
	ListConnectionsFunc    func(cmd *cobra.Command, args []string) error
	CreateConnectionsFunc  func(cmd *cobra.Command, args []string) error
	DeleteConnectionFunc   func(cmd *cobra.Command, args []string) error
	ListRoutingRulesFunc   func(cmd *cobra.Command, args []string) error
	GetRoutingRuleFunc     func(cmd *cobra.Command, args []string) error
	CreateRoutingRuleFunc  func(cmd *cobra.Command, args []string) error
	UpdateRoutingRuleFunc  func(cmd *cobra.Command, args []string) error
	DeleteRoutingRuleFunc  func(cmd *cobra.Command, args []string) error
}

func NewGatewayServiceMock() *GatewayServiceMock {
	return &GatewayServiceMock{
		CreateFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		DescribeFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ListFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		UpdateFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ListStorageClassesFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		CreateStorageClassFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		DeleteStorageClassFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ListConnectionsFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		CreateConnectionsFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		DeleteConnectionFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ListRoutingRulesFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		GetRoutingRuleFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		CreateRoutingRuleFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		UpdateRoutingRuleFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		DeleteRoutingRuleFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (m *GatewayServiceMock) Create(cmd *cobra.Command, args []string) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) Describe(cmd *cobra.Command, args []string) error {
	if m.DescribeFunc != nil {
		return m.DescribeFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) List(cmd *cobra.Command, args []string) error {
	if m.ListFunc != nil {
		return m.ListFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) Update(cmd *cobra.Command, args []string) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) ListStorageClasses(cmd *cobra.Command, args []string) error {
	if m.ListStorageClassesFunc != nil {
		return m.ListStorageClassesFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) CreateStorageClass(cmd *cobra.Command, args []string) error {
	if m.CreateStorageClassFunc != nil {
		return m.CreateStorageClassFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) DeleteStorageClass(cmd *cobra.Command, args []string) error {
	if m.DeleteStorageClassFunc != nil {
		return m.DeleteStorageClassFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) ListConnections(cmd *cobra.Command, args []string) error {
	if m.ListConnectionsFunc != nil {
		return m.ListConnectionsFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) CreateConnections(cmd *cobra.Command, args []string) error {
	if m.CreateConnectionsFunc != nil {
		return m.CreateConnectionsFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) DeleteConnection(cmd *cobra.Command, args []string) error {
	if m.DeleteConnectionFunc != nil {
		return m.DeleteConnectionFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) ListRoutingRules(cmd *cobra.Command, args []string) error {
	if m.ListRoutingRulesFunc != nil {
		return m.ListRoutingRulesFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) GetRoutingRule(cmd *cobra.Command, args []string) error {
	if m.GetRoutingRuleFunc != nil {
		return m.GetRoutingRuleFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) CreateRoutingRule(cmd *cobra.Command, args []string) error {
	if m.CreateRoutingRuleFunc != nil {
		return m.CreateRoutingRuleFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) UpdateRoutingRule(cmd *cobra.Command, args []string) error {
	if m.UpdateRoutingRuleFunc != nil {
		return m.UpdateRoutingRuleFunc(cmd, args)
	}
	return nil
}

func (m *GatewayServiceMock) DeleteRoutingRule(cmd *cobra.Command, args []string) error {
	if m.DeleteRoutingRuleFunc != nil {
		return m.DeleteRoutingRuleFunc(cmd, args)
	}
	return nil
}
