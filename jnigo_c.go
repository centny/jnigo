package jnigo

/*
#include <stdio.h>
#include <stdlib.h>
#include "jnigo_c.h"
#cgo darwin CPPFLAGS: -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include/darwin
#cgo darwin CFLAGS: -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include/darwin
#cgo darwin LDFLAGS: -L/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/jre/lib/server -ljvm
*/
import "C"
import (
	"errors"
	"unsafe"
)

const (
	JNI_TRUE  = C.JNI_TRUE
	JNI_FALSE = C.JNI_FALSE
	JNI_OK    = C.JNI_OK
)

const (
	JNI_VERSION_1_1 = C.JNI_VERSION_1_1
	JNI_VERSION_1_2 = C.JNI_VERSION_1_2
	JNI_VERSION_1_4 = C.JNI_VERSION_1_4
	JNI_VERSION_1_6 = C.JNI_VERSION_1_6
)

type jboolean C.jboolean
type jbyte C.jbyte
type jchar C.jchar
type jshort C.jshort
type jint C.jint
type jlong C.jlong
type jfloat C.jfloat
type jdouble C.jdouble

type Val struct {
	z   bool
	b   byte
	c   byte
	s   int16
	i   int
	j   int64
	f   float32
	d   float64
	l   unsafe.Pointer
	typ byte
}

func (v *Val) jval() C.jval {
	var zv C.jboolean
	if v.z {
		zv = C.jboolean(1)
	} else {
		zv = C.jboolean(0)
	}
	return C.jval{
		z:   zv,
		b:   C.jbyte(v.b),
		c:   C.jchar(v.c),
		s:   C.jshort(v.s),
		i:   C.jint(v.i),
		j:   C.jlong(v.j),
		f:   C.jfloat(v.f),
		d:   C.jdouble(v.d),
		l:   C.jobject(v.l),
		typ: C.char(v.typ),
	}
}

type Vmo C.c_vmo

func (v *Vmo) Err() error {
	if v.valid == 0 {
		cc := (*C.char)(unsafe.Pointer(&v.msg))
		return errors.New(C.GoString(cc))
	} else {
		return nil
	}
}

type Class C.c_class

func (c *Class) Err() error {
	v := Vmo(C.c_vmo(c.vmo))
	return v.Err()
}
func (c *Class) FindMethod(name, vsig, rsig string, static_ int) (Method, error) {
	return JNIGO_findClsMethod(*c, name, vsig, rsig, static_)
}
func (c *Class) FindField(name, rsig string, static_ int) (Field, error) {
	return JNIGO_findClsField(*c, name, rsig, static_)
}
func (c *Class) Call(name, vsig, rsig string, args []Val) (Object, error) {
	m, err := c.FindMethod(name, vsig, rsig, 1)
	if err != nil {
		return Object{}, err
	}
	return JNIGO_callA(m, args)
}
func (c *Class) New(args ...interface{}) (Object, error) {
	sig, vals, err := CovArgs(args...)
	if err != nil {
		return Object{}, err
	}
	m, err := c.FindMethod("<init>", sig, "V", 0)
	if err != nil {
		return Object{}, err
	}
	return JNIGO_newA(m, vals)
}

//
type Object C.c_object

func (o *Object) FindMethod(name, vsig, rsig string) (Method, error) {
	return JNIGO_findObjMethod(*o, name, vsig, rsig)
}
func (o *Object) FindField(name, rsig string) (Field, error) {
	return JNIGO_findObjField(*o, name, rsig)
}
func (o *Object) Err() error {
	v := Vmo(C.c_vmo(o.vmo))
	return v.Err()
}

