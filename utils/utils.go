package utils

import "errors"

func Zip(lists ...[]string) ([][]string, error) {
	if len(lists) == 0 {
		return lists, nil
	}

	size := len(lists[0])

	for _, l := range lists[1:] {
		if size != len(l) {
			return nil, errors.New("Zip function only supports lists with the same length")
		}
		size = len(l)
	}

	res := make([][]string, 0)

	for c := 0; c < size; c++ {
		temp := make([]string, 0)
		for r := 0; r < len(lists); r++ {
			temp = append(temp, lists[r][c])
		}
		res = append(res, temp)
	}

	return res, nil
}

func GetStats(nums []int) (int, int, int, float64) {
	if len(nums) == 0 {
		return -1, -1, -1, -1
	}

	min := nums[0]
	max := nums[0]
	total := nums[0]

	for _, v := range nums[1:] {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}

		total += v
	}

	return min, max, total, float64(total) / float64(len(nums))
}

func Reverse(lists ...[]string) {
	for i := 0; i < len(lists)/2; i++ {
		lists[i], lists[len(lists)-1-i] = lists[len(lists)-1-i], lists[i]
	}
}
