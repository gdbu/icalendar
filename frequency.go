package icalendar

const (
	// FrequencySecondly represents a frequency of every second
	FrequencySecondly = "SECONDLY"
	// FrequencyMinutely represents a frequency of every second
	FrequencyMinutely = "MINUTELY"
	// FrequencyHourly represents a frequency of every minute
	FrequencyHourly = "HOURLY"
	// FrequencyDaily represents a frequency of every hour
	FrequencyDaily = "DAILY"
	// FrequencyWeekly represents a frequency of every day
	FrequencyWeekly = "WEEKLY"
	// FrequencyMonthly represents a frequency of every month
	FrequencyMonthly = "MONTHLY"
	// FrequencyYearly represents a frequency of every year
	FrequencyYearly = "YEARLY"
)

// Frequency represents an event frequency
type Frequency string
