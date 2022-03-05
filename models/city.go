package models

// Cities represents the return from our nearby cities route
type Cities struct {
	ReferenceCity    string         `json:"referenceCity"`
	NearbyPopulation int64          `json:"nearbyPopulation"`
	NearbyCities     []GeoAPICities `json:"nearbyCities"`
}
