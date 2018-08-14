package icalendar

import "strconv"

// Coordinate represents a geo coordinate
type Coordinate struct {
	// Latitude
	Lat float64 `json:"lat"`
	// Longitude
	Lon float64 `json:"lon"`
}

func (c *Coordinate) String() (out string) {
	// Convert latitude to string
	lat := strconv.FormatFloat(c.Lat, 'f', 8, 64)
	// Convert longitude to string
	lon := strconv.FormatFloat(c.Lon, 'f', 8, 64)
	// Concatenate lat/lon with semi-colon as the delimiter
	return lat + ";" + lon
}
