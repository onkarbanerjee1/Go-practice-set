package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// a = matrix from given file, b= a's transpose, c= a*b
	var a [10][10]int
	var b [10][10]int
	var c [10][10]int

	n := InitMatrices(&a, &b, f)

	fmt.Println("\n The given matrix is :")
	Display(&a, n)

	fmt.Println("\n\n It's transpose is :")
	Display(&b, n)

	Multiply(&a, &b, &c, n)
	fmt.Println("\n\n Their product is :")
	Display(&c, n)

}

// function to read the fiile and initiate the matirx and it's transpose and validate if squarre as well
func InitMatrices(a *[10][10]int, b *[10][10]int, f *os.File) int {
	row := 0
	col := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if col == 0 {
			col = len(line)
		} else if col != len(line) {
			log.Fatal("Input not in valid 2D matrix format")
		}
		for j, k := range line {
			a[row][j], _ = strconv.Atoi(k)
			b[j][row] = a[row][j]
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// check if square or not
	if row != col {
		log.Fatal("Not a square matrix")
	}
	return row
}

// function to calculate the product of two given matrices and store into another given matrix
func Multiply(a *[10][10]int, b *[10][10]int, c *[10][10]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = 0
			for k := 0; k < n; k++ {
				c[i][j] = c[i][j] + (a[i][k] * b[k][j])
			}

		}
	}
}

// function to display the contents of the matrix
func Display(m *[10][10]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf(" %d", m[i][j])
		}
		fmt.Println()
	}
}
