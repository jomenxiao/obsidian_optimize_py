package types

import "github.com/ThinkInAIXYZ/go-mcp/protocol"

type ToolMeta struct {
	Name        string
	Description string
	InputSchema interface{}
}

type ResourceMeta struct {
	Name        string
	Description string
	Prompts     *protocol.Prompt
}
