package agent

import (
	"encoding/json"

	"github.com/openai/openai-go/v3"
)

type ToolDefinition struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	InputSchema openai.FunctionParameters `json:"input_schema"`
	Function    func(input json.RawMessage) (string, error)
}

var ReadHealthDataDefinition = ToolDefinition{
	Name:        "read_health_data",
	Description: "Read health data from health ",
	InputSchema: openai.FunctionParameters{},
	Function: func(input json.RawMessage) (string, error) {
		// Simulate fetching health stats
		return "", nil
	},
}
