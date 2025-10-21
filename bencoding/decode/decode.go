package decode

func Decode(input string) any {
	res, _ := GetValue(input, 0)
	return res
}

func DecodeString(input string, idx int) (string, int) {
	stringLength := 0
	for input[idx] != byte(':') {
		stringLength = ((stringLength * 10) + int(input[idx]-'0'))
		idx += 1
	}
	start := idx + 1
	res_str := input[start : start+stringLength]
	return res_str, start + stringLength
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
		var list_item any
		list_item, idx = GetValue(input, idx)
		res = append(res, list_item)
	}
	return res, idx + 1
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

func GetValue(input string, idx int) (any, int) {

	switch input[idx] {
	case 'l':
		return DecodeList(input, idx+1)

	case 'i':
		return DecodeInteger(input, idx+1)

	case 'd':
		return DecodeDictionary(input, idx+1)

	default:
		return DecodeString(input, idx)

	}
}
