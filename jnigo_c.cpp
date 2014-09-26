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
c_object JNIGO_errOccurred() {
	return jvm_g.errOccurred().toc();
}
void JNIGO_errClear() {
	jvm_g.errClear();
}
c_object JNIGO_newAry(const char* sig, int len) {
	return jvm_g.newAry(string(sig), len).toc();
}
//
c_class JNIGO_findClass(const char* name) {
	return jvm_g.findClass(string(name)).toc();
}
//
c_method JNIGO_findClsMethod(c_class cls, const char* name, const char* vsig,
		const char* rsig, int static_) {
	Class rcls;
	rcls.fromc(cls);
	return rcls.findMethod(string(name), string(vsig), string(rsig), static_).toc();
}
c_field JNIGO_findClsField(c_class cls, const char* name, const char* rsig,
		int static_) {
	Class rcls;
	rcls.fromc(cls);
	return rcls.findField(string(name), string(rsig), static_).toc();
}
c_method JNIGO_findObjMethod(c_object obj, const char* name, const char* vsig,
		const char* rsig) {
	Object robj;
	robj.fromc(obj);
	return robj.findMethod(string(name), string(vsig), string(rsig)).toc();
}
c_field JNIGO_findObjField(c_object obj, const char* name, const char* rsig) {
	Object robj;
	robj.fromc(obj);
	return robj.findField(string(name), string(rsig)).toc();
}
int JNIGO_len(c_object obj) {
	Object robj;
	robj.fromc(obj);
	return robj.len();
}
void JNIGO_cary(c_object obj, void* buf, int idx, int len) {
	Object robj;
	robj.fromc(obj);
	robj.cary(buf, idx, len);
}
void JNIGO_sary(c_object obj, void* buf, int idx, int len) {
	Object robj;
	robj.fromc(obj);
	robj.sary(buf, idx, len);
}
//
c_object JNIGO_newS(const jbyte* bys, int len) {
	return jvm_g.newS(bys, len).toc();
}
c_object JNIGO_newA(c_method m, const jval * args, int len) {
	Method rm;
	rm.fromc(m);
	return rm.newA(args, len).toc();
}
c_object JNIGO_callA(c_method m, const jval * args, int len) {
	Method rm;
	rm.fromc(m);
	return rm.callA(args, len).toc();
}
c_object JNIGO_callA_o(c_object obj, const char* name, const char* vsig,
		const char* rsig, const jval * args, int len) {
	Object robj;
	robj.fromc(obj);
	return robj.callA(string(name), string(vsig), string(rsig), args, len).toc();
}
c_object JNIGO_get(c_field f) {
	Field rf;
	rf.fromc(f);
	return rf.get().toc();
}
c_object JNIGO_get_o(c_object obj, int idx) {
	Object robj;
	robj.fromc(obj);
	return robj.get(idx).toc();
}
void JNIGO_set_o(c_object obj, int idx, jval v) {
	Object robj;
	robj.fromc(obj);
	robj.set(idx, v);
}
void JNIGO_set(c_field f, jval arg) {
	Field rf;
	rf.fromc(f);
	rf.set(arg);
}
#ifdef __cplusplus
}
#endif
