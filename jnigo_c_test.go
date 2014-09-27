package jnigo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"testing"
)

func init() {
	SetVersion(JNI_VERSION_1_6)
	SetIgnoreUnrecognized(0)
	AddOption("-Djava.class.path=java/bin")
	fmt.Println(Init())
}
func TestStr(t *testing.T) {
	tv := "abc--"
	str := NewS(tv)
	tstr, err := str.String()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(tstr)
	if tv != tstr {
		t.Error("not right")
		os.Exit(1)
	}
	// fmt.Println(tstr)
}

func TestStaticM(t *testing.T) {
	cls, err := FindClass("jnigo/StaticM")
	if err != nil {
		t.Error(err.Error())
		return
	}
	cls.CallV("showz", "V", false)
	cls.CallV("showz", "V", false)
	cls.CallV("showb", "V", uint8(1))
	cls.CallV("showb", "V", Byte(1))
	cls.CallV("showc", "V", Char(1))
	cls.CallV("shows", "V", int16(1))
	cls.CallV("showi", "V", 1)
	cls.CallV("showj", "V", int64(1))
	cls.CallV("showf", "V", float32(1))
	cls.CallV("showd", "V", float64(1))
	// //
	obj, err := cls.CallV("getv", "V")
	obj, err = cls.CallV("getz", "Z")
	fmt.Println(obj.Value())
	obj, _ = cls.CallV("getb", "B")
	fmt.Println(obj.Value())
	obj, _ = cls.CallV("getc", "C")
	fmt.Println(obj.Value())
	obj, _ = cls.CallV("gets", "S")
	fmt.Println(obj.Value())
	obj, _ = cls.CallV("geti", "I")
	fmt.Println(obj.Value())
	obj, _ = cls.CallV("getj", "J")
	fmt.Println(obj.Value())
	obj, _ = cls.CallV("getf", "F")
	fmt.Println(obj.Value())
	obj, _ = cls.CallV("getd", "D")
	fmt.Println(obj.Value())
	// //	//
	fmt.Println("TestStaticM-------->")
}

