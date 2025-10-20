package decode

import (
	"unicode"
)

func Decode(input string) any {
	var resultList []any
	i := 0
	for i < len(input) {
		var resultItem any
		symbol := input[i]
		if unicode.IsDigit(rune(symbol)) {
			resultItem, i = DecodeString(input, i)
		} else if symbol == byte('i') {
			resultItem, i = DecodeInteger(input, i+1)
		} else if symbol == byte('l') {
			resultItem, i = DecodeList(input, i+1)
		} else {
			resultItem, i = DecodeDictionary(input, i+1)
		}
		resultList = append(resultList, resultItem)
	}
	return resultList
}

func DecodeString(input string, idx int) (string, int) {
	len := 0
	for input[idx] != byte(':') {
		len = ((len * 10) + int(input[idx]-'0'))
		idx += 1
	}
	start := idx + 1
	res_str := input[start : start+len]
	return res_str, start + len
}
func DecodeInteger(input string, idx int) (int, int) {
	res_int := 0
	is_neg := false
	if input[idx] == '-' {
		is_neg = true
		idx += 1
	}
	for input[idx] != byte('e') {
		res_int = res_int*10 + int(input[idx]-'0')
		idx += 1
	}
	if is_neg {
		res_int *= -1
	}
	return res_int, idx + 1
}
func DecodeList(input string, idx int) ([]any, int) {

	var res []any
	for input[idx] != byte('e') {
		symbol := input[idx]
		var list_item any
		if unicode.IsDigit(rune(symbol)) { // it is string lenght starting
			list_item, idx = DecodeString(input, idx)
		} else if symbol == byte('i') { // it is an integer
			list_item, idx = DecodeInteger(input, idx+1)
		} else if symbol == byte('d') {
			list_item, idx = DecodeDictionary(input, idx+1)
		} else { // it is list itself
			list_item, idx = DecodeList(input, idx+1)
		}

		res = append(res, list_item)
	}
	return res, idx + 1
}

func GetValue(input string, idx int) (any, int) {
	element := input[idx]

	if unicode.IsDigit(rune(element)) {
		return DecodeString(input, idx)
	} else if element == byte('i') {
		return DecodeInteger(input, idx+1)
	} else if element == byte('l') {
		return DecodeList(input, idx+1)
	} else {
		return DecodeDictionary(input, idx+1)
	}
}

func DecodeDictionary(input string, idx int) (map[string]any, int) {
	// in bencoding there are only string keys
	res := make(map[string]any)

	for input[idx] != byte('e') {
		var key string
		var value any
		key, idx = DecodeString(input, idx)
		value, idx = GetValue(input, idx)
		res[key] = value
	}
	return res, idx + 1

}
