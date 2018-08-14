package icalendar

import (
	"net/url"
	"time"

	"github.com/Hatch1fy/errors"
	"github.com/Hatch1fy/uuid"
)

const (
	header = "BEGIN:VCALENDAR\r\nVERSION:2.0\r\nPRODID:-//%s//%s//EN\r\nCALSCALE:GREGORIAN\r\nMETHOD:PUBLISH\r\nBEGIN:VEVENT\r\n"
	footer = "END:VEVENT\r\nEND:VCALENDAR"

	dateFmt = "20060102T150405Z"
)

// Event represents a Calendar event
type Event struct {
	// This property defines the persistent, globally unique identifier for the calendar component.
	UID string `json:"uid"`
	// This property defines the organizer for the calendar component.
	Organizer string `json:"organizer"`
	// This property defines a short summary or subject for the calendar component.
	Summary string `json:"summary"`
	// This property provides a more complete description of the calendar component, than that provided by the "SUMMARY" property.
	Description string `json:"description"`
	// This property defines a Uniform Resource Locator (URL) associated with the iCalendar object.
	URL string `json:"url"`
	// This property defines the revision sequence number of the calendar component within a sequence of revisions.
	Sequence int64
	// This property defines the overall status or confirmation for the calendar component.
	Status string
	// This property defines whether an event is transparent or not to busy time searches.
	Transparent string
	// This property specifies when the calendar component begins.
	Start time.Time
	// This property specifies when the calendar component ends.
	End time.Time
	// This property specified when the calendar component was created.
	Created time.Time
	// This property defines a rule or repeating pattern for recurring events.
	RepeatRule *RepeatRule
	// This property defines the categories for a calendar component.
	Categories []string
	// The property defines the intended venue for the activity defined by a calendar component.
	Location string
	// This property specifies information related to the global position for the activity specified by a calendar component.
	Geo *Coordinate
}

// validateURL will ensure the event has a valid URL (if set)
func (e *Event) validateURL() (err error) {
	if len(e.URL) == 0 {
		// URL is not set, return early
		return
	}

	// Attempt to parse the url as a url.URL
	_, err = url.Parse(e.URL)
	return
}

// Sanitize will ensure the required fields have default data (if empty)
func (e *Event) Sanitize() {
	if e.Created.IsZero() {
		// Created is not set, set it as the current time
		e.Created = time.Now()
	}

	if e.UID == "" {
		// UID is not set, create a new uuid.UUID
		u := uuid.New()
		// Set UID as the string representation of uuid
		e.UID = u.String()
	}
}

// Validate will validate an event
func (e *Event) Validate() (err error) {
	var errs errors.ErrorList
	// Ensure URL is valid (if set)
	errs.Push(e.validateURL())
	// Return the list of collected errors
	// TODO: Dive deeper into validations if necessary, some notables:
	//			- Repeat rules
	//			- Start/end time being inverse
	return errs.Err()
}

func (e *Event) String() (out string) {
	var buf []byte
	// Set header
	buf = append(buf, getHeader("Hatchify", "Hatch app")...)
	// Set UID (will default if unset)
	buf = appendString(buf, "UID:", e.UID, "\r\n")
	// Set organizer
	buf = appendString(buf, "ORGANIZER:MAILTO:", e.Organizer, "\r\n")
	// Set summary
	buf = appendString(buf, "SUMMARY:", e.Summary, "\r\n")
	// Set description
	buf = appendString(buf, "DESCRIPTION:", e.Description, "\r\n")
	// Set description (for Microsoft).. my heart gently weeps
	buf = appendString(buf, "X-ALT-DESC;FMTTYPE=text/html:", e.Description, "\r\n")
	// Set url
	buf = appendString(buf, "URL:", e.URL, "\r\n")
	// Set sequence
	buf = appendInt64(buf, "SEQUENCE:", e.Sequence, "\r\n")
	// Set status
	buf = appendString(buf, "STATUS:", e.Status, "\r\n")
	// Set transparent
	buf = appendString(buf, "TRANSPARENT:", e.Transparent, "\r\n")
	// Set start timestamp
	buf = appendTime(buf, "DTSTART:", e.Start, "\r\n")
	// Set end timestamp
	buf = appendTime(buf, "DTEND:", e.End, "\r\n")
	// Set creation timestamp (will default if unset)
	buf = appendTime(buf, "DTSTAMP:", e.Created, "\r\n")
	// Set repeating rule (see RepeatingRule type)
	buf = appendStringer(buf, "RRULE:", e.RepeatRule, "\r\n")
	// Set categories
	buf = appendStringSlice(buf, "CATEGORIES:", e.Categories, "\r\n")
	// Set location
	buf = appendString(buf, "LOCATION:", e.Location, "\r\n")
	// Set geo coordinates (see Coordinate type)
	buf = appendStringer(buf, "GEO:", e.Geo, "\r\n")
	// Set footer
	buf = append(buf, footer...)
	// Convert the byteslice buffer to a string and return
	return string(buf)
}
