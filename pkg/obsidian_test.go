package pkg

import (
	"testing"

	"obsidianOptimizeMCP/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newMockClient() OfileClient {
	cfg := &types.Config{
		ObsidianURL:   "http://127.0.0.1:27123/",
		ObsidianToken: "ebbd2b2eee2ad6b5bbbd56eef3da035c7d4fa39d8844cef566b12103a5cdf32b",
	}
	return NewOfileClient(cfg)
}

func TestNewOfileClient(t *testing.T) {
	client := newMockClient()
	assert.NotNil(t, client)

}

func TestOfile_ReadFile(t *testing.T) {
	client := newMockClient()
	_, err := client.ReadFile("LLM/code-prompt/ObsidianOptimize.md")
	require.NoError(t, err)
}

func TestOfile_CreateOrUpdateFile(t *testing.T) {
	client := newMockClient()
	err := client.CreateOrUpdateFile("test.md", []byte("test"))
	require.NoError(t, err)
}

func TestOfile_DeleteFile(t *testing.T) {
	client := newMockClient()
	err := client.DeleteFile("test.md")
	require.NoError(t, err)
}

func TestOfile_ListFiles(t *testing.T) {
	client := newMockClient()
	files, err := client.ListFiles("LLM")
	require.NoError(t, err)
	assert.NotEmpty(t, files)
}
