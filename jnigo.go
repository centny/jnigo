package jnigo

/*
#include <stdio.h>
#include <stdlib.h>
#include <jni.h>
#include "jnigo.h"
#cgo darwin CFLAGS: -fprofile-arcs -ftest-coverage -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include/ -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include/darwin
#cgo darwin LDFLAGS: -fprofile-arcs -ftest-coverage -L/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/jre/lib/server -ljvm
*/
import "C"
import (
	"fmt"
	"reflect"
	"strings"
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

//

type Jsize C.jsize
type Jarray C.jarray

// type Jobject C.jobject
type Jval C.jval

type Jboolean C.jboolean
type Jbyte C.jbyte
type Jchar C.jchar
type Jshort C.jshort
type Jint C.jint
type Jlong C.jlong
type Jfloat C.jfloat
type Jdouble C.jdouble

//
type Char byte
type Byte byte
type EmptyObjAry string

//
type VMOption struct {
	OptionString string
}

var GVM Jvm

func Init(os ...string) int {
	GVM.Version = JNI_VERSION_1_6
	GVM.IgnoreUnrecognized = JNI_TRUE
	for _, o := range os {
		GVM.AddVMOption2(o)
	}
	return int(GVM.Init())
}
func Destory() int {
	return int(GVM.Destroy())
}

type Jvm struct {
	Version            Jint
	IgnoreUnrecognized Jboolean
	//
	options []VMOption
	//
	env *C.JNIEnv
	jvm *C.JavaVM
}

func (j *Jvm) covary(arg interface{}) (string, C.jobject, error) {
	if arg == nil {
		return "", nil, Err("arg is nil")
	}
	pval := reflect.ValueOf(arg)
	ptype := reflect.TypeOf(arg)
	if ptype.Kind() != reflect.Slice {
		return "", nil, Err("not slice for:%v", arg)
	}
	var _bval_ Byte
	var _cval_ Char
	var _oval_ Object
	var _aval_ EmptyObjAry
	var __oval__ *Object
	switch ptype.Elem() {
	case reflect.TypeOf(_aval_):
		if pval.Len() < 1 {
			return "", nil, Err("empty slice for EmptyObjAry")
		}
		ovals := arg.([]EmptyObjAry)
		cls := j.FindClass(strings.Replace(string(ovals[0]), ".", "/", -1))
		if cls == nil {
			return "", nil, Err("class not found:%s", ovals[0])
		}
		cvals := C.JNIGO_NewObjectArray(j.env, 0, cls.cls, nil)
		return "[L" + string(ovals[0]) + ";", C.jobject(cvals), nil
	case reflect.TypeOf(_oval_):
		if pval.Len() < 1 {
			return "", nil, Err("empty slice(using EmptyObjAry?)")
		}
		ovals := arg.([]Object)
		vlen := C.jsize(len(ovals))
		cvals := C.JNIGO_NewObjectArray(j.env, vlen, ovals[0].Cls.cls, nil)
		for i, b := range ovals {
			C.JNIGO_SetObjectArrayElement(j.env, cvals, C.jsize(i), b.jobj)
		}
		return "[L" + ovals[0].Cls.Name + ";", C.jobject(cvals), nil
	case reflect.TypeOf(__oval__):
		if pval.Len() < 1 {
			return "", nil, Err("empty slice(using EmptyObjAry?)")
		}
		ovals := arg.([]*Object)
		vlen := C.jsize(len(ovals))
		cvals := C.JNIGO_NewObjectArray(j.env, vlen, ovals[0].Cls.cls, nil)
		for i, b := range ovals {
			C.JNIGO_SetObjectArrayElement(j.env, cvals, C.jsize(i), b.jobj)
		}
		return "[L" + ovals[0].Cls.Name + ";", C.jobject(cvals), nil
	case reflect.TypeOf(_bval_):
		vals := []C.jbyte{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, C.jbyte(pval.Index(i).Interface().(Byte)))
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewByteArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetByteArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[B", C.jobject(cvals), nil
	case reflect.TypeOf(_cval_):
		vals := []C.jchar{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, C.jchar(pval.Index(i).Interface().(Char)))
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewCharArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetCharArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[C", C.jobject(cvals), nil
	}
	//
	switch ptype.Elem().Kind() {
	case reflect.Bool:
		vals := []C.jboolean{}
		for _, b := range arg.([]bool) {
			if b {
				vals = append(vals, C.jboolean(JNI_TRUE))
			} else {
				vals = append(vals, C.jboolean(JNI_FALSE))
			}
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewBooleanArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetBooleanArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[Z", C.jobject(cvals), nil
	case reflect.Uint8:
		vals := []C.jbyte{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, C.jbyte(pval.Index(i).Interface().(uint8)))
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewByteArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetByteArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[B", C.jobject(cvals), nil
	case reflect.Int16:
		vals := []C.jshort{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, C.jshort(pval.Index(i).Interface().(int16)))
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewShortArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetShortArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[S", C.jobject(cvals), nil
	case reflect.Int32:
		vals := []C.jint{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, C.jint(pval.Index(i).Interface().(int32)))
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewIntArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetIntArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[I", C.jobject(cvals), nil
	case reflect.Int:
		vals := []C.jint{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, C.jint(pval.Index(i).Interface().(int)))
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewIntArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetIntArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[I", C.jobject(cvals), nil
	case reflect.Int64:
		vals := []C.jlong{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, C.jlong(pval.Index(i).Interface().(int64)))
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewLongArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetLongArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[J", C.jobject(cvals), nil
	case reflect.Float32:
		vals := []C.jfloat{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, C.jfloat(pval.Index(i).Interface().(float32)))
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewFloatArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetFloatArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[F", C.jobject(cvals), nil
	case reflect.Float64:
		vals := []C.jdouble{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, C.jdouble(pval.Index(i).Interface().(float64)))
		}
		vlen := C.jsize(len(vals))
		cvals := C.JNIGO_NewDoubleArray(j.env, vlen)
		if vlen > 0 {
			C.JNIGO_SetDoubleArrayRegion(j.env, cvals, 0, vlen, &vals[0])
		}
		return "[D", C.jobject(cvals), nil
	default:
		return "", nil, Err("invalid type:%s", ptype.Elem().Kind().String())
	}
}

func (j *Jvm) covval(arg interface{}) (string, C.jval, error) {
	var _bval_ Byte
	var _cval_ Char
	var _oval_ Object
	var __oval__ *Object
	ptype := reflect.TypeOf(arg)
	switch ptype {
	case reflect.TypeOf(_bval_):
		return "B", C.jval{b: C.jbyte(arg.(Byte)), typ: 1}, nil
	case reflect.TypeOf(_cval_):
		return "C", C.jval{c: C.jchar(arg.(Char)), typ: 2}, nil
	case reflect.TypeOf(_oval_):
		return "L" + arg.(Object).Cls.Name + ";", C.jval{l: arg.(Object).jobj, typ: 8}, nil
	case reflect.TypeOf(__oval__):
		return "L" + arg.(*Object).Cls.Name + ";", C.jval{l: arg.(*Object).jobj, typ: 8}, nil
	}
	switch ptype.Kind() {
	case reflect.Bool:
		if arg.(bool) {
			return "Z", C.jval{z: C.jboolean(JNI_TRUE), typ: 0}, nil
		} else {
			return "Z", C.jval{z: C.jboolean(JNI_FALSE), typ: 0}, nil
		}
	case reflect.Uint8:
		return "B", C.jval{b: C.jbyte(arg.(uint8)), typ: 1}, nil
	case reflect.Int16:
		return "S", C.jval{s: C.jshort(arg.(int16)), typ: 3}, nil
	case reflect.Int32:
		return "I", C.jval{i: C.jint(arg.(int32)), typ: 4}, nil
	case reflect.Int:
		return "I", C.jval{i: C.jint(arg.(int)), typ: 4}, nil
	case reflect.Int64:
		return "J", C.jval{j: C.jlong(arg.(int64)), typ: 5}, nil
	case reflect.Float32:
		return "F", C.jval{f: C.jfloat(arg.(float32)), typ: 6}, nil
	case reflect.Float64:
		return "D", C.jval{d: C.jdouble(arg.(float64)), typ: 7}, nil
	case reflect.Slice:
		sig, val, err := j.covary(arg)
		return sig, C.jval{l: val, typ: 8}, err
	default:
		return "", C.jval{typ: -1}, Err("invalid type:%s", ptype.Kind().String())
	}
}

func (j *Jvm) CovArgs(args ...interface{}) (string, []Jval, error) {
	vals := []Jval{}
	sigs := ""
	for _, arg := range args {
		sig, val, err := j.covval(arg)
		if err != nil {
			return "", nil, err
		} else {
			sigs += sig
			vals = append(vals, Jval(val))
		}
	}
	return sigs, vals, nil
}

func (j *Jvm) AddVMOption(o VMOption) {
	j.options = append(j.options, o)
}
func (j *Jvm) AddVMOption2(o string) {
	j.AddVMOption(VMOption{
		OptionString: o,
	})
}
func (j *Jvm) Show() {
	fmt.Println("Version:", j.Version)
	fmt.Println("IgnoreUnrecognized:", j.IgnoreUnrecognized)
	fmt.Println("Options:")
	for _, o := range j.options {
		fmt.Println("\t" + o.OptionString)
	}
}
func (j *Jvm) Init() Jint {
	options := []C.JavaVMOption{}
	for _, o := range j.options {
		os := C.CString(o.OptionString)
		defer C.free(unsafe.Pointer(os))
		options = append(options, C.JavaVMOption{optionString: os})
	}
	//
	vm_args := C.JavaVMInitArgs{}
	vm_args.version = JNI_VERSION_1_6
	vm_args.nOptions = (C.jint)(len(options))
	if vm_args.nOptions > 0 {
		vm_args.options = &options[0]
	}
	vm_args.ignoreUnrecognized = JNI_TRUE
	return (Jint)(C.JNI_CreateJavaVM(&j.jvm,
		(*unsafe.Pointer)(unsafe.Pointer(&j.env)),
		unsafe.Pointer(&vm_args)))
}
func (j *Jvm) Destroy() Jint {
	return (Jint)(C.JNIGO_DestroyJavaVM(j.jvm))
}
func (j *Jvm) FindClass(name string) *Class {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cls := C.JNIGO_FindClass(j.env, cname)
	if cls == nil {
		return nil
	} else {
		return &Class{
			Vm:   j,
			Name: name,
			cls:  cls,
		}
	}
}

type Class struct {
	Vm   *Jvm
	Name string
	//
	cls C.jclass
}

func (c *Class) GetMethod(name, arg_sig, ret_sig string, static bool) *Method {
	cname, csig := C.CString(name), C.CString(fmt.Sprintf("(%s)%s", arg_sig, ret_sig))
	defer C.free(unsafe.Pointer(cname))
	defer C.free(unsafe.Pointer(csig))
	var mid C.jmethodID
	if static {
		mid = C.JNIGO_GetStaticMethodID(c.Vm.env, c.cls, cname, csig)
	} else {
		mid = C.JNIGO_GetMethodID(c.Vm.env, c.cls, cname, csig)
	}
	if mid == nil {
		return nil
	} else {
		return &Method{
			Vm:     c.Vm,
			Cls:    c,
			Obj:    nil,
			Name:   name,
			ArgSig: arg_sig,
			RetSig: ret_sig,
			mid:    mid,
		}
	}
}
func (c *Class) New(args ...interface{}) (*Object, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return nil, err
	}
	m := c.GetMethod("<init>", sig, "V", false)
	if m == nil {
		return nil, Err("constructor not found by sig:(%s)V", sig)
	}
	return m.newObjectA(vals)
}

////////////////
func (c *Class) CallObject(name, ret_sig string, args ...interface{}) (*Object, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return nil, err
	}
	m := c.GetMethod(name, sig, ret_sig, true)
	if m == nil {
		return nil, Err("method not found by sig:(%s)%s", sig, ret_sig)
	}
	return m.CallObjectMethodA(vals)
}
func (c *Class) CallVoid(name string, args ...interface{}) error {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return err
	}
	m := c.GetMethod(name, sig, "V", true)
	if m == nil {
		return Err("method not found by sig:(%s)%s", sig, "V")
	}
	m.CallVoidMethodA(vals)
	return nil
}
func (c *Class) CallBoolean(name string, args ...interface{}) (bool, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return false, err
	}
	m := c.GetMethod(name, sig, "Z", true)
	if m == nil {
		return false, Err("method not found by sig:(%s)%s", sig, "Z")
	}
	return m.CallBooleanMethodA(vals), nil
}
func (c *Class) CallByte(name string, args ...interface{}) (byte, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := c.GetMethod(name, sig, "B", true)
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "B")
	}
	return m.CallByteMethodA(vals), nil
}
func (c *Class) CallChar(name string, args ...interface{}) (byte, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := c.GetMethod(name, sig, "C", true)
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "C")
	}
	return m.CallCharMethodA(vals), nil
}
func (c *Class) CallShort(name string, args ...interface{}) (int16, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := c.GetMethod(name, sig, "S", true)
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "S")
	}
	return m.CallShortMethodA(vals), nil
}
func (c *Class) CallInt(name string, args ...interface{}) (int, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := c.GetMethod(name, sig, "I", true)
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "I")
	}
	return m.CallIntMethodA(vals), nil
}
func (c *Class) CallLong(name string, args ...interface{}) (int64, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := c.GetMethod(name, sig, "J", true)
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "J")
	}
	return m.CallLongMethodA(vals), nil
}
func (c *Class) CallFloat(name string, args ...interface{}) (float32, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := c.GetMethod(name, sig, "F", true)
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "F")
	}
	return m.CallFloatMethodA(vals), nil
}
func (c *Class) CallDouble(name string, args ...interface{}) (float64, error) {
	sig, vals, err := c.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := c.GetMethod(name, sig, "D", true)
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "D")
	}
	return m.CallDoubleMethodA(vals), nil
}

