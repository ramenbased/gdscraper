package main

import (
	"regexp"
)

func regexGetID(url string) string {
	c, err := regexp.Compile("[0-9]{3,7}")
	Er(err)
	return c.FindString(url)
}
