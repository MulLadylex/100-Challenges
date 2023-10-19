package main

func main() {

}

func singleNumber(nums []int) (res []int) {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	for num, occ := range freq {
		if occ == 1 {
			res = append(res, num)
		}
	}
	return
}
