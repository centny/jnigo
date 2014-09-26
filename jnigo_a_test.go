package jnigo

import (
	"fmt"
	"testing"
)

func TestSigName(t *testing.T) {
	fmt.Println(SigName("java/lang/String"))
	fmt.Println(SigName("java.lang.String"))
	fmt.Println(SigName("Ljava/lang/String;"))
}
