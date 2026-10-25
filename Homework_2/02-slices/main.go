package main

func Sum(numbers []int) int {
	summa := 0
	for _, number := range numbers {
		summa += number
	}
	return summa
}

func main() {
}
