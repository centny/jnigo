package jnigo

import (
	"fmt"
	"github.com/Centny/Cny4go/util"
	"os"
	"testing"
)

// func TestChkSig(t *testing.T) {
// 	var a byte = 0
// 	cov_arg(1, "", Object{}, a, []int{})
// 	fmt.Println("end...")
// }
func init() {
	Init("-Djava.class.path=java/bin")
}
func initjava() {
	os.RemoveAll("java/bin")
	os.Mkdir("java/bin", os.ModePerm)
	fmt.Println(util.Exec("javac", "-d", "java/bin", "java/src/jnigo/*"))
}
func args_t(t *testing.T, vm *Jvm, tsig string, args ...interface{}) {
	sig, _, err := vm.CovArgs(args...)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if sig != tsig {
		t.Error(sig + " is not " + tsig)
	}
}

func TestCovArgs(t *testing.T) {
	vm := &GVM
	vm.Show()
	args_t(t, vm, "II", 1, 1)
	args_t(t, vm, "IZ", 1, true)
	args_t(t, vm, "IZ", 1, false)
	args_t(t, vm, "ZBCSIJFD", true, Byte(1), Char(1), int16(1), 1, int64(1), float32(1), float64(1))
	args_t(t, vm, "ZBCSIJFD", false, uint8(1), Char(1), int16(1), int32(1), int64(1), float32(1), float64(1))
	args_t(t, vm, "[Z[B[C[S[I[J[F[D", []bool{true}, []Byte{1}, []Char{1}, []int16{1}, []int{1}, []int64{1}, []float32{1}, []float64{1})
	args_t(t, vm, "[Z[B[C[S[I[J[F[D", []bool{false}, []uint8{1}, []Char{1}, []int16{1}, []int32{1}, []int64{1}, []float32{1}, []float64{1})
	args_t(t, vm, "[Z[B[C[S[I[J[F[D", []bool{}, []Byte{}, []Char{}, []int16{}, []int{}, []int64{}, []float32{}, []float64{})
	args_t(t, vm, "[Z[B[C[S[I[J[F[D", []bool{}, []uint8{}, []Char{}, []int16{}, []int32{}, []int64{}, []float32{}, []float64{})
	args_t(t, vm, "[Ljava.lang.String;", []EmptyObjAry{"java.lang.String"})
	//
	fmt.Println(vm.covary(1))
	fmt.Println(vm.covary(nil))
	fmt.Println(vm.covary([]EmptyObjAry{}))
	fmt.Println(vm.covary([]EmptyObjAry{"asss"}))
	fmt.Println(vm.covary([]Object{}))
	fmt.Println(vm.covary([]string{""}))
	//
	fmt.Println(vm.covval("arg"))
	fmt.Println(vm.CovArgs("arg"))
	//
	// _, err = NewJvm("-jjsjfs")
	// if err == nil {
	// 	t.Error("not err")
	// 	return
	// }
	// fmt.Println(err.Error())
}

func TestStaticM(t *testing.T) {
	vm := &GVM
	cls := vm.FindClass("jnigo/StaticM")
	if cls == nil {
		t.Error("class not found")
		return
	}
	cls.CallVoid("showz", false)
	cls.CallVoid("showb", uint8(1))
	cls.CallVoid("showb", Byte(1))
	cls.CallVoid("showc", Char(1))
	cls.CallVoid("shows", int16(1))
	cls.CallVoid("showi", 1)
	cls.CallVoid("showj", int64(1))
	cls.CallVoid("showf", float32(1))
	cls.CallVoid("showd", float64(1))
	//
	fmt.Println(cls.CallVoid("getv"))
	fmt.Println(cls.CallBoolean("getz"))
	fmt.Println(cls.CallByte("getb"))
	fmt.Println(cls.CallChar("getc"))
	fmt.Println(cls.CallShort("gets"))
	fmt.Println(cls.CallInt("geti"))
	fmt.Println(cls.CallLong("getj"))
	fmt.Println(cls.CallFloat("getf"))
	fmt.Println(cls.CallDouble("getd"))
	//
	fmt.Println(cls.CallVoid("get"))
	fmt.Println(cls.CallBoolean("get"))
	fmt.Println(cls.CallByte("get"))
	fmt.Println(cls.CallChar("get"))
	fmt.Println(cls.CallShort("get"))
	fmt.Println(cls.CallInt("get"))
	fmt.Println(cls.CallLong("get"))
	fmt.Println(cls.CallFloat("get"))
	fmt.Println(cls.CallDouble("get"))
	//
	fmt.Println(cls.CallVoid("getv", ""))
	fmt.Println(cls.CallBoolean("getz", ""))
	fmt.Println(cls.CallByte("getb", ""))
	fmt.Println(cls.CallChar("getc", ""))
	fmt.Println(cls.CallShort("gets", ""))
	fmt.Println(cls.CallInt("geti", ""))
	fmt.Println(cls.CallLong("getj", ""))
	fmt.Println(cls.CallFloat("getf", ""))
	fmt.Println(cls.CallDouble("getd", ""))
}

