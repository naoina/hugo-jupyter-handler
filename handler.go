package jupyterhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gohugoio/hugo/hugolib"
	"github.com/gohugoio/hugo/source"
)

type JupyterNotebookHandler struct {
	markdownHandler hugolib.Handler
}

func NewJupyterNotebookHandler() *JupyterNotebookHandler {
	return &JupyterNotebookHandler{
		markdownHandler: hugolib.FindHandler("markdown"),
	}
}

func (h JupyterNotebookHandler) Read(f *source.File, s *hugolib.Site) hugolib.HandledResult {
	ipynb := JupyterNotebook{}
	if err := json.Unmarshal(f.Bytes(), &ipynb); err != nil {
		panic(err)
	}
	buf, err := json.Marshal(ipynb.Metadata.FrontMatter)
	if err != nil {
		panic(err)
	}
	buf = append(buf, '\n', '\n')
	for _, cell := range ipynb.Cells {
		source := strings.Join(cell.Source, "")
		switch cell.CellType {
		case "markdown":
			for name, a := range cell.Attachments {
				target := fmt.Sprintf("attachment:%s", name)
				for mime, data := range a {
					dataURL := fmt.Sprintf("data:%s;base64,%s", mime, data)
					source = strings.Replace(source, target, dataURL, -1)
				}
			}
			buf = append(buf, source...)
		case "code":
			buf = append(buf, fmt.Sprintf("```%s\n%s\n```", ipynb.Metadata.KernelSpec.Language, source)...)
			for _, o := range cell.Outputs {
				buf = append(buf, '\n')
				switch o.OutputType {
				case "stream":
					for _, s := range o.Text {
						buf = append(buf, fmt.Sprintf("```text\n%s\n```", strings.TrimSpace(s))...)
					}
				}
			}
		}
		buf = append(buf, '\n')
	}
	f.Contents = bytes.NewReader(buf)
	r := h.markdownHandler.Read(f, s)
	r.Page().Markup = "markdown"
	return r
}

func (h JupyterNotebookHandler) FileConvert(*source.File, *hugolib.Site) hugolib.HandledResult {
	return hugolib.HandledResult{}
}

func (h JupyterNotebookHandler) Extensions() []string {
	return []string{"ipynb"}
}

func (h JupyterNotebookHandler) PageConvert(p *hugolib.Page) hugolib.HandledResult {
	return h.markdownHandler.PageConvert(p)
}

func init() {
	hugolib.RegisterHandler(NewJupyterNotebookHandler())
}
