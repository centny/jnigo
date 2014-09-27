/*
 * jnigo.h
 *
 *  Created on: Sep 25, 2014
 *      Author: cny
 */

#ifndef JNIGO_H_
#define JNIGO_H_
#include <jni.h>
#include <string>
#include <vector>
#include "jnigo_c.h"
using namespace std;
//
class JVM;
class Class;
class Method;
class Field;
class Object;
class Err;
///
///
class VmObj {
public:
	int valid;
	string msg;
	JVM *jvm;
public:
	VmObj();
	void init(JVM *jvm, string msg);
	void covjval(jvalue *vals, const jval * args, int len);
	void fromc_(c_vmo c);
	c_vmo toc_();
};
//class Err: public VmObj {
//public:
//	JVM *jvm;
//	jthrowable err;
//public:
//	Err();
//	void fromc(c_err c);
//	c_err toc();
//};
class JVM: public VmObj {
public:
	JavaVM *jvm_; /* denotes a Java VM */
	JNIEnv *env_; /* pointer to native method interface */
	int version;
	int ignoreUnrecognized;
	vector<string> options; /* the VM option */
	JavaVMOption *options_;
public:
	JVM();
	virtual ~JVM();
	//
	void addOption(string option);
	int init();
	void destory();
	Class findClass(string name);
	Object errOccurred();
	void errClear();
//	void free(VmObj **v);
	Object newAry(string sig, int len);
	Object newS(const jbyte* bys, int len);
};
///
class Class: public VmObj {
public:
	jclass cls;
	string name;
public:
	Class();
	Method findMethod(string name, string vsig, string rsig, int static_ = 1);
	Field findField(string name, string rsig, int static_ = 1);
	void fromc(c_class c);
	c_class toc();
};
///
class Object: public VmObj {
public:
	Class cls;
	string sig;
	jval val;
public:
	Object();
	Method findMethod(string name, string vsig, string rsig);
	Field findField(string name, string rsig);
	//
	void set(int idx, jval val);
	Object get(int idx);
	//
	int len();
	void cary(void* bs, int idx, int len);
	void sary(void* bs, int idx, int len);
	//
	Object callA(string name, string vsig, string rsig, const jval * args,
			int len);
	//
	Object as(string name);
	//
	void fromc(c_object c);
	c_object toc();
};

class Method: public VmObj {
public:
	Class cls;
	Object obj;
	jmethodID mid;
	string name;
	string vsig;
	string rsig;
	int static_;
public:
	Method();
	Object callA(const jval * args, int len);
	Object newA(const jval * args, int len);
	//
	void fromc(c_method c);
	c_method toc();
};
class Field: public VmObj {
public:
	Class cls;
	Object obj;
	jfieldID fid;
	string name;
	string rsig;
	int static_;
public:
	Field();
	void set(jval val);
	Object get();
	//
	void fromc(c_field c);
	c_field toc();
};
#endif /* JNIGO_H_ */
