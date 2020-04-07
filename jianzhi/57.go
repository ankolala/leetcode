package main

import "fmt"



func findContinuousSequence(target int) [][]int {

	i, j, sum  := 1, 2, 3

	final := make([][]int,0)

	for j < target {


		if sum < target {
			j += 1
			sum += j
		}

		if sum == target {
			result := make([]int, 0)

			for t:=i;t<=j;t++{
				result = append(result, t)
			}

			final  = append(final, result)
			j += 1
			sum += j
		}

		if sum > target {
			sum -= i
			i += 1
		}

	}
	return final


}

func main() {

	fmt.Println(findContinuousSequence(2))
}
