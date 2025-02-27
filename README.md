# iCalendar
iCalendar is a library to assist in producing iCalendar events (.ics).

## Features
Using an ```icalendar.Event``` one can create iCalendar events, which can be sent to any email service provider who supports ics format calendar events. The majority of providers should support ics format events. We currently test for:

- GMail
- Outlook
- Apple Mail *(intended to support, no testing has been performed yet)*

## Usage

```go
package main

import (
	"log"
	"os"
	"time"

	"github.com/Hatch1fy/icalendar"
)

func main() {
	var (
		e   icalendar.Event
		f   *os.File
		err error
	)

	e.Organizer = "johndoe@gmail.com"
	e.Summary = "Birthday party!"
	e.Description = "It's party time!\n\nGet your swim trunks on and let's enjoy the summer sun!"
	e.Start = time.Date(2018, 6, 12, 0, 0, 0, 0, time.UTC)
	e.End = e.Start.Add(time.Hour * 3)

	if f, err = os.Create("birthday.ics"); err != nil {
		log.Fatalf("Error creating iCalendar file: %v", err)
	}
	defer f.Close()

	f.Write([]byte(e.String()))
}

```