func TestObjectM(t *testing.T) {
	vm := &GVM
	cls := vm.FindClass("jnigo/ObjectM")
	if cls == nil {
		t.Error("class not found")
		return
	}
	fmt.Println(cls.New())
	fmt.Println(cls.New("..."))
	obj, err := cls.New(1, 2)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(obj.CallInt("getA"))
	fmt.Println(obj.CallInt("getB"))
	obj.CallVoid("setA", 211)
	obj.CallVoid("setB", 222)
	fmt.Println(obj.CallInt("getA"))
	fmt.Println(obj.CallInt("getB"))
	//
	//
	obj.CallVoid("showz", false)
	obj.CallVoid("showb", uint8(1))
	obj.CallVoid("showb", Byte(1))
	obj.CallVoid("showc", Char(1))
	obj.CallVoid("shows", int16(1))
	obj.CallVoid("showi", 1)
	obj.CallVoid("showj", int64(1))
	obj.CallVoid("showf", float32(1))
	obj.CallVoid("showd", float64(1))
	//
	fmt.Println(obj.CallVoid("getv"))
	fmt.Println(obj.CallBoolean("getz"))
	fmt.Println(obj.CallByte("getb"))
	fmt.Println(obj.CallChar("getc"))
	fmt.Println(obj.CallShort("gets"))
	fmt.Println(obj.CallInt("geti"))
	fmt.Println(obj.CallLong("getj"))
	fmt.Println(obj.CallFloat("getf"))
	fmt.Println(obj.CallDouble("getd"))
	//
	fmt.Println(obj.CallVoid("get"))
	fmt.Println(obj.CallBoolean("get"))
	fmt.Println(obj.CallByte("get"))
	fmt.Println(obj.CallChar("get"))
	fmt.Println(obj.CallShort("get"))
	fmt.Println(obj.CallInt("get"))
	fmt.Println(obj.CallLong("get"))
	fmt.Println(obj.CallFloat("get"))
	fmt.Println(obj.CallDouble("get"))
	//
	fmt.Println(obj.CallVoid("getv", ""))
	fmt.Println(obj.CallBoolean("getz", ""))
	fmt.Println(obj.CallByte("getb", ""))
	fmt.Println(obj.CallChar("getc", ""))
	fmt.Println(obj.CallShort("gets", ""))
	fmt.Println(obj.CallInt("geti", ""))
	fmt.Println(obj.CallLong("getj", ""))
	fmt.Println(obj.CallFloat("getf", ""))
	fmt.Println(obj.CallDouble("getd", ""))
}
func TestArray(t *testing.T) {
	vm := &GVM
	clsa := vm.FindClass("jnigo/A")
	if clsa == nil {
		t.Error("class not found")
		return
	}
	clsary := vm.FindClass("jnigo/Ary")
	if clsary == nil {
		t.Error("class not found")
		return
	}
	obja, _ := clsa.New()
	objb, _ := clsa.New()
	fmt.Println(clsary.CallVoid("showas", []Object{*obja, *objb}))
	// vm.NewObjectArray()
	objary, _ := clsary.New()
	// fmt.Println(vm.CovArgs([]Object{obja, objb}))
	fmt.Println(objary.CallVoid("show", []Object{*obja, *objb}))
}
func TestAbc(t *testing.T) {
	vm := &GVM
	clsa := vm.FindClass("jnigo/A")
	if clsa == nil {
		t.Error("class not found")
		return
	}
	clsb := vm.FindClass("jnigo/B")
	if clsb == nil {
		t.Error("class not found")
		return
	}
	obja, err := clsa.New()
	if err != nil {
		t.Error(err.Error())
		return
	}
	obja.CallVoid("show")
	//
	objb, err := clsb.New(obja)
	if err != nil {
		t.Error(err.Error())
		return
	}
	objb.CallVoid("show")

	obja_0, err := clsa.New()
	if err != nil {
		t.Error(err.Error())
		return
	}
	obja_1, err := clsa.New()
	if err != nil {
		t.Error(err.Error())
		return
	}
	clsc := vm.FindClass("jnigo/C")
	if clsc == nil {
		t.Error("class not found")
		return
	}
	objc, err := clsc.New()
	if err != nil {
		t.Error(err.Error())
		return
	}
	objc.CallVoid("setAs", []Object{*obja_0, *obja_1})
	fmt.Println(objc.CallObject("getAs", "[Ljnigo/A;"))
	objc.CallVoid("showas")
}