///////////////

type Method struct {
	Vm     *Jvm
	Cls    *Class
	Obj    *Object
	Name   string
	ArgSig string
	RetSig string
	//
	mid C.jmethodID
}

func (m *Method) tobj() C.jobject {
	if m.Obj == nil {
		return C.jobject(m.Cls.cls)
	} else {
		return m.Obj.jobj
	}
}
func (m *Method) FindRetClass() *Class {
	return m.Vm.FindClass(strings.Trim(m.RetSig, "[L;"))
}
func (m *Method) call_args(vals []Jval) (*C.JNIEnv, C.jobject, C.jmethodID, *C.jval, C.int) {
	rvals := []C.jval{}
	for _, val := range vals {
		rvals = append(rvals, C.jval(val))
	}
	var tval *C.jval = nil
	if len(rvals) > 0 {
		tval = &rvals[0]
	}
	return m.Vm.env, m.tobj(), m.mid, tval, C.int(len(rvals))
}
func (m *Method) newObjectA(vals []Jval) (*Object, error) {
	res := C.JNIGO_NewObjectA(m.call_args(vals))
	return &Object{
		Vm:   m.Vm,
		Cls:  m.Cls,
		jobj: res,
	}, nil
}
func (m *Method) CallObjectMethodA(vals []Jval) (*Object, error) {
	var cls *Class = nil
	if strings.HasPrefix(m.RetSig, "[L") {
		cls = m.FindRetClass()
		if cls == nil {
			return nil, Err("invalid return class:%s", m.RetSig)
		}
	}
	res := C.JNIGO_CallObjectMethodA(m.call_args(vals))
	return &Object{
		Vm:   m.Vm,
		Cls:  cls,
		jobj: res,
	}, nil
}
func (m *Method) CallVoidMethodA(vals []Jval) {
	C.JNIGO_CallVoidMethodA(m.call_args(vals))
}
func (m *Method) CallBooleanMethodA(vals []Jval) bool {
	return C.JNIGO_CallBooleanMethodA(m.call_args(vals)) == JNI_TRUE
}
func (m *Method) CallByteMethodA(vals []Jval) byte {
	return byte(C.JNIGO_CallByteMethodA(m.call_args(vals)))
}
func (m *Method) CallCharMethodA(vals []Jval) byte {
	return byte(C.JNIGO_CallCharMethodA(m.call_args(vals)))
}
func (m *Method) CallShortMethodA(vals []Jval) int16 {
	return int16(C.JNIGO_CallShortMethodA(m.call_args(vals)))
}
func (m *Method) CallIntMethodA(vals []Jval) int {
	return int(C.JNIGO_CallIntMethodA(m.call_args(vals)))
}
func (m *Method) CallLongMethodA(vals []Jval) int64 {
	return int64(C.JNIGO_CallLongMethodA(m.call_args(vals)))
}
func (m *Method) CallFloatMethodA(vals []Jval) float32 {
	return float32(C.JNIGO_CallFloatMethodA(m.call_args(vals)))
}
func (m *Method) CallDoubleMethodA(vals []Jval) float64 {
	return float64(C.JNIGO_CallDoubleMethodA(m.call_args(vals)))
}

