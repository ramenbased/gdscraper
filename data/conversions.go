package data

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Er(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func emptyString(s string) string {
	if s == "" {
		return "NULL"
	}
	if s != "" {
		log.Fatalf("emptyString in data conversions Error cus of this shit: %v", s)
	}
	return s
}

func c_NoSpace(s string) string {
	return strings.TrimSpace(s)
}

func c_WeekInt(s string) any {
	r, err := regexp.Compile("\\d{1,}")
	Er(err)
	i, err := strconv.Atoi(r.FindString(s))
	if err != nil {
		return emptyString(s)
	}
	return i
}

func c_StringFloat(s string) any {
	rs := strings.TrimSpace(s)
	f, err := strconv.ParseFloat(rs, 64)
	if err != nil {
		//TODO: return sql.Null something
		return emptyString(s)
	}
	return f
}
