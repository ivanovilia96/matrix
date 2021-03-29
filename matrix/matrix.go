package main

import (
	"fmt"
	"math/rand"
)

const (
	BLACK = iota
	WHITE
	RED
)

// используется линейный поиск потому что массив не отсортирован. так можно было бы бинарный (если был бы отсортирован)
func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// рекурсивная функция, которая ищет с 4х сторон от элемента, если находит такой же цвет, то вызывает сама себя но уже с того элемента который находит
func checkNearElements(matrix [][]int, elementPositionRow, elementPositionColumn int, counter *int, checkedElements map[int][]int) {
	curentCheckingColor := matrix[elementPositionRow][elementPositionColumn]
	// проверка на то, что этот элемент мы уже проверяли, если проверяли, то выход из рекурсии
	if val, isDefaultValue := checkedElements[elementPositionRow]; !isDefaultValue || !Find(val, elementPositionColumn) {
		// т.к. мы его сейчас проверяем, то добавляем в проверенные элементы
		checkedElements[elementPositionRow] = append(checkedElements[elementPositionRow], elementPositionColumn)
		*counter += 1
		// элемент находится не на первой линии, тогда есть возможность проверить элемент над ним с помощью рекурсии
		if elementPositionRow != 0 {
			if curentCheckingColor == matrix[elementPositionRow-1][elementPositionColumn] {
				checkNearElements(matrix, elementPositionRow-1, elementPositionColumn, counter, checkedElements)
			}
		}

		// элемент находится не на последней линии, тогда есть возможность проверить элемент под ним с помощью рекурсии
		if elementPositionRow != len(matrix)-1 {
			if curentCheckingColor == matrix[elementPositionRow+1][elementPositionColumn] {
				checkNearElements(matrix, elementPositionRow+1, elementPositionColumn, counter, checkedElements)
			}
		}

		// элемент является не первым в row, тогда есть возможность проверить элемент перед ним с помощью рекурсии
		if elementPositionColumn != 0 {
			if curentCheckingColor == matrix[elementPositionRow][elementPositionColumn-1] {
				checkNearElements(matrix, elementPositionRow, elementPositionColumn-1, counter, checkedElements)
			}
		}

		// элемент является не последним в row, тогда есть возможность проверить элемент после него с помощью рекурсии
		if elementPositionColumn != len(matrix[0])-1 {
			if curentCheckingColor == matrix[elementPositionRow][elementPositionColumn+1] {
				checkNearElements(matrix, elementPositionRow, elementPositionColumn+1, counter, checkedElements)
			}
		}
	}
}

// создает матрицу и возвращает её
func createMatrix(rowCount, columnCount int, colors []int) [][]int {
	var matrix [][]int
	for i := 0; i < rowCount; i++ {
		matrix = append(matrix, []int{})
		for j := 0; j < columnCount; j++ {
			matrix[i] = append(matrix[i], colors[rand.Intn(len(colors))])
		}
	}
	return matrix
}

func searchInMatrix(matrix [][]int) map[int]int {
	//фрагмент поиска в матрице (будет содержать цвет: макс кол-во ребер )
	mapOfBlockCount := make(map[int]int)
	/*
		 храню checkedElements что бы не ходить по кругу на рекурсии
		 ( в эту переменную функция checkNearElements автоматически будет записывать элементы ) служит и как выход из рекурсии и как механизм оптимизации
		скорости по средствам кеширования
	*/
	checkedElements := make(map[int][]int)
	// идем по row
	for i := 0; i < len(matrix); i++ {
		// идем по каждому элементу в row
		for j := 0; j < len(matrix[0]); j++ {
			blocksCount := 0
			checkNearElements(matrix, i, j, &blocksCount, checkedElements)
			// если значение под цвет уже присутствует, то мы сравниваем его с текущим. если нет, записываем новое т.к есть значение по умолчанию у типов
			if mapOfBlockCount[matrix[i][j]] < blocksCount {
				mapOfBlockCount[matrix[i][j]] = blocksCount
			}
		}
	}
	return mapOfBlockCount
}

// ищет из мэп наибольшие значения, возвращает список этих значений ( если у нескольких цветов будет одинаковое кол-во пересечений, вернет несколько )
func searchTheBiggestColorsCount(colorMap map[int]int) []string {
	biggestCountValue := 0

	// ищем наибольшее значение
	for _, value := range colorMap {
		if biggestCountValue < value {
			biggestCountValue = value
		}
	}

	listOfBiggestColors := []int{}
	// ищем дубликаты наибольшего значения
	for key, value := range colorMap {
		if biggestCountValue == value {
			listOfBiggestColors = append(listOfBiggestColors, key)
		}
	}
	resultSetOfBiggestNumbers := []string{}

	for i := 0; i < len(listOfBiggestColors); i++ {
		switch listOfBiggestColors[i] {
		case BLACK:
			resultSetOfBiggestNumbers = append(resultSetOfBiggestNumbers, fmt.Sprintf("color: Black кол-во пересечений: %v \n", colorMap[listOfBiggestColors[i]]))
		case WHITE:
			resultSetOfBiggestNumbers = append(resultSetOfBiggestNumbers, fmt.Sprintf("color: White кол-во пересечений: %v \n", colorMap[listOfBiggestColors[i]]))
		case RED:
			resultSetOfBiggestNumbers = append(resultSetOfBiggestNumbers, fmt.Sprintf("color: Red кол-во пересечений: %v \n", colorMap[listOfBiggestColors[i]]))
		default:
			println("этого не должно было вывестись, если вы добавили цвет в searchTheBiggestColorCount функцию в switch")
		}
	}

	return resultSetOfBiggestNumbers
}

func main() {
	const (
		rowCount    = 15
		columnCount = 5
	)

	// если сюда добавить новые цвета, то их нужно будет добавить в searchTheBiggestColorCount функцию в switch
	colors := []int{BLACK, WHITE, RED}

	matrix := createMatrix(rowCount, columnCount, colors)
	mapOfColorsWithCountOfContactPoints := searchInMatrix(matrix)
	result := searchTheBiggestColorsCount(mapOfColorsWithCountOfContactPoints)

	for i := 0; i < rowCount; i++ {
		fmt.Printf("%v  \n", matrix[i])
	}
	fmt.Printf("%v  - Аналитика для справки (ключ - enum цвета, значение - кол-во граней)\n", mapOfColorsWithCountOfContactPoints)
	fmt.Printf("%v  - это цвет\\цвета с наибольшими пересечениями \n", result)
}
