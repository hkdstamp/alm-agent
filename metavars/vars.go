package metavars

// meta packege should not import others to avoid circular reference.

var (
	// ServerID is identifier of VM around provider.
	// such as Instance ID.
	ServerID string

	// ReportDisabled  true => do not send error report to rollbar
	ReportDisabled bool

	// AgentID is unique identifer of alm-agent
	AgentID string
)
