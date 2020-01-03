package stringtil

import (
	"strings"
)

// @method Reverse
// @description Reverses a string
// @param str string
// @return string Returns the reversed string
func Reverse(str string) (string) {
	var _str string
	for i := len(str)-1; i > -1 ; i-- {
		_str += string(str[i])
	}
	return _str
}

// @method Repeat
// @description Repeats string n times
// @param str string String to be repeated
// @param count int No of times the string should be repeated
func Repeat(str string, count int) string {
	var _str string = str
	if count > 1 {
		_str += Repeat(str, count-1)
	} else {
		_str = ""
	}
	return _str
}

// @method Ucfirst
// @description Replaced the first letter of the string to uppercase
// @param str string String to be repeated
// @return string Returns the Converted string
func Ucfirst(str string) (string) {
	var _s string = strings.ToUpper(str[0:1]) + str[1:]
	return _s
}

// @method Ucwords
// @description Replaces the first letter of each word to Uppercase
// @param str string String to be repeated
// @return string Returns the Converted string
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

// @method Substr
// @description Returns substring of string
// @param str string haystack
// @param startIndex int Start Index
// @param length int Length
// @return string Returns the substring
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

// @method ParseUrl
// @description Splits the query params and returns in Map format
// @param str string query string
// @return map Returns the field map
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

func 