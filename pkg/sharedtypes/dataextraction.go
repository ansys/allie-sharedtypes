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

type CodeGenerationElement struct {
	Guid string             `json:"guid"`
	Type CodeGenerationType `json:"type"`

	NamePseudocode string `json:"name_pseudocode"` // Function name without dependencies
	NameFormatted  string `json:"name_formatted"`  // Name of the function with spaces and without parameters
	Description    string `json:"description"`

	Name              string   `json:"name"`
	Dependencies      []string `json:"dependencies"`
	Summary           string   `json:"summary"`
	ReturnType        string   `json:"return"`
	ReturnElementList []string `json:"return_element_list"`
	ReturnDescription string   `json:"return_description"` // Return description
	Remarks           string   `json:"remarks"`

	// Only for type "function" or "method"
	Parameters []XMLMemberParam `json:"parameters"`
	Example    XMLMemberExample `json:"example"`

	// Only for type "enum"
	EnumValues []string `json:"enum_values"`
}

// Enum values for CodeGenerationType
type CodeGenerationType string

const (
	Function  CodeGenerationType = "Function"
	Method    CodeGenerationType = "Method"
	Class     CodeGenerationType = "Class"
	Parameter CodeGenerationType = "Parameter"
	Enum      CodeGenerationType = "Enum"
	Module    CodeGenerationType = "Module"
)

type XMLMemberExample struct {
	Description string               `xml:",chardata" json:"description"` // Text content of <example>
	Code        XMLMemberExampleCode `xml:"code,omitempty" json:"code"`   // Optional <code> element
}

type XMLMemberExampleCode struct {
	Type string `xml:"type,attr" json:"type"` // Attribute for <code>
	Text string `xml:",chardata" json:"text"` // Text content of <code>
}

type XMLMemberParam struct {
	Name        string `xml:"name" json:"name"`             // Attribute for <param>
	Type        string `xml:"type,omitempty" json:"type"`   // Attribute for <param>
	Description string `xml:",chardata" json:"description"` // Text content of <param>
}

type CodeGenerationExample struct {
	Guid                   string            `json:"guid"`
	Name                   string            `json:"name"`
	Dependencies           []string          `json:"dependencies"`
	DependencyEquivalences map[string]string `json:"dependency_equivalences"`
	Chunks                 []string          `json:"chunks"`
}

type CodeGenerationUserGuideSection struct {
	Name            string   `json:"name"`
	Title           string   `json:"title"`
	IsFirstChild    bool     `json:"is_first_child"`
	NextSibling     string   `json:"next_sibling"`
	NextParent      string   `json:"next_parent"`
	DocumentName    string   `json:"document_name"`
	Parent          string   `json:"parent"`
	Content         string   `json:"content"`
	Level           int      `json:"level"`
	Link            string   `json:"link"`
	ReferencedLinks []string `json:"referenced_links"`
	Chunks          []string `json:"chunks"`
}
