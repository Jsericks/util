package util

import (
	"os"
	"strconv"
	"strings"
)

func FetchEnvStringWithDefault(k, d string) string {
	if s := os.Getenv(k); s != "" {
		return s
	}
	return d
}

func FetchEnvIntWithDefault(k string, d int) (int, error) {
	v := os.Getenv(k)
	if v == "" {
		return d, nil
	}
	i, err := strconv.Atoi(v)
	if err == nil {
		return i, err
	}
	return d, err
}

func FetchEnvStringAndSplitWithDefault(k, d string) []string {
	s := FetchEnvStringWithDefault(k, d)
	return strings.Split(s, ",")
}
