package schemas

type Index struct {
	Name        string         `json:"name"`
	StorageType string         `json:"storage_type"`
	Mappings    *IndexMappings `json:"mappings"`
}

type IndexMappings struct {
	Properties *IndexProperty `json:"properties"`
}

type IndexProperty map[string]*IndexPropertyT

type IndexPropertyT struct {
	Type           string `json:"type"`
	Index          bool   `json:"index"`
	Store          bool   `json:"store"`
	Sortable       bool   `json:"sortable"`
	Aggregatable   bool   `json:"aggregatable"`
	Highlightable  bool   `json:"highlightable"`
	Analyzer       string `json:"analyzer"`
	SearchAnalyzer string `json:"search_analyzer"`
	Format         string `json:"format"`
}
