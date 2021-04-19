package main

func Repeat(text string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += text
	}
	return result
}
