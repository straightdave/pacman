package main

var (
	levelMap1 = convertToMatrix(
		10, []int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 3, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 3, 3, 3, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 2, 0, 0, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 3, 0, 0, 0, 0, 1,
			1, 0, 0, 3, 0, 0, 0, 3, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		})
)

func convertToMatrix(n int, raw []int) [][]int {
	var res [][]int
	for i := range n {
		res = append(res, raw[i*n:(i+1)*n])
	}
	return res
}
