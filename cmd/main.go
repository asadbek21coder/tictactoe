package main

import "fmt"

func main() {
	var matrix [3][3]string = [3][3]string{
		[3]string{"1", "2", "3"},
		[3]string{"4", "5", "6"},
		[3]string{"7", "8", "9"},
	}
	var matrix2 [8][3]int = [8][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9},
		[3]int{1, 4, 7},
		[3]int{2, 5, 8},
		[3]int{3, 6, 9},
		[3]int{1, 5, 9},
		[3]int{3, 5, 7},
	}

	var row map[string]int = map[string]int{
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 1,
		"5": 1,
		"6": 1,
		"7": 2,
		"8": 2,
		"9": 2,
	}
	var column map[string]int = map[string]int{
		"1": 0,
		"2": 1,
		"3": 2,
		"4": 0,
		"5": 1,
		"6": 2,
		"7": 0,
		"8": 1,
		"9": 2,
	}
	var i int = 1
	for i <= 9 {
		var input string
		fmt.Scan(&input)
		// fmt.Println(mos[input], mos2[input])
		x, y := row[input], column[input]
		if i%2 == 0 {
			matrix[x][y] = "0"
		} else {
			matrix[x][y] = "x"
		}
		Display(matrix)
		var win bool = false
		for j := 0; j < 8; j++ {
			if matrix[matrix2[j][0]] == matrix[matrix2[j][1]] && matrix[matrix2[j][0]] == matrix[matrix2[j][2]] {
				win = true
				break
			}
		}
		if win {
			if i%2 == 0 {
				fmt.Println("2-O'yinchi yutdi!")
			} else {
				fmt.Println("1-O'yinchi yutdi!")
			}
			break
		}
		i++
	}
}

func Display(matrix [3][3]string) {

	fmt.Println("-------------------------------------------------")

	fmt.Printf("\t%v \t| \t%v \t| \t%v \t|\n", matrix[0][0], matrix[0][1], matrix[0][2])

	fmt.Println("-------------------------------------------------")

	fmt.Printf("\t%v \t| \t%v \t| \t%v \t|\n", matrix[1][0], matrix[1][1], matrix[1][2])

	fmt.Println("-------------------------------------------------")

	fmt.Printf("\t%v \t| \t%v \t| \t%v \t|\n", matrix[2][0], matrix[2][1], matrix[2][2])

}
