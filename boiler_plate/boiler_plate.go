package boiler_plate

import (
	"io/ioutil"
	"strings"
)

func Get(input string) string {
	bytes , err := ioutil.ReadFile(input)
	if err != nil { panic(err) }

	return string(bytes)
}

func GetC(input string) []string {
	return strings.Split(Get(input),"")
}

func GetN(input string) []string {
	return strings.Split(Get(input),"\n")
}

func GetNN(input string) []string {
	return strings.Split(Get(input),"\n\n")
}

func Ntos(input_strings []string) []string {
	for i, _ := range input_strings {
		for j, _ := range input_strings[i] {
			if string(input_strings[i][j]) == "\n" {
				out := []rune(input_strings[i])
				out[j] = ' '
				input_strings[i] = string(out);
			}
		}
	}

	return input_strings
}