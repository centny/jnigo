package jnigo

import (
	"fmt"
	"testing"
)

func init() {
	JNIGO_addOption("-Djava.class.path=src/java/bin")
	fmt.Println(JNIGO_init())
}
func TestStr(t *testing.T) {
	tv := "abc--"
	str := JNIGO_newS(tv)
	tstr, _ := str.String()
	if tv != tstr {
		t.Error("not right")
	}
}

// func TestStaticM(t *testing.T) {
// 	cls, err := JNIGO_findClass("jnigo/StaticM")
// }
