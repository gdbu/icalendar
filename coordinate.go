package icalendar

import "strconv"

// Coordinate represents a geo coordinate
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (c *Coordinate) String() (out string) {
	return strconv.FormatFloat(c.Lat, 'f', 8, 64) + ";" + strconv.FormatFloat(c.Lon, 'f', 8, 64)
}
