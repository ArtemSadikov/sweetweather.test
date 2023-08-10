package calculate

import (
	"regexp"
	"strconv"
	"sweetweather.test/pkg/calculate"
)

type Calculate struct{}

func (Calculate) Calculate(data string) (int, error) {
	var (
		result int
		values []int
	)

	re := regexp.MustCompile(`([-+]?\d+)`)
	matches := re.FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		number, _ := strconv.Atoi(match[1])
		sign := match[0]
		if sign == "-" {
			number *= -1
		}

		values = append(values, number)
	}

	result = calculate.SumArray(values...)

	return result, nil
}

func NewCalculate() *Calculate {
	return &Calculate{}
}
