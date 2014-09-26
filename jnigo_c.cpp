/*
 * jnigo_c.c
 *
 *  Created on: Sep 25, 2014
 *      Author: cny
 */

#include "jnigo.h"
//
JVM jvm_g;

#ifdef __cplusplus
extern "C" {
#endif
void JNIGO_setVersion(int v) {
	jvm_g.version = v;
}
void JNIGO_setIgnoreUnrecognized(int i) {
	jvm_g.ignoreUnrecognized = i;
}
void JNIGO_addOption(const char* option) {
	jvm_g.addOption(string(option));
}
int JNIGO_init() {
	return jvm_g.init();
}
void JNIGO_destory() {
	jvm_g.destory();
}
//void JNIGO_free(void **p) {
//	jvm_g.free((VmObj**) p);
//}
struct c_err JNIGO_errOccurred() {
	return jvm_g.errOccurred().toc();
}
void JNIGO_errClear() {
	jvm_g.errClear();
}
//
struct c_class JNIGO_findClass(const char* name) {
	return jvm_g.findClass(string(name)).toc();
}
//
struct c_method JNIGO_findClsMethod(struct c_class cls, const char* name, const char* vsig,
		const char* rsig, int static_) {
	Class rcls;
	rcls.fromc(cls);
	return rcls.findMethod(string(name), string(vsig), string(rsig), static_).toc();
}
struct c_field JNIGO_findClsField(struct c_class cls, const char* name, const char* rsig,
		int static_) {
	Class rcls;
	rcls.fromc(cls);
	return rcls.findField(string(name), string(rsig), static_).toc();
}
struct c_method JNIGO_findObjMethod(struct c_object obj, const char* name, const char* vsig,
		const char* rsig) {
	Object robj;
	robj.fromc(obj);
	return robj.findMethod(string(name), string(vsig), string(rsig)).toc();
}
struct c_field JNIGO_findObjField(struct c_object obj, const char* name, const char* rsig) {
	Object robj;
	robj.fromc(obj);
	return robj.findField(string(name), string(rsig)).toc();
}
//
struct c_object JNIGO_newA(struct c_method m, const struct jval * args, int len) {
	Method rm;
	rm.fromc(m);
	return rm.newA(args, len).toc();
}
struct c_object JNIGO_callA(struct c_method m, const struct jval * args, int len) {
	Method rm;
	rm.fromc(m);
	return rm.callA(args, len).toc();
}
struct c_object JNIGO_get(struct c_field f) {
	Field rf;
	rf.fromc(f);
	return rf.get().toc();
}
void JNIGO_set(struct c_field f, struct jval arg) {
	Field rf;
	rf.fromc(f);
	rf.set(arg);
}
#ifdef __cplusplus
}
#endif
