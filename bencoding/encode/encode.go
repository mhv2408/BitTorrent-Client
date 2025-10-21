package encode

import (
	"fmt"
	"sort"
)

func Encode(input any) string {

	switch val := input.(type) {
	case string:
		return EncodeString(val)
	case int:
		return EncodeInteger(val)
	case []any:
		return EncodeList(val)
	case map[string]any:
		return EncodeMap(val)
	default:
		return ""
	}
}
func EncodeString(input string) string {
	return fmt.Sprintf("%d:%s", len(input), input)
}
func EncodeInteger(input any) string {
	return fmt.Sprintf("i%de", input)
}
func EncodeList(input []any) string {

	res := ""
	for _, element := range input {
		res += Encode(element)
	}

	return fmt.Sprintf("l%se", res)
}

func EncodeMap(input map[string]any) string {
	res := ""
	var Keys []string
	for key := range input {
		Keys = append(Keys, key)
	}
	sort.Strings(Keys)

	for _, Key := range Keys {
		res += EncodeString(Key)
		res += Encode(input[Key])
	}

	return fmt.Sprintf("d%se", res)
}
