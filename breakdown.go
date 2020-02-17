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

// int value of -1 => '*'
type CronField struct {
	fieldVals        []int
	postSepFieldVals []int
	sep              rune
}

type CronBreakdown struct {
	minute         CronField
	hour           CronField
	dayMonth       CronField
	month          CronField
	dayWeek        CronField
	validationErrs []error
}

func BuildCronField() (cf CronField) {
	cf = CronField{}

	return
}

func BuildBreakdown() (cb *CronBreakdown) {
	cb = &CronBreakdown{}

	return
}
