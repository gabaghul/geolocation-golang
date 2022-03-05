package models

// GeoAPICities represents Geo City API's return
type GeoAPICities struct {
	Data     []GeoAPICity `json:"data"`
	Metadata Metadata     `json:"metadata"`
}

// GepAPICity represents city data returned from API
type GeoAPICity struct {
	Id          int64   `json:"id"`
	WikiDataID  string  `json:"wikiDataId"`
	Type        string  `json:"type"`
	City        string  `json:"city"`
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionCode  string  `json:"regionCode"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Population  int64   `json:"population"`
}

// Metadata represents info for pagination control (not important for us, but i'll create a struct just in case)
type Metadata struct {
	CurrentOffset int64 `json:"currentOffset"`
	TotalCount    int64 `json:"totalCount"`
}
