package awsrate

// https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html
type AWSRateFormat struct{}

const (
	Minute uint8 = 1
	Hour   uint8 = 2
	Day    uint8 = 3
)

type CronBreakdown struct {
	timeValue uint8
	interval  int
}

var format = AWSRateFormat{}
