package main

import (
	"fmt"
	"testing"
)

func TestCreateMatrix(t *testing.T) {
	const (
		rowCount    = 5
		columnCount = 5
	)

	colors := []int{BLACK, WHITE, RED}

	matrix := createMatrix(rowCount, columnCount, colors)
	if len(matrix) != rowCount || len(matrix[0]) != columnCount {
		t.Error(fmt.Printf("Error when creating matrix in TestCreateMatrix func excepted row : %v recieved: %v , excepted column count %v, recieved %v", rowCount, len(matrix), columnCount, len(matrix[0])))
	}

	for _, row := range matrix {
		for _, v := range row {
			if !Find(colors, v) {
				t.Error(" error when checking colors in matrix in TestCreateMatrix func")
			}
		}
	}
}

func BenchmarkCreateMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const (
			rowCount    = 1
			columnCount = 1
		)
		colors := []int{BLACK, WHITE, RED}
		createMatrix(rowCount, columnCount, colors)
	}
}

func TestSearchInMatrix(t *testing.T) {
	colors := []int{BLACK, WHITE, RED}
	matrix := [][]int{
		{colors[0], colors[1], colors[1]},
		{colors[0], colors[1], colors[2]},
	}

	mapOfColorsWithCountOfContactPoints := searchInMatrix(matrix)
	if mapOfColorsWithCountOfContactPoints[BLACK] != 2 {
		t.Error(" error in searchInMatrix, excepted another count of data in your map")
	}

	if mapOfColorsWithCountOfContactPoints[WHITE] != 3 {
		t.Error(" error in searchInMatrix, excepted another count of data in your map")
	}
	if mapOfColorsWithCountOfContactPoints[RED] != 1 {
		t.Error(" error in searchInMatrix, excepted another count of data in your map")
	}
}

func BenchmarkSearchInMatrix(b *testing.B) {
	colors := []int{BLACK, WHITE, RED}
	matrix := [][]int{
		{colors[0], colors[1], colors[1]},
		{colors[0], colors[1], colors[2]},
	}
	for i := 0; i < b.N; i++ {
		searchInMatrix(matrix)
	}
}

func TestSearchTheBiggestColorsCountWithDoubleEqualResult(t *testing.T) {
	colorMap := make(map[int]int)
	colorMap[0] = 8
	colorMap[1] = 7
	colorMap[2] = 8
	mapOfColorsWithCountOfContactPoints := searchTheBiggestColorsCount(colorMap)

	if len(mapOfColorsWithCountOfContactPoints) != 2 {
		t.Error(fmt.Sprintf(" error in searchTheBiggestColorsCount, excepted length of array: %v recieved :%v", 2, len(mapOfColorsWithCountOfContactPoints)))
	}
}

func TestSearchTheBiggestColorsCount(t *testing.T) {
	colorMap := make(map[int]int)
	colorMap[0] = 8
	colorMap[1] = 17
	colorMap[2] = 1
	mapOfColorsWithCountOfContactPoints := searchTheBiggestColorsCount(colorMap)

	if len(mapOfColorsWithCountOfContactPoints) != 1 {
		t.Error(fmt.Sprintf(" error in searchTheBiggestColorsCount, excepted length of array: %v recieved :%v", 1, len(mapOfColorsWithCountOfContactPoints)))
	}
}

func BenchmarkSearchTheBiggestColorsCount(b *testing.B) {
	colorMap := make(map[int]int)
	colorMap[0] = 8
	colorMap[1] = 7
	colorMap[2] = 18
	for i := 0; i < b.N; i++ {
		searchTheBiggestColorsCount(colorMap)
	}
}
