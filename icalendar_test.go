package icalendar

import (
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestICalendar(t *testing.T) {
	var (
		e   Event
		f   *os.File
		err error
	)

	e.Organizer = "josh@montoya.io"
	e.Summary = "Morning stand-up"
	e.Description = "Rise and shine! It's our favorite time of the day!\n\nLet's sync up and figure out how everyone's day is looking\n\n<a href=\"https://zoom.us/j/4777894195\" target=\"_blank\">https://zoom.us/j/4777894195</a>"
	if e.URL, err = url.Parse("https://zoom.us/j/4777894195"); err != nil {
		t.Fatalf("Error parsing URL: %v", err)
	}

	e.Start = time.Now().UTC()
	e.End = time.Now().UTC().Add(time.Hour)
	//e.RepeatRule =
	fmt.Println("Start", e.Start.Format(dateFmt))

	if f, err = os.Create("test.ics"); err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer f.Close()

	f.Write([]byte(e.String()))
}
