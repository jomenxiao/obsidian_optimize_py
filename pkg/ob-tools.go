package pkg

import (
	"encoding/json"
	"obsidianOptimizeMCP/types"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	log "github.com/sirupsen/logrus"
)

// ToolObsidianCreate implements Tool interface for creating a file
type ToolObsidianCreate struct {
	OfileClient OfileClient
}

func (t *ToolObsidianCreate) Profile(config *types.Config) *types.ToolMeta {
	t.OfileClient = NewOfileClient(config)
	return &types.ToolMeta{
		Name:        "obsidian_create_file",
		Description: "Create or overwrite a file in Obsidian vault",
		InputSchema: types.ObsidianFileRequest{},
	}
}

func (t *ToolObsidianCreate) Call(req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var r types.ObsidianFileRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &r); err != nil {
		return nil, err
	}
	log.Debugf("Creating or updating file: %s content: %s", r.Path, r.Content)
	err := t.OfileClient.CreateOrUpdateFile(r.Path, []byte(r.Content))
	if err != nil {
		return nil, err
	}
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{Type: "text", Text: "File created/updated successfully"},
		},
	}, nil
}

// ToolObsidianRead implements Tool interface for reading a file
type ToolObsidianRead struct {
	OfileClient OfileClient
}

func (t *ToolObsidianRead) Profile(config *types.Config) *types.ToolMeta {
	t.OfileClient = NewOfileClient(config)
	return &types.ToolMeta{
		Name:        "obsidian_read_file",
		Description: "Read a file from Obsidian vault",
		InputSchema: types.ObsidianFileRequest{},
	}
}

func (t *ToolObsidianRead) Call(req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var r types.ObsidianFileRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &r); err != nil {
		return nil, err
	}
	data, err := t.OfileClient.ReadFile(r.Path)
	if err != nil {
		return nil, err
	}
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{Type: "text", Text: string(data)},
		},
	}, nil
}

// ToolObsidianDelete implements Tool interface for deleting a file
type ToolObsidianDelete struct {
	OfileClient OfileClient
}

func (t *ToolObsidianDelete) Profile(config *types.Config) *types.ToolMeta {
	t.OfileClient = NewOfileClient(config)
	return &types.ToolMeta{
		Name:        "obsidian_delete_file",
		Description: "Delete a file from Obsidian vault",
		InputSchema: types.ObsidianFileRequest{},
	}
}

func (t *ToolObsidianDelete) Call(req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var r types.ObsidianFileRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &r); err != nil {
		return nil, err
	}

	log.Debugf("Deleting file: %s", r.Path)
	err := t.OfileClient.DeleteFile(r.Path)
	if err != nil {
		return nil, err
	}
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{Type: "text", Text: "File deleted successfully"},
		},
	}, nil
}

// ToolObsidianList implements Tool interface for listing files in a path
type ToolObsidianList struct {
	OfileClient OfileClient
}

func (t *ToolObsidianList) Profile(config *types.Config) *types.ToolMeta {
	t.OfileClient = NewOfileClient(config)
	return &types.ToolMeta{
		Name:        "obsidian_list_files",
		Description: "List files in a path of Obsidian vault",
		InputSchema: types.ObsidianFileListRequest{},
	}
}

func (t *ToolObsidianList) Call(req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var r types.ObsidianFileListRequest
	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &r); err != nil {
		return nil, err
	}
	log.Debugf("Listing files in path: %s", r.Path)
	files, err := t.OfileClient.ListFiles(r.Path)
	if err != nil {
		return nil, err
	}
	b, _ := json.Marshal(files)
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{Type: "text", Text: string(b)},
		},
	}, nil
}

func init() {
	Tools = append(Tools, &ToolObsidianCreate{})
	Tools = append(Tools, &ToolObsidianRead{})
	Tools = append(Tools, &ToolObsidianDelete{})
	Tools = append(Tools, &ToolObsidianList{})
}
