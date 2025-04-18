package pkg

import (
	"obsidianOptimizeMCP/types"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/mitchellh/mapstructure"
)

type PromptObsidianMdOptimize struct {
	OfileClient OfileClient
}

func (p *PromptObsidianMdOptimize) Profile(config *types.Config) *types.ResourceMeta {
	p.OfileClient = NewOfileClient(config)
	return &types.ResourceMeta{
		Name:        "obsidian_md_optimize",
		Description: "Optimize markdown content",
		Prompts: &protocol.Prompt{
			Name:        "obsidian_md_optimize",
			Description: "Optimize markdown content",
			Arguments: []protocol.PromptArgument{
				{
					Name:        "path",
					Description: "Path of the markdown file",
					Required:    true,
				},
			},
		},
	}
}

func (p *PromptObsidianMdOptimize) Call(req *protocol.GetPromptRequest) (*protocol.GetPromptResult, error) {
	var r types.ObsidianMdOptimizeRequest
	if err := mapstructure.Decode(req.Arguments, &r); err != nil {
		return nil, err
	}
	data, err := p.OfileClient.ReadFile(r.Path)
	if err != nil {
		return nil, err
	}
	messages := []protocol.PromptMessage{
		{
			Role:    protocol.RoleAssistant,
			Content: protocol.TextContent{Type: "text", Text: types.PromptObsidianMdOptimize},
		},
		{
			Role:    protocol.RoleUser,
			Content: protocol.TextContent{Type: "text", Text: string(data)},
		},
	}
	return &protocol.GetPromptResult{
		Messages: messages,
	}, nil
}

func init() {
	Prompts = append(Prompts, &PromptObsidianMdOptimize{})
}
