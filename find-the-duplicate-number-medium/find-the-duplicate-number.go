package main

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	renewedSlow := 0
	for renewedSlow != slow {
		renewedSlow = nums[renewedSlow]
		slow = nums[slow]
	}
	return slow
}

func main() {

}
