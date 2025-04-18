package types

var (
	ContentMdHead      = map[string]string{"Content-Type": "text/markdown"}
	ContentJsonHead    = map[string]string{"Content-Type": "application/json"}
	ContentTypeAllHead = map[string]string{"Content-Type": "*/*"}

	AcceptJsonHead = map[string]string{"Accept": "application/json"}
	AcceptMdHead   = map[string]string{"Accept": "text/markdown"}
	AcceptAllHead  = map[string]string{"Accept": "*/*"}
)

type ObsidianFileRequest struct {
	Path    string `json:"path" description:"文件路径" required:"true"`
	Content string `json:"content,omitempty" description:"文件内容（创建/更新用）"`
}

type ObsidianFileListRequest struct {
	Path string `json:"path" description:"文件夹路径" required:"true"`
}

type FileListResponse struct {
	Files []string `json:"files"`
}

type ObsidianMdOptimizeRequest struct {
	Path    string `json:"path" description:"文件路径" required:"true"`
	Content string `json:"content" description:"Markdown content to optimize"`
}
