package reverse_string

func ReverseString(input string) (output string) {
	var r []rune
	letters:= []rune(input)
	for i,_:= range letters{
		r= append(r, (letters[len(letters)-i-1]))
	}
	return string(r)
}
