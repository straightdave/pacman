package main

var (
	raw = []int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 3, 3, 3, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 3, 0, 0, 0, 0, 1,
		1, 0, 0, 3, 0, 0, 0, 3, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}

	levelMap1 = convertToMatrix(10, raw)
)

func convertToMatrix(n int, raw []int) [][]int {
	// loading to 2D matrix + pivoting

	var res [][]int
	for i := range n {
		var t []int
		for j := range n {
			n := j*10 + i
			t = append(t, raw[n])
		}
		res = append(res, t)
	}
	return res
}
