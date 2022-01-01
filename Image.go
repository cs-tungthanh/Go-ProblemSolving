package main

import "fmt"

/*
	Problem: https://leetcode.com/problems/flood-fill
	Param:
		sr: starting row
		sc: starting col
		newColor: value of color
	Solution: using Depth first search
*/
// img := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
// result := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
// newImg := floodFill(img, 1, 1, 2)
// fmt.Println("img", img)
// fmt.Println("newImg", newImg)
// fmt.Println("result", result)
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	h, w, currentColor := len(image), len(image[0]), image[sr][sc]
	if currentColor == newColor {
		return image
	}
	image[sr][sc] = newColor
	if (sr-1) >= 0 && image[sr-1][sc] == currentColor {
		image = floodFill(image, sr-1, sc, newColor)
	}
	if (sr+1) <= h-1 && image[sr+1][sc] == currentColor {
		image = floodFill(image, sr+1, sc, newColor)
	}
	if (sc-1) >= 0 && image[sr][sc-1] == currentColor {
		image = floodFill(image, sr, sc-1, newColor)
	}
	if (sc+1) <= w-1 && image[sr][sc+1] == currentColor {
		image = floodFill(image, sr, sc+1, newColor)
	}
	return image
}

/*
	Problem: https://leetcode.com/problems/max-area-of-island/
	Solution: using DFS
*/
func maxAreaOfIsland(grid [][]int) int {
	maxValue := 0
	for ri, rv := range grid {
		for ci, cv := range rv {
			if cv == 1 {
				l := 0
				dfs(grid, ri, ci, &l)
				if l > maxValue {
					maxValue = l
				}
			}
		}
	}
	return maxValue
}

func dfs(grid [][]int, sr, sc int, l *int) {
	if grid[sr][sc] == 0 || grid[sr][sc] == -1 {
		return
	}
	*l++
	grid[sr][sc] = -1
	if sr >= 1 {
		dfs(grid, sr-1, sc, l)
	}
	if sc >= 1 {
		dfs(grid, sr, sc-1, l)
	}
	if sr+1 <= len(grid)-1 {
		dfs(grid, sr+1, sc, l)
	}
	if sc+1 <= len(grid[0])-1 {
		dfs(grid, sr, sc+1, l)
	}
}

func main() {
	grid := [][]int{
		{1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	predict := maxAreaOfIsland(grid)
	fmt.Println("Predict", predict)
	fmt.Println("GT", 8, grid)
}
