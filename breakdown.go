package cron2e

// int value of -1 => '*'
type CronField struct {
	fieldVals        []int
	postSepFieldVals []int
	sep              rune
}

type CronBreakdown struct {
	second           CronField
	minute           CronField
	hour             CronField
	dayMonth         CronField
	dayWeek          CronField
	month            CronField
	validationErrors []error
}

func BuildCronField() (cf CronField) {
	cf = CronField{}

	return
}

func BuildBreakdown() (cb *CronBreakdown) {
	cb = &CronBreakdown{}

	return
}
