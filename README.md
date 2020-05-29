
# stringtil
A small string utils library which contains PHP inbuilt functions written in Go.

Functions:
--------------------
Reverse - Reverses a string

`stringtil.Reverse(str string)`

Repeat - Repeats a string n times

`stringtil.Repeat(str string, count int)`

Ucfirst - Converts the first letter of the string to uppercase

`stringtil.Ucfirst(str string)`

Ucwords - Converts the first letter of each word to uppercase

`stringtil.Ucwords(str string)`

Substr - Returns substring of string

`stringtil.Substr(str string, startIndex int, length ...int)`

ParseUrl - Splits the query params and returns in Map format

`stringtil.ParseUrl(queryString string)`

Strtr - Translates certain characters in a string

`stringtil.Strtr(str string, fromStr string, toStr string)`
