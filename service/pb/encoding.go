package pb

import (
	"encoding/json"
	"fmt"

	"strings"

	"github.com/artefactual-labs/archivematica-workflow/service/pbjson"
)

// UnmarshalJSON implements Unmarshaler. This method updates the current
// WorkflowData message with the information found in the JSON blob.
func (w *WorkflowData) UnmarshalJSON(data []byte) error {
	jwd := pbjson.NewWorkflowData()

	var err error
	if err = json.Unmarshal(data, jwd); err != nil {
		return err
	}

	for _, item := range jwd.WatchedDirectories {
		wd := WatchedDirectory{
			ChainId:  item.ChainID,
			OnlyDirs: item.OnlyDirs,
			Path:     item.Path,
			Type:     WatchedDirectory_WatchedDirectoryType(WatchedDirectory_WatchedDirectoryType_value[strings.ToUpper(item.UnitType)]),
		}
		w.WatchedDirectories = append(w.WatchedDirectories, &wd)
	}

	w.Chains = make(map[string]*Chain)
	for key, value := range jwd.Chains {
		c := Chain{
			Id:          key,
			LinkId:      value.LinkID,
			Description: i18NFromJSONtoProto(value.Description),
		}
		w.Chains[key] = &c
	}

	w.Links = make(map[string]*Link)
	for key, value := range jwd.Links {
		l := Link{
			Id:                key,
			Description:       i18NFromJSONtoProto(value.Description),
			Group:             i18NFromJSONtoProto(value.Group),
			FallbackLinkId:    value.FallbackLinkID,
			FallbackJobStatus: jobStatusFromJSONtoProto(value.FallbackJobStatus),
		}
		l.Config, err = linkConfigFromJSONToProto(value.Config)
		if err != nil {
			return err
		}
		for code, props := range value.ExitCodes {
			lec := Link_LinkExitCode{
				Code:      uint32(code),
				JobStatus: jobStatusFromJSONtoProto(props.JobStatus),
				LinkId:    props.LinkID,
			}
			l.ExitCodes = append(l.ExitCodes, &lec)
		}
		w.Links[key] = &l
	}

	return nil
}

// MarshalJSON implements Marshaler. This methos returns the current
// WorkflowData message encoded as JSON.
func (w *WorkflowData) MarshalJSON() ([]byte, error) {
	jwd := pbjson.NewWorkflowData()

	for _, item := range w.WatchedDirectories {
		wd := pbjson.WatchedDirectory{
			ChainID:  item.ChainId,
			OnlyDirs: item.OnlyDirs,
			Path:     item.Path,
			UnitType: item.Type.String(),
		}
		jwd.WatchedDirectories = append(jwd.WatchedDirectories, &wd)
	}

	for id, item := range w.Chains {
		jwd.Chains[id] = &pbjson.Chain{
			Description: i18NFromProtoToJSON(item.Description),
			LinkID:      item.LinkId,
		}
	}

	for id, item := range w.Links {
		jwd.Links[id] = &pbjson.Link{
			FallbackLinkID: item.FallbackLinkId,
		}
	}

	data, err := json.Marshal(jwd)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func i18NFromJSONtoProto(field *pbjson.I18nField) map[string]string {
	f := make(map[string]string)
	for key, value := range *field {
		f[key] = value
	}
	return f
}

func i18NFromProtoToJSON(field map[string]string) *pbjson.I18nField {
	f := pbjson.I18nField(field)
	return &f
}

func linkConfigFromJSONToProto(config *pbjson.LinkConfig) (*Link_LinkConfig, error) {
	l := Link_LinkConfig{
		Manager: config.Manager,
		Model:   config.Model,
	}
	if l.Model == "StandardTaskConfig" {
		attrs := &Link_LinkConfig_Standard_{}
		attrs.Standard = &Link_LinkConfig_Standard{
			Arguments:          config.Arguments,
			Execute:            config.Execute,
			FilterFileEnd:      config.FilterFileEnd,
			FilterFileStart:    config.FilterFileStart,
			FilterSubdir:       config.FilterSubdir,
			RequiresOutputLock: config.RequiresOutputLock,
			StderrFile:         config.StderrFile,
			StdoutFile:         config.StdoutFile,
		}
		l.Attrs = attrs
	} else if l.Model == "MicroServiceChainChoice" {
		attrs := &Link_LinkConfig_UserChoice_{}
		attrs.UserChoice = &Link_LinkConfig_UserChoice{
			ChainIds: config.Choices,
		}
		l.Attrs = attrs
	} else if l.Model == "TaskConfigSetUnitVariable" {
		attrs := &Link_LinkConfig_SetUnitVar_{}
		attrs.SetUnitVar = &Link_LinkConfig_SetUnitVar{
			Variable:      config.Variable,
			VariableValue: config.VariableValue,
			ChainId:       config.ChainID,
		}
		l.Attrs = attrs
	} else if l.Model == "MicroServiceChoiceReplacementDic" {
		attrs := &Link_LinkConfig_UserChoiceDict_{}
		attrs.UserChoiceDict = &Link_LinkConfig_UserChoiceDict{}
		for _, item := range config.Replacements {
			repl := Link_LinkConfig_UserChoiceDict_Replacement{
				Id:          item.Id,
				Description: i18NFromJSONtoProto(item.Description),
				Items:       item.Items,
			}
			attrs.UserChoiceDict.Replacements = append(attrs.UserChoiceDict.Replacements, &repl)
		}
		l.Attrs = attrs
	} else if l.Model == "TaskConfigUnitVariableLinkPull" {
		attrs := &Link_LinkConfig_GetUnitVar_{}
		attrs.GetUnitVar = &Link_LinkConfig_GetUnitVar{
			ChainId:  config.ChainID,
			Variable: config.Variable,
		}
		l.Attrs = attrs
	} else if l.Model == "TaskConfigAssignMagicLink" {
		attrs := &Link_LinkConfig_Magic_{}
		attrs.Magic = &Link_LinkConfig_Magic{
			LinkId: config.LinkID,
		}
		l.Attrs = attrs
	} else if l.Manager == "linkTaskManagerLoadMagicLink" {
		// We do nothing here, but we don't want to return an error.
	} else {
		return nil, fmt.Errorf("LinkConfigFromJSONToProto: unknown model=%v manager=%v", l.Model, l.Manager)
	}
	return &l, nil
}

var ss = map[string]Job_Status{
	"Unknown":                Job_UNKNOWN,
	"Awaiting decision":      Job_AWAITING_DECISION,
	"Completed successfully": Job_COMPLETED_SUCCESSFULLY,
	"Executing commands":     Job_EXECUTING_COMMANDS,
	"Failed":                 Job_FAILED,
}

func jobStatusFromJSONtoProto(jobStatus string) Job_Status {
	s, ok := ss[jobStatus]
	if !ok {
		return Job_UNKNOWN
	}
	return s
}
