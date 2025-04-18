package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"obsidianOptimizeMCP/types"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type OfileClient interface {
	CreateOrUpdateFile(filepath string, content []byte) error
	ReadFile(filepath string) ([]byte, error)
	DeleteFile(filepath string) error
	ListFiles(folder string) ([]string, error)
}

type Ofile struct {
	BaseURL string
	Client  *http.Client
	Token   string // Bearer token for authentication
}

func NewOfileClient(config *types.Config) OfileClient {
	if !strings.HasPrefix(config.ObsidianURL, "http://") && !strings.HasPrefix(config.ObsidianURL, "https://") {
		config.ObsidianURL = "http://" + config.ObsidianURL
	}
	return &Ofile{
		BaseURL: strings.TrimRight(config.ObsidianURL, "/"),
		Client:  &http.Client{Timeout: 30 * time.Second},
		Token:   config.ObsidianToken,
	}
}

// unified HTTP request handler
func (c *Ofile) doRequest(method, url string, body io.Reader, heads []map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
	for _, head := range heads {
		for k, v := range head {
			req.Header.Set(k, v)
		}
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Debugf("Failed to %s %s: %v", method, url, err)
		return nil, err
	}
	return resp, nil
}

func (c *Ofile) CreateOrUpdateFile(filepath string, content []byte) error {
	u := fmt.Sprintf("%s/vault/%s", c.BaseURL, url.PathEscape(filepath))
	resp, err := c.doRequest(http.MethodPut, u, bytes.NewReader(content), []map[string]string{types.ContentMdHead, types.AcceptAllHead})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Accept any 2xx code as success
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("failed to create/update file: %s", resp.Status)
	}
	log.Debugf("created/updated file: %s content: %s", filepath, string(content))
	return nil
}

func (c *Ofile) ReadFile(filepath string) ([]byte, error) {
	u := fmt.Sprintf("%s/vault/%s", c.BaseURL, strings.Trim(filepath, "/"))
	resp, err := c.doRequest(http.MethodGet, u, nil, []map[string]string{types.AcceptMdHead})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Debugf("Failed to read file: %s with status: %s", filepath, resp.Status)
		return nil, fmt.Errorf("failed to read file: %s", resp.Status)
	}
	return io.ReadAll(resp.Body)
}

func (c *Ofile) DeleteFile(filepath string) error {
	u := fmt.Sprintf("%s/vault/%s", c.BaseURL, strings.Trim(filepath, "/"))
	resp, err := c.doRequest(http.MethodDelete, u, nil, []map[string]string{types.AcceptAllHead})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("failed to delete file: %s", resp.Status)
	}
	log.Debugf("deleted file: %s", filepath)
	return nil
}

func (c *Ofile) ListFiles(folder string) ([]string, error) {
	u := fmt.Sprintf("%s/vault/%s/", c.BaseURL, strings.Trim(folder, "/"))
	resp, err := c.doRequest(http.MethodGet, u, nil, []map[string]string{types.AcceptJsonHead})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Debugf("Failed to list files: %s with url %s", resp.Status, u)
		return nil, fmt.Errorf("failed to list files: %s", resp.Status)
	}
	filesResponse := types.FileListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&filesResponse); err != nil {
		return nil, err
	}
	log.Debugf("listed files: %s", folder)
	return filesResponse.Files, nil
}
