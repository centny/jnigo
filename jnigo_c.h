/*
 * jnigo_c.h
 *
 *  Created on: Sep 25, 2014
 *      Author: cny
 */

#ifndef JNIGO_C_H_
#define JNIGO_C_H_
#include <jni.h>

//
#ifdef __cplusplus
extern "C" {
#endif
//
typedef struct {
	jboolean z;
	jbyte b;
	jchar c;
	jshort s;
	jint i;
	jlong j;
	jfloat f;
	jdouble d;
	jobject l;
	char typ;
} jval;
typedef struct {
	int valid;
	char msg[512];
	void* jvm;
} c_vmo;
//typedef struct {
//	c_vmo vmo;
//	//
//	void* jvm;
//	jthrowable err;
//} c_err;
typedef struct {
	c_vmo vmo;
	//
	void* jvm;
	jclass cls;
	char name[1024];
} c_class;
typedef struct {
	c_vmo vmo;
	//
	void *jvm;
	c_class cls;
	char sig[1024];
	jval val;
} c_object;
typedef struct {
	c_vmo vmo;
	//
	void* jvm;
	c_class cls;
	c_object obj;
	jmethodID mid;
	char name[1024];
	char vsig[1024];
	char rsig[1024];
	int static_;
} c_method;
typedef struct {
	c_vmo vmo;
	//
	void* jvm;
	c_class cls;
	c_object obj;
	jfieldID fid;
	char name[1024];
	char rsig[1024];
	int static_;
} c_field;
//
void JNIGO_setVersion(int v);
void JNIGO_setIgnoreUnrecognized(int i);
void JNIGO_addOption(const char* option);
int JNIGO_init();
void JNIGO_destory();
c_object JNIGO_errOccurred();
void JNIGO_errClear();
c_object JNIGO_newAry(const char* sig, int len);
//
c_class JNIGO_findClass(const char* name);
c_method JNIGO_findClsMethod(c_class cls, const char* name, const char* vsig,
		const char* rsig, int static_);
c_field JNIGO_findClsField(c_class cls, const char* name, const char* rsig,
		int static_);
c_method JNIGO_findObjMethod(c_object obj, const char* name, const char* vsig,
		const char* rsig);
c_field JNIGO_findObjField(c_object obj, const char* name, const char* rsig);
int JNIGO_len(c_object obj);
void JNIGO_cary(c_object obj, void* buf, int idx, int len);
void JNIGO_sary(c_object obj, void* buf, int idx, int len);
c_object JNIGO_as(c_object obj, const char* name);
//
c_object JNIGO_newS(const jbyte* bys, int len);
c_object JNIGO_newA(c_method m, const jval * args, int len);
c_object JNIGO_callA(c_method m, const jval * args, int len);
c_object JNIGO_callA_o(c_object obj, const char* name, const char* vsig,
		const char* rsig, const jval * args, int len);
c_object JNIGO_get(c_field f);
c_object JNIGO_get_o(c_object obj, int idx);
void JNIGO_set_o(c_object obj, int idx, jval v);
void JNIGO_set(c_field f, jval arg);
//
#ifdef __cplusplus
}
#endif
//
#endif /* JNIGO_C_H_ */
