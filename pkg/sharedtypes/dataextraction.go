package sharedtypes

// DataExtractionDocumentData represents the data extracted from a document.
type DataExtractionDocumentData struct {
	DocumentName      string    `json:"documentName"`
	DocumentId        string    `json:"documentId"`
	Guid              string    `json:"guid"`
	Level             string    `json:"level"`
	ChildIds          []string  `json:"childIds"`
	ParentId          string    `json:"parentId"`
	PreviousSiblingId string    `json:"previousSiblingId"`
	NextSiblingId     string    `json:"nextSiblingId"`
	LastChildId       string    `json:"lastChildId"`
	FirstChildId      string    `json:"firstChildId"`
	Text              string    `json:"text"`
	Keywords          []string  `json:"keywords"`
	Summary           string    `json:"summary"`
	Embedding         []float32 `json:"embedding"`
}
