package main

func Increment(value *int) {
	if value == nil {
		return
	}

	*value += 1
}

func main() {
}
