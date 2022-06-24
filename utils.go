package main

import "github.com/gookit/goutil/strutil"

func ToLowerCase(str string) string {
	return strutil.Lowercase(str)
}

func StringSlice(str string, start int, end int) string {
	return str[start:end]
}