func (o *Object) Val() Val {
	return Val{
		z:   o.val.z != 0,
		b:   byte(o.val.b),
		c:   byte(o.val.c),
		s:   int16(o.val.s),
		i:   int(o.val.i),
		j:   int64(o.val.j),
		f:   float32(o.val.f),
		d:   float64(o.val.d),
		l:   unsafe.Pointer(o.val.l),
		typ: byte(o.val.typ),
	}
}
func (o *Object) Sig() string {
	cc := (*C.char)(unsafe.Pointer(&o.sig))
	return C.GoString(cc)
}
func (o *Object) Call(name, vsig, rsig string, args []Val) (Object, error) {
	return JNIGO_callA_o(*o, name, vsig, rsig, args)
}
func (o *Object) CallV(name, rsig string, args ...interface{}) (Object, error) {
	sig, vals, err := CovArgs(args...)
	if err != nil {
		return Object{}, err
	}
	return o.Call(name, sig, rsig, vals)
}
func (o *Object) Z() bool {
	return o.val.z != 0
}
func (o *Object) B() byte {
	return byte(o.val.b)
}
func (o *Object) C() byte {
	return byte(o.val.c)
}
func (o *Object) S() int16 {
	return int16(o.val.s)
}
func (o *Object) I() int {
	return int(o.val.i)
}
func (o *Object) J() int64 {
	return int64(o.val.j)
}
func (o *Object) F() float32 {
	return float32(o.val.f)
}
func (o *Object) D() float64 {
	return float64(o.val.d)
}
func (o *Object) Len() int {
	return JNIGO_len(*o)
}
func (o *Object) Zs() []bool {
	l := o.Len()
	if l < 1 {
		return []bool{}
	}
	bys := make([]C.jboolean, l)
	JNIGO_cary(*o, unsafe.Pointer(&bys[0]), 0, l)
	tbs := make([]bool, l)
	for i, b := range bys {
		tbs[i] = b != 0
	}
	return tbs
}
func (o *Object) Bs() []byte {
	l := o.Len()
	if l < 1 {
		return []byte{}
	}
	bys := make([]C.jbyte, l)
	JNIGO_cary(*o, unsafe.Pointer(&bys[0]), 0, l)
	tbs := make([]byte, l)
	for i, b := range bys {
		tbs[i] = byte(b)
	}
	return tbs
}
func (o *Object) Cs() []byte {
	l := o.Len()
	if l < 1 {
		return []byte{}
	}
	bys := make([]C.jchar, l)
	JNIGO_cary(*o, unsafe.Pointer(&bys[0]), 0, l)
	tbs := make([]byte, l)
	for i, b := range bys {
		tbs[i] = byte(b)
	}
	return tbs
}
func (o *Object) Ss() []int16 {
	l := o.Len()
	if l < 1 {
		return []int16{}
	}
	bys := make([]C.jshort, l)
	JNIGO_cary(*o, unsafe.Pointer(&bys[0]), 0, l)
	tbs := make([]int16, l)
	for i, b := range bys {
		tbs[i] = int16(b)
	}
	return tbs
}
func (o *Object) Is() []int {
	l := o.Len()
	if l < 1 {
		return []int{}
	}
	bys := make([]C.jint, l)
	JNIGO_cary(*o, unsafe.Pointer(&bys[0]), 0, l)
	tbs := make([]int, l)
	for i, b := range bys {
		tbs[i] = int(b)
	}
	return tbs
}
func (o *Object) Js() []int64 {
	l := o.Len()
	if l < 1 {
		return []int64{}
	}
	bys := make([]C.jlong, l)
	JNIGO_cary(*o, unsafe.Pointer(&bys[0]), 0, l)
	tbs := make([]int64, l)
	for i, b := range bys {
		tbs[i] = int64(b)
	}
	return tbs
}
func (o *Object) Fs() []float32 {
	l := o.Len()
	if l < 1 {
		return []float32{}
	}
	bys := make([]C.jfloat, l)
	JNIGO_cary(*o, unsafe.Pointer(&bys[0]), 0, l)
	tbs := make([]float32, l)
	for i, b := range bys {
		tbs[i] = float32(b)
	}
	return tbs
}
func (o *Object) Ds() []float64 {
	l := o.Len()
	if l < 1 {
		return []float64{}
	}
	bys := make([]C.jdouble, l)
	JNIGO_cary(*o, unsafe.Pointer(&bys[0]), 0, l)
	tbs := make([]float64, l)
	for i, b := range bys {
		tbs[i] = float64(b)
	}
	return tbs
}
func (o *Object) Sary(ptr unsafe.Pointer, l int) {
	JNIGO_sary(*o, ptr, 0, l)
}
func (o *Object) Get(idx int) (Object, error) {
	return JNIGO_get_o(*o, idx)
}
func (o *Object) Set(idx int, v Val) {
	JNIGO_set_o(*o, idx, v)
}
func (o *Object) String() (string, error) {
	u := JNIGO_newS("UTF-8")
	b, err := o.Call("getBytes", "Ljava/lang/String;", "[B", []Val{
		u.Val(),
	})
	if err == nil {
		return string(b.Bs()), nil
	} else {
		return "", err
	}
}
func (o *Object) ToString() string {
	obj, _ := o.Call("toString", "", "Ljava/lang/String;", nil)
	str, _ := obj.String()
	return str
}

type Method C.c_method

func (m *Method) Err() error {
	v := Vmo(C.c_vmo(m.vmo))
	return v.Err()
}

type Field C.c_field

func (f *Field) Err() error {
	v := Vmo(C.c_vmo(f.vmo))
	return v.Err()
}

/*



*/
///////////
//
func JNIGO_setVersion(v int) {
	C.JNIGO_setVersion(C.int(v))
}

func JNIGO_setIgnoreUnrecognized(i int) {
	C.JNIGO_setIgnoreUnrecognized(C.int(i))
}

func JNIGO_addOption(option string) {
	cs := C.CString(option)
	C.JNIGO_addOption(cs)
	C.free(unsafe.Pointer(cs))
}

func JNIGO_init() int {
	return int(C.JNIGO_init())
}

func JNIGO_destory() {
	C.JNIGO_destory()
}

func JNIGO_errOccurred() error {
	v := Object(C.JNIGO_errOccurred())
	if v.Err() != nil {
		return nil
	}
	u := JNIGO_newS("UTF-8")
	v, _ = v.Call("toString", "", "Ljava/lang/String", nil)
	b, _ := v.Call("getBytes", "Ljava/lang/String;", "[B", []Val{
		u.Val(),
	})
	return errors.New(string(b.Bs()))
}

