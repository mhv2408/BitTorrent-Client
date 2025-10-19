package decode

import "unicode"

func DecodeString(input string, idx int) (string, int) {
	len := 0
	for input[idx] != byte(':') {
		len = (len * 10) + int(input[idx])
		idx += 1
	}
	start := idx + 1
	res_str := input[start : start+len]
	return res_str, start + len
}
func DecodeInteger(input string, idx int) (int, int) {
	res_int := 0
	for input[idx] != byte('e') {
		res_int = res_int*10 + int(input[idx])
		idx += 1
	}
	return res_int, idx + 1
}
func DecodeList(input string, idx int) ([]interface{}, int) {

	var res []interface{}
	for input[idx] != byte('e') {
		symbol := input[idx]
		var list_item interface{}
		if unicode.IsDigit(rune(symbol)) { // it is string lenght starting
			list_item, idx = DecodeString(input, idx)
		} else if symbol == byte('i') { // it is an integer
			list_item, idx = DecodeInteger(input, idx)
		} else if symbol == byte('d') {
			list_item, idx = DecodeDictionary(input, idx)
		} else { // it is list itself
			list_item, idx = DecodeList(input, idx)
		}

		res = append(res, list_item)

	}
	return res, idx
}

func GetValue(input string, idx int) (interface{}, int) {
	for true {
		element := input[idx]

		if unicode.IsDigit(rune(element)) {
			return DecodeString(input, idx)
		} else if element == byte('i') {
			return DecodeInteger(input, idx)
		} else if element == byte('l') {
			return DecodeList(input, idx)
		} else {
			return DecodeDictionary(input, idx)
		}
	}
	return nil, -1
}

func DecodeDictionary(input string, idx int) (map[string]interface{}, int) {
	var res map[string]interface{} // in bencoding there are only string keys

	for input[idx] != byte('e') {
		key, idx := DecodeString(input, idx)
		value, idx := GetValue(input, idx)

		res[key] = value
	}

	return res, idx

}
