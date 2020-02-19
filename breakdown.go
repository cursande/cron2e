package cron2e

const Wildcard int = -1

// When referring to specific field types.
const (
	Minute   uint8 = 1
	Hour     uint8 = 2
	DayMonth uint8 = 3
	Month    uint8 = 4
	DayWeek  uint8 = 5
)

type CronValue struct {
	fieldVal        int
	postSepFieldVal int
	sep             rune
}

type CronBreakdown struct {
	minutes        []CronValue
	hours          []CronValue
	dayMonths      []CronValue
	months         []CronValue
	dayWeeks       []CronValue
	validationErrs []error
}

func BuildCronValue() (cv CronValue) {
	cv = CronValue{}

	return
}

func BuildBreakdown() (cb *CronBreakdown) {
	cb = &CronBreakdown{}

	return
}