func JNIGO_errClear() {
	C.JNIGO_errClear()
}
func JNIGO_newAry(sig string, l int) (Object, error) {
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))
	obj := Object(C.JNIGO_newAry(csig, C.int(l)))
	return obj, obj.Err()
}

// //
func JNIGO_findClass(name string) (Class, error) {
	cs := C.CString(SigName(name))
	defer C.free(unsafe.Pointer(cs))
	v := Class(C.JNIGO_findClass(cs))
	return v, v.Err()
}

func JNIGO_findClsMethod(cls Class, name, vsig, rsig string, static_ int) (Method, error) {
	cname := C.CString(name)
	cvsig := C.CString(SigName(vsig))
	crsig := C.CString(SigName(rsig))
	defer func() {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(cvsig))
		C.free(unsafe.Pointer(crsig))
	}()
	v := Method(C.JNIGO_findClsMethod(C.c_class(cls), cname, cvsig, crsig, C.int(static_)))
	return v, v.Err()
}

func JNIGO_findClsField(cls Class, name, rsig string, static_ int) (Field, error) {
	cname := C.CString(name)
	crsig := C.CString(SigName(rsig))
	defer func() {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(crsig))
	}()
	v := Field(C.JNIGO_findClsField(C.c_class(cls), cname, crsig, C.int(static_)))
	return v, v.Err()
}

func JNIGO_findObjMethod(obj Object, name, vsig, rsig string) (Method, error) {
	cname := C.CString(name)
	cvsig := C.CString(SigName(vsig))
	crsig := C.CString(SigName(rsig))
	defer func() {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(cvsig))
		C.free(unsafe.Pointer(crsig))
	}()
	v := Method(C.JNIGO_findObjMethod(C.c_object(obj), cname, cvsig, crsig))
	return v, v.Err()
}
func JNIGO_findObjField(obj Object, name, rsig string) (Field, error) {
	cname := C.CString(name)
	crsig := C.CString(SigName(rsig))
	defer func() {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(crsig))
	}()
	v := Field(C.JNIGO_findObjField(C.c_object(obj), cname, crsig))
	return v, v.Err()
}
func JNIGO_len(obj Object) int {
	return int(C.JNIGO_len(C.c_object(obj)))
}
func JNIGO_cary(obj Object, buf unsafe.Pointer, idx, l int) {
	C.JNIGO_cary(C.c_object(obj), buf, C.int(idx), C.int(l))
}
func JNIGO_sary(obj Object, buf unsafe.Pointer, idx, l int) {
	C.JNIGO_sary(C.c_object(obj), buf, C.int(idx), C.int(l))
}
func JNIGO_newS(bys string) Object {
	if len(bys) > 0 {
		jbs := toByte([]byte(bys))
		return Object(C.JNIGO_newS(&jbs[0], C.int(len(bys))))
	} else {
		return Object(C.JNIGO_newS(nil, 0))
	}

}
func JNIGO_newA(m Method, args []Val) (Object, error) {
	if args != nil && len(args) > 0 {
		jvs := make([]C.jval, len(args))
		for i, v := range args {
			jvs[i] = v.jval()
		}
		v := Object(C.JNIGO_newA(C.c_method(m), &jvs[0], C.int(len(jvs))))
		return v, v.Err()
	} else {
		v := Object(C.JNIGO_newA(C.c_method(m), nil, 0))
		return v, v.Err()
	}
}

func JNIGO_callA(m Method, args []Val) (Object, error) {
	if args != nil && len(args) > 0 {
		jvs := make([]C.jval, len(args))
		for i, v := range args {
			jvs[i] = v.jval()
		}
		v := Object(C.JNIGO_callA(C.c_method(m), &jvs[0], C.int(len(jvs))))
		return v, v.Err()
	} else {
		v := Object(C.JNIGO_callA(C.c_method(m), nil, 0))
		return v, v.Err()
	}
}
func JNIGO_callA_o(obj Object, name, vsig, rsig string, args []Val) (Object, error) {
	m, err := JNIGO_findObjMethod(obj, name, vsig, rsig)
	if err != nil {
		return Object{}, err
	}
	return JNIGO_callA(m, args)
}
func JNIGO_get(f Field) (Object, error) {
	v := Object(C.JNIGO_get(C.c_field(f)))
	return v, v.Err()
}
func JNIGO_get_o(obj Object, idx int) (Object, error) {
	v := Object(C.JNIGO_get_o(C.c_object(obj), C.int(idx)))
	return v, v.Err()
}
func JNIGO_set_o(obj Object, idx int, v Val) {
	C.JNIGO_set_o(C.c_object(obj), C.int(idx), v.jval())
}
func JNIGO_set(f Field, arg Val) {
	C.JNIGO_set(C.c_field(f), arg.jval())
}

func toByte(bys []byte) []C.jbyte {
	jbs := make([]C.jbyte, len(bys))
	for i, b := range bys {
		jbs[i] = C.jbyte(b)
	}
	return jbs
}
