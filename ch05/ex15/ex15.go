package ex15

import "fmt"

func max(vals ...int) (max int, err error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("No argument!")
	}
	max = vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func min(vals ...int) (min int, err error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("No argument!")
	}
	min = vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min, nil
}
