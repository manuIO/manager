package cockroachdb

import (
	"fmt"
	"strconv"
	"strings"
)

const sep string = ","

func toString(slice []uint) string {
	text := []string{}
	for _, v := range slice {
		text = append(text, fmt.Sprintf("%d", v))
	}

	return strings.Join(text, sep)
}

func fromString(text string) []uint {
	chans := []uint{}

	parts := strings.Split(text, sep)
	for _, v := range parts {
		val, _ := strconv.ParseUint(v, 0, 0)
		chans = append(chans, uint(val))
	}

	return chans
}