type Object struct {
	Vm   *Jvm
	Cls  *Class
	jobj C.jobject
}

func (o *Object) GetMethod(name, arg_sig, ret_sig string) *Method {
	m := o.Cls.GetMethod(name, arg_sig, ret_sig, false)
	if m != nil {
		m.Obj = o
	}
	return m
}

////////////////
func (o *Object) CallObject(name, ret_sig string, args ...interface{}) (*Object, error) {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return nil, err
	}
	m := o.GetMethod(name, sig, ret_sig)
	if m == nil {
		return nil, Err("method not found by sig:(%s)%s", sig, ret_sig)
	}
	return m.CallObjectMethodA(vals)
}
func (o *Object) CallVoid(name string, args ...interface{}) error {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return err
	}
	m := o.GetMethod(name, sig, "V")
	if m == nil {
		return Err("method not found by sig:(%s)%s", sig, "V")
	}
	m.CallVoidMethodA(vals)
	return nil
}
func (o *Object) CallBoolean(name string, args ...interface{}) (bool, error) {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return false, err
	}
	m := o.GetMethod(name, sig, "Z")
	if m == nil {
		return false, Err("method not found by sig:(%s)%s", sig, "Z")
	}
	return m.CallBooleanMethodA(vals), nil
}
func (o *Object) CallByte(name string, args ...interface{}) (byte, error) {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := o.GetMethod(name, sig, "B")
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "B")
	}
	return m.CallByteMethodA(vals), nil
}
func (o *Object) CallChar(name string, args ...interface{}) (byte, error) {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := o.GetMethod(name, sig, "C")
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "C")
	}
	return m.CallCharMethodA(vals), nil
}
func (o *Object) CallShort(name string, args ...interface{}) (int16, error) {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := o.GetMethod(name, sig, "S")
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "S")
	}
	return m.CallShortMethodA(vals), nil
}
func (o *Object) CallInt(name string, args ...interface{}) (int, error) {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := o.GetMethod(name, sig, "I")
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "I")
	}
	return m.CallIntMethodA(vals), nil
}
func (o *Object) CallLong(name string, args ...interface{}) (int64, error) {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := o.GetMethod(name, sig, "J")
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "J")
	}
	return m.CallLongMethodA(vals), nil
}
func (o *Object) CallFloat(name string, args ...interface{}) (float32, error) {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := o.GetMethod(name, sig, "F")
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "F")
	}
	return m.CallFloatMethodA(vals), nil
}
func (o *Object) CallDouble(name string, args ...interface{}) (float64, error) {
	sig, vals, err := o.Vm.CovArgs(args...)
	if err != nil {
		return 0, err
	}
	m := o.GetMethod(name, sig, "D")
	if m == nil {
		return 0, Err("method not found by sig:(%s)%s", sig, "D")
	}
	return m.CallDoubleMethodA(vals), nil
}

