package cron2e

import (
	"fmt"
	"os"
)

func Run() (translation string) {
	if len(os.Args) < 2 {
		fmt.Println("Enter a cron expression")
		return
	}
	expr := os.Args[1]

	format, err := FormatForExpression(expr)

	if err != nil {
		fmt.Println(err)
		return
	}

	translation, errs := format.Translate(expr)

	if len(errs) > 0 {
		for i := 0; i < len(errs); i++ {
			fmt.Println(errs[i])
		}
		return
	}

	fmt.Println(translation)

	return
}
