package main

import "fmt"

// 转置数组

func transpose(matrix [][]int) [][]int {

	// 获得行, 列数目
	m, n := len(matrix), len(matrix[0])

	// 创建新的二维矩阵 新矩阵的行数是原矩阵的列数
	grid := make([][]int, n)

	for i := range grid {
		// 每个新行开辟一个 之前行数为长度的新数据
		grid[i] = make([]int, m)
	}

	fmt.Println(grid)

	for i, row := range matrix {
		for j, v := range row {
			grid[j][i] = v
		}
	}

	return grid
}

func main() {

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	fmt.Println(matrix)

	matrix = transpose(matrix)

	fmt.Println(matrix)

}
