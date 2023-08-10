package calculate

import (
	"regexp"
	"strconv"
	"sweetweather.test/pkg/calculate"
)

type Calculate struct{}

func (Calculate) Calculate(data string) (int, error) {
	var (
		result    int
		values    []int
		signArray []string
	)

	re := regexp.MustCompile(`([-+]?\d+)`)
	matches := re.FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		number, _ := strconv.Atoi(match[1])
		values = append(values, number)
	}

	re = regexp.MustCompile(`[-+]`)
	signArray = re.FindAllString(data, -1)

	resultArray := make([]int, len(signArray)+1)
	resultArray[0] = values[0]

	for i, sign := range signArray {
		number := values[i+1]
		if sign == "-" {
			number *= -1
		}
		resultArray[i+1] = number
	}

	result = calculate.SumArray(values...)

	return result, nil
}

func NewCalculate() *Calculate {
	return &Calculate{}
}
