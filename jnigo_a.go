package jnigo

import (
	"regexp"
	"strings"
)

// //
var sig_m *regexp.Regexp = regexp.MustCompile("^[ZBCSIJFDLV\\[].*$")

func SigName(name string) string {
	name = strings.Trim(name, "\t ")
	name = strings.Replace(name, ".", "/", -1)
	if sig_m.MatchString(name) {
		return name
	} else {
		return "L" + name + ";"
	}
}
