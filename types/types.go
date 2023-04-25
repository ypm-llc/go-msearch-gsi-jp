package types

// Geometry value of GeoJSON
type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

// Properties value of GeoJSON
type Properties struct {
	AddressCode string `json:"addressCode"`
	Title       string `json:"title"`
	DataSource  string `json:"dataSource"`
}

// Feature value of GeoJSON
type Feature struct {
	Geometry   *Geometry   `json:"geometry"`
	Type       string      `json:"type"`
	Properties *Properties `json:"properties"`
}

// Results of Reverse Geocode
type AddressResults struct {
	MuniCd string `json:"muniCd"`
	Lv01Nm string `json:"lv01Nm"`
}

// Address value of Reverse Geocode
type Address struct {
	Results *AddressResults `json:"results"`
}
