package jnigo

/*
#include <stdio.h>
#include <stdlib.h>
#include <jni.h>
#include "jnigo_c.h"
#cgo darwin CPPFLAGS: -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include/darwin
#cgo darwin CFLAGS: -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include -I/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/include/darwin
#cgo darwin LDFLAGS: -L/Library/Java/JavaVirtualMachines/jdk1.7.0_65.jdk/Contents/Home/jre/lib/server -ljvm
*/
import "C"
import (
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

// struct jval {
// 	jboolean ;
// 	jbyte b;
// 	jchar c;
// 	jshort s;
// 	jint i;
// 	jlong j;
// 	jfloat f;
// 	jdouble d;
// 	jobject l;
// 	char typ;
// } jval;
type c_vmo C.c_vmo
type c_err C.c_err
type c_class C.c_class
type c_object C.c_object
type c_method C.c_object
type c_field C.c_field

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

func JNIGO_errOccurred() c_err {
	return c_err(JNIGO_errOccurred())
}

func JNIGO_errClear() {
	C.JNIGO_errClear()
}

// //
func JNIGO_findClass(name string) c_class {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	return c_class(C.JNIGO_findClass(cs))
}

func JNIGO_findClsMethod(cls c_class, name, vsig, rsig string, static_ int) c_method {
	cname := C.CString(name)
	cvsig := C.CString(vsig)
	crsig := C.CString(rsig)
	defer func() {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(cvsig))
		C.free(unsafe.Pointer(crsig))
	}()
	return c_method(C.JNIGO_findClsMethod(C.c_class(cls), cname, cvsig, crsig, C.int(static_)))
}

func JNIGO_findClsField(cls c_class, name, rsig string, static_ int) c_field {
	cname := C.CString(name)
	crsig := C.CString(rsig)
	defer func() {
		C.free(unsafe.Pointer(cname))
		C.free(unsafe.Pointer(crsig))
	}()
	return c_field(C.JNIGO_findClsField(C.c_class(cls), cname, crsig, C.int(static_)))
}

// struct c_method JNIGO_findObjMethod(struct c_object obj, const char* name,
// 		const char* vsig, const char* rsig);
// struct c_field JNIGO_findObjField(struct c_object obj, const char* name,
// 		const char* rsig);
// //
// struct c_object JNIGO_newA(struct c_method m, const struct jval * args,
// 		int len);
// struct c_object JNIGO_callA(struct c_method m, const struct jval * args,
// 		int len);
// struct c_object JNIGO_get(struct c_field f);
// void JNIGO_set(struct c_field f, struct jval arg);
