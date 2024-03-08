package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type QueryOperator string

const pattern = `\w+:[^:]+(?: |$)`

func IsValidFilter(input string) bool {
	matched, err := regexp.MatchString(pattern, input)
	if err != nil {
		return false
	}

	if matched {
		return true
	}

	return false
}

func queryOperatorByValueType(value string) QueryOperator {
	if strings.HasPrefix(value, "0") {
		return "in"
	}

	if strings.ToLower(value) == "true" || strings.ToLower(value) == "false" {
		return "eq"
	}

	if isNumber(value) {
		return "eq"
	}

	return "in"
}

func isNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func BuildFilterQuery(input string) string {
	keyValuePairs := regexp.MustCompile(pattern).FindAllString(input, -1)
	var finalString string

	for _, pair := range keyValuePairs {
		keyValue := strings.Split(pair, ":")
		key := keyValue[0]
		value := keyValue[1]
		operator := queryOperatorByValueType(value)
		finalString += fmt.Sprintf("%s:%s(%s),", key, operator, value)
	}

	finalString = strings.TrimSuffix(finalString, ",")

	return finalString
}
