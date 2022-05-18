package reverse_string

import "testing"


var tests=[]struct{
	input string
	output string
}{
	{"Hello world!", "!dlrow olleH"},
	{"123456789", "987654321"},
	{"Hello my chicken nuggets", "steggun nekcihc ym olleH"},
}


func TestReverseString (t *testing.T){
		for i:=0;i<len(tests);i++{
			if s:=ReverseString(tests[i].input); s!= tests[i].output{
				t.Errorf("Test with value `%s` is bad. Result is `%s` , need - '%s' ",tests[i].input,s, tests[i].output)
			}else{
				t.Logf("Test with value `%s` is ok ",tests[i].input)
			}

		}


}
