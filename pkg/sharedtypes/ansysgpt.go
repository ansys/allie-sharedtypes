package sharedtypes

// DefaultFields represents the default fields for the user query.
type AnsysGPTDefaultFields struct {
	QueryWord         string
	FieldName         string
	FieldDefaultValue string
}

// ACSSearchResponse represents the response from the ACS search.
type ACSSearchResponse struct {
	Physics             string  `json:"physics"`
	SourceTitleLvl3     string  `json:"sourceTitle_lvl3"`
	SourceURLLvl3       string  `json:"sourceURL_lvl3"`
	TokenSize           int     `json:"tokenSize"`
	SourceTitleLvl2     string  `json:"sourceTitle_lvl2"`
	Weight              float64 `json:"weight"`
	SourceURLLvl2       string  `json:"sourceURL_lvl2"`
	Product             string  `json:"product"`
	Content             string  `json:"content"`
	TypeOFasset         string  `json:"typeOFasset"`
	Version             string  `json:"version"`
	SearchScore         float64 `json:"@search.score"`
	SearchRerankerScore float64 `json:"@search.rerankerScore"`
}

// AnsysGPTCitation represents the citation from the AnsysGPT.
type AnsysGPTCitation struct {
	Title     string  `json:"Title"`
	URL       string  `json:"URL"`
	Relevance float64 `json:"Relevance"`
}
