package tools

import (
	"regexp"
	"strconv"
)

func RemoveMask(document *string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return re.ReplaceAllString(*document, "")
}

func ConvertStrToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	return num, err
}
