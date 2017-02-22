package pbjson

// I18nField ...
type I18nField map[string]string

// WatchedDirectory ...
type WatchedDirectory struct {
	ChainID  string `json:"chain_id"`
	OnlyDirs bool   `json:"only_dirs"`
	Path     string `json:"path"`
	UnitType string `json:"unit_type"`
}

// Chain ...
type Chain struct {
	Description *I18nField `json:"description"`
	LinkID      string     `json:"link_id"`
}

// Link ...
type Link struct {
	Config            *LinkConfig           `json:"config"`
	Description       *I18nField            `json:"description"`
	ExitCodes         map[int]*LinkExitCode `json:"exit_codes"`
	FallbackJobStatus string                `json:"fallback_job_status"`
	FallbackLinkID    string                `json:"fallback_link_id"`
	Group             *I18nField            `json:"group"`
}

// LinkConfig ...
type LinkConfig struct {
	Manager string `json:"@manager"`
	Model   string `json:"@model"`

	// StandardTaskConfig
	Arguments          string `json:"arguments"`
	Execute            string `json:"execute"`
	FilterFileEnd      string `json:"filter_file_end"`
	FilterSubdir       string `json:"filter_subdir"`
	RequiresOutputLock bool   `json:"requires_output_lock"`
	StderrFile         string `json:"stderr_file"`
	StdoutFile         string `json:"stdout_file"`

	// MicroServiceChainChoice
	Choices []string `json:"chain_choices"`

	// TaskConfigSetUnitVariable
	Variable      string `json:"variable"`
	VariableValue string `json:"variable_value"`
	ChainID       string `json:"chain_id"`

	// MicroServiceChoiceReplacementDic
	Replacements []*LinkConfigReplacement `json:"replacements"`

	// TaskConfigUnitVariableLinkPull
	// @ Variable, ChainID

	// TaskConfigAssignMagicLink
	LinkID string `json:"link_id"`
}

// LinkConfigReplacement ...
type LinkConfigReplacement struct {
	Id          string            `json:"id"`
	Description *I18nField        `json:"description"`
	Items       map[string]string `json:"items"`
}

// LinkExitCode ...
type LinkExitCode struct {
	JobStatus string `json:"job_status"`
	LinkID    string `json:"link_id"`
}

// WorkflowData is our proxy type to marshal WorkflowData or to
// unmarshal JSON-encoded data into a
type WorkflowData struct {
	WatchedDirectories []*WatchedDirectory `json:"watched_directories"`
	Chains             map[string]*Chain   `json:"chains"`
	Links              map[string]*Link    `json:"links"`
}

// NewWorkflowData returns a pointer to a new WorkflowData
func NewWorkflowData() *WorkflowData {
	return &WorkflowData{
		WatchedDirectories: []*WatchedDirectory{},
		Chains:             map[string]*Chain{},
		Links:              map[string]*Link{},
	}
}