func TestTary(t *testing.T) {
	vm := &GVM
	clsta := vm.FindClass("jnigo/Tary")
	if clsta == nil {
		t.Error("class not found")
		return
	}
	objta, err := clsta.New()
	if err != nil {
		t.Error(err.Error())
		return
	}
	zs, _ := objta.CallObject("zs", "[Z")
	zs.AsBoolAry(func(o *Object, i int, v bool) {
		fmt.Println(v)
	})
	bs, _ := objta.CallObject("bs", "[B")
	bs.AsByteAry(func(o *Object, i int, v byte) {
		fmt.Println(v)
	})
	cs, _ := objta.CallObject("cs", "[C")
	cs.AsCharAry(func(o *Object, i int, v byte) {
		fmt.Println(v)
	})
	ss, _ := objta.CallObject("ss", "[S")
	ss.AsShortAry(func(o *Object, i int, v int16) {
		fmt.Println(v)
	})
	// fmt.Println(ss.AsShortAry(nil))
	is, _ := objta.CallObject("is", "[I")
	is.AsIntAry(func(o *Object, i int, v int) {
		fmt.Println(v)
	})
	js, _ := objta.CallObject("js", "[J")
	js.AsLongAry(func(o *Object, i int, v int64) {
		fmt.Println(v)
	})
	fs, _ := objta.CallObject("fs", "[F")
	fs.AsFloatAry(func(o *Object, i int, v float32) {
		fmt.Println(v)
	})
	ds, _ := objta.CallObject("ds", "[D")
	ds.AsDoubleAry(func(o *Object, i int, v float64) {
		fmt.Println(v)
	})
	ls, _ := objta.CallObject("ls", "[Ljnigo/A;")
	for i := 0; i < ls.Len(); i++ {
		oo := ls.GetObject(i)
		oo.CallVoid("show")
	}
	es, _ := objta.CallObject("es", "[I")
	// fmt.Println(es, err)
	es.AsBoolAry(nil)
	es.AsByteAry(nil)
	es.AsCharAry(nil)
	es.AsShortAry(nil)
	es.AsIntAry(nil)
	es.AsLongAry(nil)
	es.AsFloatAry(nil)
	es.AsDoubleAry(nil)
	fmt.Println("----->")
}

