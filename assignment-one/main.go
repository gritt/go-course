package main

func main() {

	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range i {

		print(value)

		if value%2 == 0 {
			print(" is even \n")
			continue
		}

		print(" is odd \n")
	}
}