func TestObjectM(t *testing.T) {
	cls, err := FindClass("jnigo/ObjectM")
	if err != nil {
		t.Error(err.Error())
		return
	}
	// fmt.Println(cls)
	// fmt.Println(cls.New())
	// fmt.Println(cls.New("..."))
	nobj, err := cls.New(1, 2)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(nobj.jvm, "jvm")
	fmt.Println(nobj.Val())
	fmt.Println("-------xxxxx<------")
	obj, err := nobj.CallV("getA", "I")
	fmt.Println("------->011")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("------->021")
	fmt.Println(obj.I())
	fmt.Println("------->031")
	fmt.Println(obj.jvm)
	obj, err = nobj.CallV("getB", "I")
	// fmt.Println(obj)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(obj.I())
	nobj.CallV("setA", "V", 221)
	nobj.CallV("setB", "V", 222)
	obj, err = nobj.CallV("getA", "I")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(obj.I())
	obj, err = nobj.CallV("getB", "I")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(obj.I())
	//
	//
	// obj.CallVoid("setA", 211)
	// obj.CallVoid("setB", 222)
	// fmt.Println(obj.CallInt("getA"))
	// fmt.Println(obj.CallInt("getB"))
	// //
	// //
	// obj.CallVoid("showz", false)
	// obj.CallVoid("showb", uint8(1))
	// obj.CallVoid("showb", Byte(1))
	// obj.CallVoid("showc", Char(1))
	// obj.CallVoid("shows", int16(1))
	// obj.CallVoid("showi", 1)
	// obj.CallVoid("showj", int64(1))
	// obj.CallVoid("showf", float32(1))
	// obj.CallVoid("showd", float64(1))
	// //
	// fmt.Println(obj.CallVoid("getv"))
	// fmt.Println(obj.CallBoolean("getz"))
	// fmt.Println(obj.CallByte("getb"))
	// fmt.Println(obj.CallChar("getc"))
	// fmt.Println(obj.CallShort("gets"))
	// fmt.Println(obj.CallInt("geti"))
	// fmt.Println(obj.CallLong("getj"))
	// fmt.Println(obj.CallFloat("getf"))
	// fmt.Println(obj.CallDouble("getd"))
	// //
	// fmt.Println(obj.CallVoid("get"))
	// fmt.Println(obj.CallBoolean("get"))
	// fmt.Println(obj.CallByte("get"))
	// fmt.Println(obj.CallChar("get"))
	// fmt.Println(obj.CallShort("get"))
	// fmt.Println(obj.CallInt("get"))
	// fmt.Println(obj.CallLong("get"))
	// fmt.Println(obj.CallFloat("get"))
	// fmt.Println(obj.CallDouble("get"))
	// //
	// fmt.Println(obj.CallVoid("getv", ""))
	// fmt.Println(obj.CallBoolean("getz", ""))
	// fmt.Println(obj.CallByte("getb", ""))
	// fmt.Println(obj.CallChar("getc", ""))
	// fmt.Println(obj.CallShort("gets", ""))
	// fmt.Println(obj.CallInt("geti", ""))
	// fmt.Println(obj.CallLong("getj", ""))
	// fmt.Println(obj.CallFloat("getf", ""))
	// fmt.Println(obj.CallDouble("getd", ""))
	//
	fmt.Println("TestObjectM-------->")
}
func TestArray(t *testing.T) {
	// 	vm := GVM
	clsa, err := FindClass("jnigo/A")
	if err != nil {
		t.Error(err.Error())
		return
	}
	clsary, err := FindClass("jnigo/Ary")
	if err != nil {
		t.Error(err.Error())
		return
	}
	obja, _ := clsa.New()
	objb, _ := clsa.New()
	_, err = clsary.CallV("showas", "V", []Object{obja, objb})
	if err != nil {
		t.Error(err.Error())
		return
	}
	// vm.NewObjectArray()
	objary, _ := clsary.New()
	_, err = objary.CallV("show", "V", []*Object{&obja, &objb})
	if err != nil {
		t.Error(err.Error())
		return
	}
	//
	fmt.Println("TestArray-------->")
}
func TestAshow(t *testing.T) {
	clsa, _ := FindClass("Ljnigo/A;")
	obja, _ := clsa.New()
	objv, _ := obja.As("Ljava/lang/Object;")
	obja.CallV("show", "V", objv)
	obja.CallV("show", "V", []Object{objv})
	fmt.Println()
	fmt.Println()
	show, err := obja.CallV("show", "Ljava/lang/String;",
		true,              //for java boolean
		Byte(1),           //for java byte
		Char(1),           //for java char
		int16(1),          //for java short
		1,                 //for java int
		int64(1),          //for java long
		float32(1),        //for java float
		float64(1),        //for java double
		objv,              //for java Object
		"jjjjj",           //for java String
		[]bool{false},     //for java boolean[]
		[]Byte{1, 2},      //for java byte[]
		[]Char{3, 4},      //for java char[]
		[]int16{11, 12},   //for java short[]
		[]int{21, 22},     //for java int[]
		[]int64{31, 32},   //for java long[]
		[]float32{41, 42}, //for java float[]
		[]float64{51, 52}, //for java double[]
		[]Object{objv},    //for java Object[]
		[]string{"aaa"},   //for java String[]
	)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(show.ToString())
	fmt.Println("TestAshow-------->")
}
func TestErr(t *testing.T) {
	clsa, err := FindClass("Ljnigo/Err;")
	if err != nil {
		t.Error(err.Error())
		return
	}
	_, err = clsa.CallV("err1", "V")
	if err == nil {
		t.Error("not error")
		return
	}
	fmt.Println(err.Error())
	//
	_, err = clsa.CallV("err2", "V")
	if err == nil {
		t.Error("not error")
		return
	}
	fmt.Println(err.Error())
	//
	_, err = clsa.CallV("err3", "V")
	if err == nil {
		t.Error("not error")
		return
	}
	fmt.Println(err.Error())
}

func TestInput(t *testing.T) {
	in, err := New("java.io.FileInputStream", "jnigo.cb")
	if err != nil {
		t.Error(err.Error())
		return
	}
	err = in.CallVoid("close")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestInput2(t *testing.T) {
	db, err := sql.Open("mysql", "cny:123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		t.Error(err.Error())
		return
	}
	tx, err := db.Begin()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("exec----010")
	for i := 0; i < 10; i++ {
		input, err := New("java.io.FileInputStream", "jnigo.cb")
		if err != nil {
			t.Error(err.Error())
			return
		}
		err = input.CallVoid("close")
		if err != nil {
			t.Error(err.Error())
			return
		}
		fmt.Println("exec----100")
	}
	tx.Commit()
	fmt.Println("----------------------------------------->")
}
