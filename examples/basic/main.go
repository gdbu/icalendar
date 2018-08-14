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
