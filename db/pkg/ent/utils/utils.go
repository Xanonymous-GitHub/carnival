package utils

import (
	"fmt"
	"strings"
)

// Wait golang 1.18 release to use generic features.

func ToValuesFromEnum(enums ...interface{}) (kinds []string) {
	s := fmt.Sprint(enums[0])
	f := strings.Fields(s[1 : len(s)-1])
	for _, k := range f {
		kinds = append(kinds, k)
	}
	return
}

func ToMapFromEnum(enums ...interface{}) map[string]int32 {
	r := make(map[string]int32)
	s := fmt.Sprint(enums[0])
	f := strings.Fields(s[1 : len(s)-1])
	for i, k := range f {
		r[k] = int32(i) + 1
	}
	return r
}
