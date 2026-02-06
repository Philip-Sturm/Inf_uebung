package slices

// RemoveNegatives entfernt alle negativen Zahlen aus dem Slice,
// indem es das Slice direkt über den übergebenen Pointer verändert.
func RemoveNegatives(nums []int) []int {
	var sl []int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			sl = append(sl, nums[i])
		}
	}
	return sl
}
