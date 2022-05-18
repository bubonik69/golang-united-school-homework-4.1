package reverse_string

import "fmt"

func ReverseString(input string) (output string) {
	var temp []rune
	letters:= []rune(input)
	fmt.Println(letters)
	for i,_:= range letters{
		temp= append(temp, (letters[len(letters)-i-1]))
	}
	output= string(temp)

	return output
}
