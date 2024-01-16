package cuda

// Range
type Range struct {
	Number      string `json:"number"`
	Description string `json:"description"`
}

type Ranges struct {
	Ranges []Range `json:"ranges"`
}

type RangesSimple struct {
	Ranges []string `json:"ranges"` // range numbers only
}

// Cluster
type Cluster struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type Clusters struct {
	Clusters []Cluster `json:"clusters"`
}

type ClustersSimple struct {
	Clusters []string `json:"clusters"` // cluster names only
}

// Box
type Box struct {
	Name                string `json:"name"`
	Enabled             bool   `json:"enabled"`
	IP                  string `json:"ip"`
	Description         string `json:"description"`
	Product             string `json:"product"`
	Model               string `json:"boxModel"`
	Revision            string `json:"revision"`
	ConfTemplate        string `json:"confTemplate,omitempty"`
	ConfTemplateBinding string `json:"confTemplateBinding"`
}

type Boxes struct {
	Boxes []Box `json:"boxes"`
}

type BoxesSimple struct {
	Boxes []string `json:"boxes"`
}
