package main

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	var sums []int
	for _, num := range numsToSum {
		sums = append(sums, Sum(num))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sums = append(sums, Sum(nums))
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