///////////////
func (o *Object) Len() int {
	return int(C.JNIGO_GetArrayLength(o.Vm.env, C.jarray(o.jobj)))
}
func (o *Object) GetObject(idx int) *Object {
	res := C.JNIGO_GetObjectArrayElement(o.Vm.env, C.jobjectArray(o.jobj), C.jsize(idx))
	if res == nil {
		return nil
	} else {
		return &Object{
			Vm:   o.Vm,
			Cls:  o.Cls,
			jobj: res,
		}
	}
}
func (o *Object) AsBoolAry(f func(o *Object, i int, v bool)) []bool {
	vlen := o.Len()
	if vlen < 1 {
		return []bool{}
	}
	lvs := make([]C.jboolean, vlen)
	C.JNIGO_GetBooleanArrayRegion(o.Vm.env, C.jbooleanArray(o.jobj), 0, C.jsize(vlen), &lvs[0])
	gvs := make([]bool, vlen)
	for i, v := range lvs {
		gvs[i] = v == JNI_TRUE
		if f != nil {
			f(o, i, gvs[i])
		}
	}
	return gvs
}
func (o *Object) AsByteAry(f func(o *Object, i int, v byte)) []byte {
	vlen := o.Len()
	if vlen < 1 {
		return []byte{}
	}
	lvs := make([]C.jbyte, vlen)
	C.JNIGO_GetByteArrayRegion(o.Vm.env, C.jbyteArray(o.jobj), 0, C.jsize(vlen), &lvs[0])
	gvs := make([]byte, vlen)
	for i, v := range lvs {
		gvs[i] = byte(v)
		if f != nil {
			f(o, i, gvs[i])
		}
	}
	return gvs
}

