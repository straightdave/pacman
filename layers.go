package main

var (
	levelMap1 = convertToMatrix(
		5, []int{
			0, 0, 0, 0, 0,
			0, 1, 0, 0, 0,
			0, 1, 2, 0, 0,
			0, 0, 0, 0, 1,
			0, 0, 1, 1, 0,
		})
)

func convertToMatrix(n int, raw []int) [][]int {
	var res [][]int
	for i := range n {
		res = append(res, raw[i*n:(i+1)*n])
	}
	return res
}
