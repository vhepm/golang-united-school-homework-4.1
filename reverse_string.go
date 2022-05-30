package reverse_string

import "fmt"

func ReverseString(input string) (output string) {
	n := 0
	split := make([]rune, len(input))
	for _, r := range input {
		split[n] = r
		n++
	}
	split = split[0:n]

	for i := 0; i < n/2; i++ {
		split[i], split[n-1-i] = split[n-1-i], split[i]
	}

	output = string(split)
	fmt.Println(output)

	return
}
