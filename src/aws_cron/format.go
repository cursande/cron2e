package awscron

// https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html
type AWSCronFormat struct{}

const (
	Wildcard int = -1
	Unset    int = -2
)

// When referring to specific field types.
const (
	Minute   uint8 = 1
	Hour     uint8 = 2
	DayMonth uint8 = 3
	Month    uint8 = 4
	DayWeek  uint8 = 5
	Year     uint8 = 6
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
	years          []CronValue
}

var format = AWSCronFormat{}
