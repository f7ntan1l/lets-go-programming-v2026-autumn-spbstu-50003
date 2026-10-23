package main

import (
	"fmt"
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
	noOptimalTemp  = -1

	ruleAtLeast = ">="
	ruleAtMost  = "<="
)

type AirConditioner struct {
	lowerBound int
	upperBound int
}

func NewAirConditioner() AirConditioner {
	return AirConditioner{
		lowerBound: defaultMinTemp,
		upperBound: defaultMaxTemp,
	}
}

func (ac *AirConditioner) Adjust(rule string, temp int) {
	switch rule {
	case ruleAtLeast:
		if temp > ac.lowerBound {
			ac.lowerBound = temp
		}
	case ruleAtMost:
		if temp < ac.upperBound {
			ac.upperBound = temp
		}
	}
}

func (ac *AirConditioner) OptimalTemperature() int {
	if ac.lowerBound <= ac.upperBound {
		return ac.lowerBound
	}

	return noOptimalTemp
}

func main() {
	var totalDepartments int
	if _, err := fmt.Scan(&totalDepartments); err != nil {
		fmt.Println(err)

		return
	}

	for range totalDepartments {
		var totalStaff int
		if _, err := fmt.Scan(&totalStaff); err != nil {
			fmt.Println(err)

			return
		}

		conditioner := NewAirConditioner()

		for range totalStaff {
			var (
				rule string
				temp int
			)

			if _, err := fmt.Scan(&rule, &temp); err != nil {
				fmt.Println(err)

				return
			}

			conditioner.Adjust(rule, temp)
			fmt.Println(conditioner.OptimalTemperature())
		}
	}
}
