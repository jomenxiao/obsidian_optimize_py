package pkg

import (
	"obsidianOptimizeMCP/types"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	log "github.com/sirupsen/logrus"
)

var Tools = []Tool{}

type Tool interface {
	Profile(config *types.Config) *types.ToolMeta
	Call(req *protocol.CallToolRequest) (*protocol.CallToolResult, error)
}

var Prompts = []Prompt{}

type Prompt interface {
	Profile(config *types.Config) *types.ResourceMeta
	Call(req *protocol.GetPromptRequest) (*protocol.GetPromptResult, error)
}

func Register(mcpServer *server.Server, config *types.Config) {
	log.Printf("Registering %d tools...", len(Tools))
	for _, t := range Tools {
		meta := t.Profile(config)
		// register tool with mcpServer
		tool, err := protocol.NewTool(meta.Name, meta.Description, meta.InputSchema)
		if err != nil {
			log.Fatalf("Failed to create tool: %v", err)
			return
		}
		mcpServer.RegisterTool(tool, t.Call)
		log.Printf("Registered tool: %s", meta.Name)

	}
	log.Printf("Registering %d prompts...", len(Prompts))
	for _, p := range Prompts {
		meta := p.Profile(config)
		mcpServer.RegisterPrompt(meta.Prompts, p.Call)
		log.Printf("Registered prompt: %s", meta.Name)
	}
}
