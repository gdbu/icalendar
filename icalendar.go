package icalendar

import (
	"fmt"
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

func getHeader(company, product string) (out []byte) {
	return []byte(fmt.Sprintf(header, company, product))
}

/*
BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//ZContent.net//Zap Calendar 1.0//EN
CALSCALE:GREGORIAN
METHOD:PUBLISH
BEGIN:VEVENT
SUMMARY:Abraham Lincoln
UID:c7614cff-3549-4a00-9152-d25cc1fe077d
SEQUENCE:0
STATUS:CONFIRMED
TRANSP:TRANSPARENT
RRULE:FREQ=YEARLY;INTERVAL=1;BYMONTH=2;BYMONTHDAY=12
DTSTART:20080212
DTEND:20080213
DTSTAMP:20150421T141403
CATEGORIES:U.S. Presidents,Civil War People
LOCATION:Hodgenville\, Kentucky
GEO:37.5739497;-85.7399606
DESCRIPTION:Born February 12\, 1809\nSixteenth President (1861-1865)\n\n\n
 \nhttp://AmericanHistoryCalendar.com
URL:http://americanhistorycalendar.com/peoplecalendar/1,328-abraham-lincol
 n
END:VEVENT
END:VCALENDAR
*/

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

func (e *Event) validateURL() (err error) {
	if len(e.URL) == 0 {
		return
	}

	_, err = url.Parse(e.URL)
	return
}

// Sanitize will ensure the required fields have default data (if empty)
func (e *Event) Sanitize() {
	if e.Created.IsZero() {
		e.Created = time.Now()
	}

	if e.UID == "" {
		u := uuid.New()
		e.UID = u.String()
	}
}

// Validate will validate an event
func (e *Event) Validate() (err error) {
	var errs errors.ErrorList
	errs.Push(e.validateURL())
	return errs.Err()
}

func (e *Event) String() (out string) {
	var buf []byte
	buf = append(buf, getHeader("Hatchify", "Hatch app")...)
	buf = appendString(buf, "UID:", e.UID, "\r\n")
	buf = appendString(buf, "ORGANIZER:MAILTO:", e.Organizer, "\r\n")
	buf = appendString(buf, "SUMMARY:", e.Summary, "\r\n")
	buf = appendString(buf, "DESCRIPTION:", e.Description, "\r\n")
	buf = appendString(buf, "X-ALT-DESC;FMTTYPE=text/html:", e.Description, "\r\n")
	buf = appendString(buf, "URL:", e.URL, "\r\n")
	buf = appendInt64(buf, "SEQUENCE:", e.Sequence, "\r\n")
	buf = appendString(buf, "STATUS:", e.Status, "\r\n")
	buf = appendString(buf, "TRANSPARENT:", e.Transparent, "\r\n")
	buf = appendTime(buf, "DTSTART:", e.Start, "\r\n")
	buf = appendTime(buf, "DTEND:", e.End, "\r\n")
	buf = appendTime(buf, "DTSTAMP:", e.Created, "\r\n")
	buf = appendStringer(buf, "RRULE:", e.RepeatRule, "\r\n")
	buf = appendStringSlice(buf, "CATEGORIES:", e.Categories, "\r\n")
	buf = appendString(buf, "LOCATION:", e.Location, "\r\n")
	buf = appendStringer(buf, "GEO:", e.Geo, "\r\n")
	buf = append(buf, footer...)
	return string(buf)
}
