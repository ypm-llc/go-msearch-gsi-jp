package types

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

type Properties struct {
	AddressCode string `json:"addressCode"`
	Title       string `json:"title"`
	DataSource  string `json:"dataSource"`
}

type Feature struct {
	Geometry   *Geometry   `json:"geometry"`
	Type       string      `json:"type"`
	Properties *Properties `json:"properties"`
}

type AddressResults struct {
	MuniCd string `json:"muniCd"`
	Lv01Nm string `json:"lv01Nm"`
}

type Address struct {
	Results *AddressResults `json:"results"`
}