func TestField(t *testing.T) {
	vm := &GVM
	clsf := vm.FindClass("jnigo/Field")
	if clsf == nil {
		t.Error("class not found")
		return
	}
	fmt.Println(clsf.Set("zz", false))
	fmt.Println(clsf.Set("jz", int64(11)))
	fmt.Println(clsf.Boolean("zz"))
	fmt.Println(clsf.Byte("bz"))
	fmt.Println(clsf.Char("cz"))
	fmt.Println(clsf.Short("sz"))
	fmt.Println(clsf.Int("iz"))
	fmt.Println(clsf.Long("jz"))
	fmt.Println(clsf.Float("fz"))
	fmt.Println(clsf.Double("dz"))
	//
	// fmt.Println(clsf.GetField("zz", "Z", false))
	objf, err := clsf.New()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(objf.Set("zz_p", false))
	fmt.Println(objf.Set("jz_p", int64(11)))
	fmt.Println(objf.Boolean("zz_p"))
	fmt.Println(objf.Byte("bz_p"))
	fmt.Println(objf.Char("cz_p"))
	fmt.Println(objf.Short("sz_p"))
	fmt.Println(objf.Int("iz_p"))
	fmt.Println(objf.Long("jz_p"))
	fmt.Println(objf.Float("fz_p"))
	fmt.Println(objf.Double("dz_p"))
	//
	at, _ := clsf.Object("as", "Ljnigo/A;")
	if at.IsNull() {
		t.Error("is null")
		return
	}
	at.CallVoid("show")
	//
	clsa := vm.FindClass("jnigo/A")
	if clsa == nil {
		t.Error("class not found")
		return
	}
	av, _ := clsa.New()
	fmt.Println(clsf.Set("as", av))
	a, err := clsf.Object("as", "Ljnigo/A;")
	if err != nil {
		t.Error(err.Error())
		return
	}
	if a.IsNull() {
		t.Error("null value")
		return
	}
	// fmt.Println(a, err)
	a.CallVoid("show")
	//
	zss, err := clsf.Object("zzs", "[I")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(zss.AsIntAry(nil))
	//
	ass, err := clsf.Object("ass", "[Ljnigo/A;")
	if err != nil {
		t.Error(err.Error())
		return
	}
	as := ass.GetObject(0)
	as.CallVoid("show")
	//
	zss, err = objf.Object("zzs_p", "[I")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(zss.AsIntAry(nil))
	//
	ass, err = objf.Object("ass_p", "[Ljnigo/A;")
	if err != nil {
		t.Error(err.Error())
		return
	}
	as = ass.GetObject(0)
	as.CallVoid("show")
	//
}

func TestString(t *testing.T) {
	vm := &GVM
	// fmt.Println(vm.covary([]string{"vv", "va"}))
	fmt.Println(vm.New("java/lang/String", "..."))
	obj := vm.NewS("v")
	// fmt.Println(obj)
	fmt.Println(obj.ToString())
	// fmt.Println(obj.AsStringAry(nil))
	fmt.Println(obj.CallInt("length"))
	// fmt.Println(obj.CallObject("toCharArray", "[C").AsCharAry(nil))
	//
	ary := vm.NewAryS("v1", "v2", "v3")
	fmt.Println(ary.Len())
	// fmt.Println(ary)
	fmt.Println(ary.AsStringAry(nil))
	//
	clsa := vm.FindClass("jnigo/A")
	obja, _ := clsa.New()
	clss, _ := clsa.CallObject("tss", "Ljava/lang/String;")
	objs, _ := obja.CallObject("ts", "Ljava/lang/String;")
	fmt.Println(clss, objs)
	// fmt.Println(objs.CallObject("toCharArray", "[C"))
	fmt.Println(clss.AsString(), objs.AsString())
}

func TestAshow(t *testing.T) {
	vm := &GVM
	clsa := vm.FindClass("Ljnigo/A;")
	obja, _ := clsa.New()
	objv, _ := obja.As("Ljava/lang/Object;")
	objs, _ := vm.NewAry("Ljava/lang/Object;", 1)
	objs.SetObject(0, objv)
	fmt.Println(obja.CallVoid("show", objv))
	fmt.Println(obja.CallVoid("show", objs))
	show, err := obja.CallObject("show", "Ljava/lang/String;",
		true, Byte(1), Char(1), int16(1),
		1, int64(1), float32(1), float64(1),
		objv, "jjjjj", []bool{false}, []Byte{1, 2},
		[]Char{3, 4}, []int16{11, 12}, []int{21, 22},
		[]int64{31, 32}, []float32{41, 42}, []float64{51, 52},
		objs, []string{"aaa", "bbb"},
	)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(show.ToString())
}
