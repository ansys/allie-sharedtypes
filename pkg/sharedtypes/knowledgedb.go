package sharedtypes

// DbFilters represents the filters for the database.
type DbFilters struct {
	// Filters for string fields
	GuidFilter         []string `json:"guid,omitempty"`
	DocumentIdFilter   []string `json:"document_id,omitempty"`
	DocumentNameFilter []string `json:"document_name,omitempty"`
	LevelFilter        []string `json:"level,omitempty"`

	// Filters for array fields
	TagsFilter     DbArrayFilter `json:"tags,omitempty"`
	KeywordsFilter DbArrayFilter `json:"keywords,omitempty"`

	// Filters for JSON fields
	MetadataFilter []DbJsonFilter `json:"metadata,omitempty"`
}

// DbArrayFilter represents the filter for an array field in the database.
type DbArrayFilter struct {
	NeedAll    bool     `json:"needAll"`
	FilterData []string `json:"filterData"`
}

// DbJsonFilter represents the filter for a JSON field in the database.
type DbJsonFilter struct {
	FieldName  string   `json:"fieldName"`
	FieldType  string   `json:"fieldType" description:"Can be either string or array."` // "string" or "array"
	FilterData []string `json:"filterData"`
	NeedAll    bool     `json:"needAll" description:"Only needed if the FieldType is array."` // only needed for array fields
}

// DbData represents the data stored in the database.
type DbData struct {
	Guid              string                 `json:"guid"`
	DocumentId        string                 `json:"document_id"`
	DocumentName      string                 `json:"document_name"`
	Text              string                 `json:"text"`
	Keywords          []string               `json:"keywords"`
	Summary           string                 `json:"summary"`
	Embedding         []float32              `json:"embeddings"`
	Tags              []string               `json:"tags"`
	Metadata          map[string]interface{} `json:"metadata"`
	ParentId          string                 `json:"parent_id"`
	ChildIds          []string               `json:"child_ids"`
	PreviousSiblingId string                 `json:"previous_sibling_id"`
	NextSiblingId     string                 `json:"next_sibling_id"`
	LastChildId       string                 `json:"last_child_id"`
	FirstChildId      string                 `json:"first_child_id"`
	Level             string                 `json:"level"`
	HasNeo4jEntry     bool                   `json:"has_neo4j_entry"`
}

// DbResponse represents the response from the database.
type DbResponse struct {
	Guid              string                 `json:"guid"`
	DocumentId        string                 `json:"document_id"`
	DocumentName      string                 `json:"document_name"`
	Text              string                 `json:"text"`
	Keywords          []string               `json:"keywords"`
	Summary           string                 `json:"summary"`
	Embedding         []float32              `json:"embeddings"`
	Tags              []string               `json:"tags"`
	Metadata          map[string]interface{} `json:"metadata"`
	ParentId          string                 `json:"parent_id"`
	ChildIds          []string               `json:"child_ids"`
	PreviousSiblingId string                 `json:"previous_sibling_id"`
	NextSiblingId     string                 `json:"next_sibling_id"`
	LastChildId       string                 `json:"last_child_id"`
	FirstChildId      string                 `json:"first_child_id"`
	Distance          float64                `json:"distance"`
	Level             string                 `json:"level"`
	HasNeo4jEntry     bool                   `json:"has_neo4j_entry"`

	// Siblings
	Parent    *DbData  `json:"parent,omitempty"`
	Children  []DbData `json:"children,omitempty"`
	LeafNodes []DbData `json:"leaf_nodes,omitempty"`
	Siblings  []DbData `json:"siblings,omitempty"`
}

// DBListCollectionsOutput represents the output of listing collections in the database.
type DBListCollectionsOutput struct {
	Success     bool     `json:"success" description:"Returns true if the collections were listed successfully. Returns false or an error if not."`
	Collections []string `json:"collections" description:"A list of collection names."`
}

// GeneralNeo4jQueryInput represents the input for executing a Neo4j query.
type GeneralNeo4jQueryInput struct {
	Query string `json:"query" description:"Neo4j query to be executed. Required for executing a query." required:"true"`
}

// GeneralNeo4jQueryOutput represents the output of executing a Neo4j query.
type GeneralNeo4jQueryOutput struct {
	Success  bool          `json:"success" description:"Returns true if the query was executed successfully. Returns false or an error if not."`
	Response neo4jResponse `json:"response" description:"Summary and records of the query execution."`
}

// neo4jResponse represents the response from the Neo4j query.
type neo4jResponse struct {
	Record          neo4jRecord     `json:"record"`
	SummaryCounters summaryCounters `json:"summaryCounters"`
}

// neo4jRecord represents the record from the Neo4j query.
type neo4jRecord []struct {
	Values []value `json:"Values"`
}

// value represents the value from the Neo4j query.
type value struct {
	Id        int      `json:"Id"`
	NodeTypes []string `json:"Labels"`
	Props     props    `json:"Props"`
}

// props represents the properties from the Neo4j query.
type props struct {
	CollectionName string   `json:"collectionName"`
	DocumentId     string   `json:"documentId"`
	DocumentTypes  []string `json:"documentTypes,omitempty"`
	Guid           string   `json:"guid,omitempty"`
}

// summaryCounters represents the summary counters from the Neo4j query.
type summaryCounters struct {
	NodesCreated         int `json:"nodes_created"`
	NodesDeleted         int `json:"nodes_deleted"`
	RelationshipsCreated int `json:"relationships_created"`
	RelationshipsDeleted int `json:"relationships_deleted"`
	PropertiesSet        int `json:"properties_set"`
	LabelsAdded          int `json:"labels_added"`
	LabelsRemoved        int `json:"labels_removed"`
	IndexesAdded         int `json:"indexes_added"`
	IndexesRemoved       int `json:"indexes_removed"`
	ConstraintsAdded     int `json:"constraints_added"`
	ConstraintsRemoved   int `json:"constraints_removed"`
}
