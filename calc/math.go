package calc

import "errors"

func Add(ary ...int) (error, int) {
	sum := 0

	if len(ary) < 2 {
		return errors.New("Provide more than 2 number"), sum
	} else {
		for _, value := range ary {
			sum = sum + value
		}
		return nil, sum
	}
}
