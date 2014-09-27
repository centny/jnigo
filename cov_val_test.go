package jnigo

import (
	"fmt"
	"testing"
	"time"
)

func args_t(t *testing.T, tsig string, args ...interface{}) {
	sig, _, err := CovArgs(args...)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(sig + "--->" + tsig)
	if sig != tsig {
		t.Error(sig + " is not " + tsig)
		// os.Exit(1)
	}
}

func TestCovArgs(t *testing.T) {
	args_t(t, "II", 1, 1)
	args_t(t, "IZ", 1, true)
	args_t(t, "IZ", 1, false)
	args_t(t, "ZBCSIJFD", true, Byte(1), Char(1), int16(1), 1, int64(1), float32(1), float64(1))
	args_t(t, "ZBCSIJFD", false, uint8(1), Char(1), int16(1), int32(1), int64(1), float32(1), float64(1))
	args_t(t, "[Z[B[C[S[I[J[F[D", []bool{true}, []Byte{1}, []Char{1}, []int16{1}, []int{1}, []int64{1}, []float32{1}, []float64{1})
	args_t(t, "[Z[B[C[S[I[J[F[D", []bool{false}, []uint8{1}, []Char{1}, []int16{1}, []int32{1}, []int64{1}, []float32{1}, []float64{1})
	args_t(t, "[Z[B[C[S[I[J[F[D", []bool{}, []Byte{}, []Char{}, []int16{}, []int{}, []int64{}, []float32{}, []float64{})
	args_t(t, "[Z[B[C[S[I[J[F[D", []bool{}, []uint8{}, []Char{}, []int16{}, []int32{}, []int64{}, []float32{}, []float64{})
	args_t(t, "Ljava/lang/String;", "abc")
	args_t(t, "Ljava/lang/String;Ljava/lang/String;", "abc1", "abc2")
	args_t(t, "[Ljava/lang/String;Ljava/lang/String;", []string{"abc1"}, "abc2")
	obj1 := JNIGO_newS("jjjj1")
	args_t(t, "Ljava/lang/String;Ljava/lang/String;", obj1, &obj1)
	// args_t(t, "[Ljava.lang.String;", []EmptyObjAry{"java.lang.String"})
	//
	fmt.Println(covary(1))
	fmt.Println(covary(nil))
	// fmt.Println(covary([]EmptyObjAry{}))
	// fmt.Println(covary([]EmptyObjAry{"asss"}))
	fmt.Println(covary([]Object{}))
	fmt.Println(covary([]Object{Object{}}))
	fmt.Println(covary([]*Object{}))
	fmt.Println(covary([]*Object{&Object{}}))
	fmt.Println(covary([]string{""}))
	//
	fmt.Println(covval("arg"))
	fmt.Println(CovArgs("arg"))
	fmt.Println(CovArgs(time.Now()))
	fmt.Println(CovArgs([]time.Time{time.Now()}))
	//
	fmt.Println("TestCovArgs-------->")
}
func TestNewAry(t *testing.T) {
	fmt.Println(NewAryv("jsjfks", nil, 0))
}
