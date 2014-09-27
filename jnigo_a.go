package jnigo

import (
	"regexp"
	"strings"
)

// //
var sig_m *regexp.Regexp = regexp.MustCompile("^[ZBCSIJFDLV\\[].*$")

func SigName(name string) string {
	name = strings.Trim(name, "\t ")
	if len(name) < 1 {
		return name
	}
	name = strings.Replace(name, ".", "/", -1)
	if sig_m.MatchString(name) {
		return name
	} else {
		return "L" + name + ";"
	}
}

func New(name string, args ...interface{}) (Object, error) {
	cls, err := FindClass(name)
	if err != nil {
		return Object{}, err
	}
	return cls.New(args...)
}
