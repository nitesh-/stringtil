package stringtil

import (
	"strings"
//	"fmt"
)

// Reverses a string
func Reverse(str string) (string) {
	var _str string
	for i := len(str)-1; i > -1 ; i-- {
		_str += string(str[i])
	}
	return _str
}

// Repeats string n times
func Repeat(str string, count int) string {
	var _str string = str
	if count > 1 {
		_str += Repeat(str, count-1)
	} else {
		_str = ""
	}
	return _str
}

// Replaced the first letter of the string to uppercase
func Ucfirst(str string) (string) {
	var _s string = strings.ToUpper(str[0:1]) + str[1:]
	return _s
}

// Replaces the first letter of each word to Uppercase
func Ucwords(str string) (string) {
	var _str string
	var strArr []string
	strArr = strings.Split(str, " ")
	for i := 0; i < len(strArr); i++ {
		strArr[i] = Ucfirst(strArr[i])
	}
	_str = strings.Join(strArr, " ")
	return _str
}

// Returns substring of string
func Substr(str string, startIndex int, length ...int) (string) {
	var _strLen int = len(str)
	if len(length) == 0 {
		return str[startIndex:]
	} else {
		if length[0] > _strLen {
			length[0] = _strLen
		}
		return str[startIndex:length[0]]
	}
}

// Splits the query params and returns in Map format
func ParseUrl(queryString string) map[string]string {
	var paramArray []string = strings.Split(queryString, "&")
	var paramArrayLen int = len(paramArray)
	var fieldArray []string
	op := make(map[string]string)
	for i := 0; i < paramArrayLen; i++ {
		fieldArray = strings.Split(paramArray[i], "=")
		op[fieldArray[0]] = fieldArray[1]
	}
	return op
}

// Finds the position of needle in haystack
func Strpos(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// Translates certain characters in a string
func Strtr(str string, fromStr string, toStr string) string {
	var lenToStr int = len(toStr)
	for i, _ := range fromStr {
		if i < lenToStr {
			str = strings.ReplaceAll(str, string(fromStr[i]), string(toStr[i]))
		}
	}
	return str
}
