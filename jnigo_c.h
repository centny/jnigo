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
typedef struct c_vmo {
	int valid;
	char msg[512];
	void* jvm;
	void* env;
} c_vmo;
typedef struct c_err {
	struct c_vmo vmo;
	//
	void* jvm;
	jthrowable err;
} c_err;
typedef struct c_class {
	struct c_vmo vmo;
	//
	void* jvm;
	jclass cls;
	char name[1024];
} c_class;
typedef struct c_object {
	struct c_vmo vmo;
	//
	void *jvm;
	struct c_class cls;
	char sig[1024];
	jval val;
} c_object;
typedef struct c_method {
	struct c_vmo vmo;
	//
	void* jvm;
	struct c_class cls;
	struct c_object obj;
	jmethodID mid;
	char name[1024];
	char vsig[1024];
	char rsig[1024];
	int static_;
} c_method;
typedef struct c_field {
	struct c_vmo vmo;
	//
	void* jvm;
	struct c_class cls;
	struct c_object obj;
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
struct c_err JNIGO_errOccurred();
void JNIGO_errClear();
//
struct c_class JNIGO_findClass(const char* name);
struct c_method JNIGO_findClsMethod(struct c_class cls, const char* name,
		const char* vsig, const char* rsig, int static_);
struct c_field JNIGO_findClsField(struct c_class cls, const char* name,
		const char* rsig, int static_);
struct c_method JNIGO_findObjMethod(struct c_object obj, const char* name,
		const char* vsig, const char* rsig);
struct c_field JNIGO_findObjField(struct c_object obj, const char* name,
		const char* rsig);
//
struct c_object JNIGO_newA(struct c_method m, const jval * args,
		int len);
struct c_object JNIGO_callA(struct c_method m, const jval * args,
		int len);
struct c_object JNIGO_get(struct c_field f);
void JNIGO_set(struct c_field f, jval arg);
//
#ifdef __cplusplus
}
#endif
//
#endif /* JNIGO_C_H_ */
