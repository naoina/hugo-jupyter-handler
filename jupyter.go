package jupyterhandler

type JupyterNotebook struct {
	Cells []struct {
		CellType string                 `json:"cell_type"`
		Metadata map[string]interface{} `json:"metadata"`
		Source   []string               `json:"source"`
		Outputs  []struct {
			OutputType string   `json:"output_type"`
			Text       []string `json:"text"`
		} `json:"outputs"`
		Attachments map[string]map[string]string `json:"attachments"`
	} `json:"cells"`
	Metadata struct {
		FrontMatter map[string]interface{} `json:"frontmatter"`

		KernelSpec struct {
			DisplayName string `json:"display_name"`
			Language    string `json:"language"`
			Name        string `json:"name"`
		} `json:"kernelspec"`

		LanguageInfo struct {
			FileExtension    string `json:"file_extension"`
			MimeType         string `json:"mimetype"`
			Name             string `json:"name"`
			NBConvertEporter string `json:"nbconvert_exporter"`
			PygmentsLexer    string `json:"pygments_lexer"`
			Version          string `json:"version"`
		} `json:"language_info"`
	} `json:"metadata"`
	NBFormat      int `json:"nbformat"`
	NBFormatMinor int `json:"nbformat_minor"`
}