func (o *Object) AsCharAry(f func(o *Object, i int, v byte)) []byte {
	vlen := o.Len()
	if vlen < 1 {
		return []byte{}
	}
	lvs := make([]C.jchar, vlen)
	C.JNIGO_GetCharArrayRegion(o.Vm.env, C.jcharArray(o.jobj), 0, C.jsize(vlen), &lvs[0])
	gvs := make([]byte, vlen)
	for i, v := range lvs {
		gvs[i] = byte(v)
		if f != nil {
			f(o, i, gvs[i])
		}
	}
	return gvs
}
func (o *Object) AsShortAry(f func(o *Object, i int, v int16)) []int16 {
	vlen := o.Len()
	if vlen < 1 {
		return []int16{}
	}
	lvs := make([]C.jshort, vlen)
	C.JNIGO_GetShortArrayRegion(o.Vm.env, C.jshortArray(o.jobj), 0, C.jsize(vlen), &lvs[0])
	gvs := make([]int16, vlen)
	for i, v := range lvs {
		gvs[i] = int16(v)
		if f != nil {
			f(o, i, gvs[i])
		}
	}
	return gvs
}
func (o *Object) AsIntAry(f func(o *Object, i int, v int)) []int {
	vlen := o.Len()
	if vlen < 1 {
		return []int{}
	}
	lvs := make([]C.jint, vlen)
	C.JNIGO_GetIntArrayRegion(o.Vm.env, C.jintArray(o.jobj), 0, C.jsize(vlen), &lvs[0])
	gvs := make([]int, vlen)
	for i, v := range lvs {
		gvs[i] = int(v)
		if f != nil {
			f(o, i, gvs[i])
		}
	}
	return gvs
}
func (o *Object) AsLongAry(f func(o *Object, i int, v int64)) []int64 {
	vlen := o.Len()
	if vlen < 1 {
		return []int64{}
	}
	lvs := make([]C.jlong, vlen)
	C.JNIGO_GetLongArrayRegion(o.Vm.env, C.jlongArray(o.jobj), 0, C.jsize(vlen), &lvs[0])
	gvs := make([]int64, vlen)
	for i, v := range lvs {
		gvs[i] = int64(v)
		if f != nil {
			f(o, i, gvs[i])
		}
	}
	return gvs
}
func (o *Object) AsFloatAry(f func(o *Object, i int, v float32)) []float32 {
	vlen := o.Len()
	if vlen < 1 {
		return []float32{}
	}
	lvs := make([]C.jfloat, vlen)
	C.JNIGO_GetFloatArrayRegion(o.Vm.env, C.jfloatArray(o.jobj), 0, C.jsize(vlen), &lvs[0])
	gvs := make([]float32, vlen)
	for i, v := range lvs {
		gvs[i] = float32(v)
		if f != nil {
			f(o, i, gvs[i])
		}
	}
	return gvs
}
func (o *Object) AsDoubleAry(f func(o *Object, i int, v float64)) []float64 {
	vlen := o.Len()
	if vlen < 1 {
		return []float64{}
	}
	lvs := make([]C.jdouble, vlen)
	C.JNIGO_GetDoubleArrayRegion(o.Vm.env, C.jdoubleArray(o.jobj), 0, C.jsize(vlen), &lvs[0])
	gvs := make([]float64, vlen)
	for i, v := range lvs {
		gvs[i] = float64(v)
		if f != nil {
			f(o, i, gvs[i])
		}
	}
	return gvs
}
