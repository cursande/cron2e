package cron2e

import (
	"fmt"
	"os"
)

func Run() string {
	if len(os.Args) < 2 {
		fmt.Println("Enter a cron expression")
		return ""
	}
	expr := os.Args[1]

	parser, err := ParserForExpression(expr)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	breakdown, err := parser.parse()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if !Validate(breakdown) {
		for i := 0; i < len(breakdown.validationErrs); i++ {
			fmt.Println(breakdown.validationErrs[i])
		}
		return ""
	}

	// Let's have a concept of DynamicTranslate, and StaticTranslate (e.g. for predefined)
	translation := Translate(breakdown)

	fmt.Println(translation)

	return translation
}
