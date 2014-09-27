#include "jnigo.h"
#include <sstream>
#include <iostream>
using namespace std;
VmObj::VmObj() {
	this->jvm = 0;
	this->valid = false;
	this->msg = "INVALID";
}
void VmObj::covjval(jvalue *vals, const jval * args, int len) {
	int i = 0;
	for (i = 0; i < len; i++) {
		switch (args[i].typ) {
		case 'Z':
			vals[i].z = args[i].z;
			break;
		case 'B':
			vals[i].b = args[i].b;
			break;
		case 'C':
			vals[i].c = args[i].c;
			break;
		case 'S':
			vals[i].s = args[i].s;
			break;
		case 'I':
			vals[i].i = args[i].i;
			break;
		case 'J':
			vals[i].j = args[i].j;
			break;
		case 'F':
			vals[i].f = args[i].f;
			break;
		case 'D':
			vals[i].d = args[i].d;
			break;
		default:
			vals[i].l = args[i].l;
			break;
		}
	}
}
void VmObj::init(JVM *jvm, string msg) {
	this->jvm = jvm;
	this->valid = msg.size() == 0;
	this->msg = msg;
}
void VmObj::fromc_(c_vmo c) {
	this->jvm = (JVM *) c.jvm;
	this->msg = string(c.msg);
	this->valid = c.valid;
}
c_vmo VmObj::toc_() {
	c_vmo vmo;
	vmo.jvm = this->jvm;
	vmo.valid = this->valid;
	this->msg.copy(vmo.msg, this->msg.size());
	vmo.msg[this->msg.size()] = 0;
	return vmo;
}
//Err::Err() {
//	this->jvm = 0;
//	this->err = 0;
//}
//void Err::fromc(c_err c) {
//	this->fromc_(c.vmo);
//	this->err = c.err;
//	this->jvm = (JVM*) c.jvm;
//}
//c_err Err::toc() {
//	c_err c;
//	c.err = this->err;
//	c.jvm = this->jvm;
//	c.vmo = this->toc_();
//	return c;
//}
/* ------ JVM ------ */
JVM::JVM() {
	this->jvm = this;
	this->jvm_ = 0;
	this->env_ = 0;
	this->valid = false;
	this->msg = "";
	this->jvm = 0;
	this->options_ = 0;
	this->version = JNI_VERSION_1_6;
	this->ignoreUnrecognized = JNI_TRUE;
}
JVM::~JVM() {
}
void JVM::addOption(string option) {
	this->options.push_back(option);
}

int JVM::init() {
	JavaVMInitArgs vm_args;
	int olen = this->options.size();
	if (olen > 0) {
		this->options_ = new JavaVMOption[this->options.size()];
		vm_args.nOptions = olen;
		vm_args.options = this->options_;
		for (int i = 0; i < olen; i++) {
			this->options_[i].optionString = (char*) this->options[i].c_str();
		}
	}
	vm_args.version = this->version;
	vm_args.ignoreUnrecognized = this->ignoreUnrecognized;
	int status = JNI_CreateJavaVM(&this->jvm_, (void**) &this->env_, &vm_args);
	if (status == JNI_OK) {
		this->valid = true;
		this->msg = "";
	} else {
		this->valid = false;
		stringstream ss;
		ss << "initial error:" << status;
		this->msg = ss.str();
	}
	return status;
}
void JVM::destory() {
	if (this->jvm) {
		this->destory();
		this->env_ = 0;
		this->jvm_ = 0;
	}
	if (this->options_) {
		delete (this->options_);
		this->options_ = 0;
	}
	this->valid = false;
	this->msg = "destory";
}
Class JVM::findClass(string name) {
	Class ncls;
	ncls.cls = this->env_->FindClass(name.c_str());
	ncls.name = name;
	ncls.jvm = this;
	if (ncls.cls) {
		ncls.init(this, "");
	} else {
		ncls.init(this, "class not found by name:" + name);
	}
	return ncls;
}
Object JVM::errOccurred() {
	Object err;
	err.jvm = this;
	err.sig = "Ljava/lang/Throwable;";
	err.cls = this->findClass(err.sig);
	if (this->env_->ExceptionCheck()) {
		err.val.l = this->env_->ExceptionOccurred();
		err.val.typ = 'L';
		err.init(this, "");
	} else {
		err.init(this, "not error");
	}
	return err;
}
void JVM::errClear() {
	this->env_->ExceptionClear();
}
//void JVM::free(VmObj **v) {
//	if (v && *v) {
//		delete *v;
//		*v = NULL;
//	}
//}
Object JVM::newAry(string sig, int len) {
	Object obj;
	obj.sig = "[" + sig;
	obj.cls = this->findClass(obj.sig.c_str());
	if (!obj.cls.valid) {
		obj.init(this, "class not found by sig:" + obj.sig);
		return obj;
	}
	if ("Z" == sig) {
		obj.val.l = this->env_->NewBooleanArray(len);
		obj.val.typ = 'L';
	} else if ("B" == sig) {
		obj.val.l = this->env_->NewByteArray(len);
		obj.val.typ = 'L';
	} else if ("C" == sig) {
		obj.val.l = this->env_->NewCharArray(len);
		obj.val.typ = 'L';
	} else if ("S" == sig) {
		obj.val.l = this->env_->NewShortArray(len);
		obj.val.typ = 'L';
	} else if ("I" == sig) {
		obj.val.l = this->env_->NewIntArray(len);
		obj.val.typ = 'L';
	} else if ("J" == sig) {
		obj.val.l = this->env_->NewLongArray(len);
		obj.val.typ = 'L';
	} else if ("F" == sig) {
		obj.val.l = this->env_->NewFloatArray(len);
		obj.val.typ = 'L';
	} else if ("D" == sig) {
		obj.val.l = this->env_->NewDoubleArray(len);
		obj.val.typ = 'L';
	} else {
		Class scls = this->findClass(sig);
		obj.val.l = this->env_->NewObjectArray(len, scls.cls, 0);
		obj.val.typ = 'L';
	}
	obj.init(this, "");
	return obj;
}
Object JVM::newS(const jbyte* bys, int len) {
	Object ary = this->newAry("B", len);
	if (len) {
		ary.sary((void*) bys, 0, len);
	}
	Class cls = this->findClass("Ljava/lang/String;");
	Method m = cls.findMethod("<init>", "[B", "V", false);
	jval val = ary.val;
	return m.newA(&val, 1);
}
/* ------ JVM End ------ */

/* ------ Class ------ */
Class::Class() {
	this->jvm = 0;
	this->cls = 0;
}
Method Class::findMethod(string name, string vsig, string rsig, int static_) {
	Method m;
	if (!this->valid) {
		Method m;
		m.init(this->jvm, "can't call find method for invalid class");
		return m;
	}
	m.jvm = this->jvm;
	m.cls = *this;
	m.name = name;
	m.vsig = vsig;
	m.rsig = rsig;
	m.static_ = static_;
	string sig_ = "(" + vsig + ")" + rsig;
	if (static_) {
		m.mid = this->jvm->env_->GetStaticMethodID(this->cls, name.c_str(),
				sig_.c_str());
	} else {
		m.mid = this->jvm->env_->GetMethodID(this->cls, name.c_str(),
				sig_.c_str());
	}
	if (m.mid) {
		m.init(this->jvm, "");
	} else {
		m.init(this->jvm, "method(" + m.name + ") not found by sig:" + sig_);
	}
	return m;
}
Field Class::findField(string name, string rsig, int static_) {
	if (!this->valid) {
		Field f;
		f.init(this->jvm, "can't call find field for invalid class");
		return f;
	}
	Field f;
	f.jvm = this->jvm;
	f.cls = *this;
	f.name = name;
	f.rsig = rsig;
	f.static_ = static_;
	if (static_) {
		f.fid = this->jvm->env_->GetStaticFieldID(this->cls, name.c_str(),
				rsig.c_str());
	} else {
		f.fid = this->jvm->env_->GetFieldID(this->cls, name.c_str(),
				rsig.c_str());
	}
	if (f.fid) {
		f.init(this->jvm, "");
	} else {
		f.init(this->jvm, "field(" + this->name + ") not found by sig:" + rsig);
	}
	return f;
}
void Class::fromc(c_class c) {
	this->fromc_(c.vmo);
	this->jvm = (JVM*) c.jvm;
	this->cls = c.cls;
	this->name = string(c.name);
}
c_class Class::toc() {
	c_class c;
	c.vmo = this->toc_();
	c.jvm = this->jvm;
	c.cls = this->cls;
	this->name.copy(c.name, this->name.size());
	c.name[this->name.size()] = 0;
	return c;
}
/* ------ Class End ------ */

/* ------ Object ------ */
Object::Object() {
	this->jvm = 0;
}
Method Object::findMethod(string name, string vsig, string rsig) {
	Method m = this->cls.findMethod(name, vsig, rsig, false);
	m.obj = *this;
	return m;
}
Field Object::findField(string name, string rsig) {
	Field f = this->cls.findField(name, rsig, false);
	f.obj = *this;
	return f;
}
void Object::set(int idx, jval val) {
	if ("Z" == this->sig) {
		this->jvm->env_->SetBooleanArrayRegion((jbooleanArray) this->val.l, idx,
				1, &val.z);
	} else if ("B" == this->sig) {
		this->jvm->env_->SetByteArrayRegion((jbyteArray) this->val.l, idx, 1,
				&val.b);
	} else if ("C" == this->sig) {
		this->jvm->env_->SetCharArrayRegion((jcharArray) this->val.l, idx, 1,
				&val.c);
	} else if ("S" == this->sig) {
		this->jvm->env_->SetShortArrayRegion((jshortArray) this->val.l, idx, 1,
				&val.s);
	} else if ("I" == this->sig) {
		this->jvm->env_->SetIntArrayRegion((jintArray) this->val.l, idx, 1,
				&val.i);
	} else if ("J" == this->sig) {
		this->jvm->env_->SetLongArrayRegion((jlongArray) this->val.l, idx, 1,
				&val.j);
	} else if ("F" == this->sig) {
		this->jvm->env_->SetFloatArrayRegion((jfloatArray) this->val.l, idx, 1,
				&val.f);
	} else if ("D" == this->sig) {
		this->jvm->env_->SetDoubleArrayRegion((jdoubleArray) this->val.l, idx,
				1, &val.d);
	} else {
		this->jvm->env_->SetObjectArrayElement((jobjectArray) this->val.l, idx,
				val.l);
	}
}
Object Object::get(int idx) {
	Object obj;
	obj.sig = this->sig.replace(0, 1, "");
	obj.cls = this->jvm->findClass(obj.sig);
	if (obj.cls.valid) {
		obj.init(this->jvm, obj.cls.msg);
		return obj;
	}
	jval val;
	if ("Z" == this->sig) {
		this->jvm->env_->GetBooleanArrayRegion((jbooleanArray) this->val.l, idx,
				1, &val.z);
		val.typ = 'Z';
	} else if ("B" == this->sig) {
		this->jvm->env_->GetByteArrayRegion((jbyteArray) this->val.l, idx, 1,
				&val.b);
		val.typ = 'B';
	} else if ("C" == this->sig) {
		this->jvm->env_->GetCharArrayRegion((jcharArray) this->val.l, idx, 1,
				&val.c);
		val.typ = 'C';
	} else if ("S" == this->sig) {
		this->jvm->env_->GetShortArrayRegion((jshortArray) this->val.l, idx, 1,
				&val.s);
		val.typ = 'S';
	} else if ("I" == this->sig) {
		this->jvm->env_->GetIntArrayRegion((jintArray) this->val.l, idx, 1,
				&val.i);
		val.typ = 'I';
	} else if ("J" == this->sig) {
		this->jvm->env_->GetLongArrayRegion((jlongArray) this->val.l, idx, 1,
				&val.j);
		val.typ = 'J';
	} else if ("F" == this->sig) {
		this->jvm->env_->GetFloatArrayRegion((jfloatArray) this->val.l, idx, 1,
				&val.f);
		val.typ = 'F';
	} else if ("D" == this->sig) {
		this->jvm->env_->GetDoubleArrayRegion((jdoubleArray) this->val.l, idx,
				1, &val.d);
		val.typ = 'D';
	} else {
		val.l = this->jvm->env_->GetObjectArrayElement(
				(jobjectArray) this->val.l, idx);
		val.typ = 'L';
	}
	obj.val = val;
	return obj;
}
int Object::len() {
	return this->jvm->env_->GetArrayLength((jarray) this->val.l);
}
void Object::cary(void* bs, int idx, int len) {
	if ("[Z" == this->sig) {
		this->jvm->env_->GetBooleanArrayRegion((jbooleanArray) this->val.l, idx,
				len, (jboolean*) bs);
	} else if ("[B" == this->sig) {
		this->jvm->env_->GetByteArrayRegion((jbyteArray) this->val.l, idx, len,
				(jbyte*) bs);
	} else if ("[C" == this->sig) {
		this->jvm->env_->GetCharArrayRegion((jcharArray) this->val.l, idx, len,
				(jchar*) bs);
	} else if ("[S" == this->sig) {
		this->jvm->env_->GetShortArrayRegion((jshortArray) this->val.l, idx,
				len, (jshort*) bs);
	} else if ("[I" == this->sig) {
		this->jvm->env_->GetIntArrayRegion((jintArray) this->val.l, idx, len,
				(jint*) bs);
	} else if ("[J" == this->sig) {
		this->jvm->env_->GetLongArrayRegion((jlongArray) this->val.l, idx, len,
				(jlong*) bs);
	} else if ("[F" == this->sig) {
		this->jvm->env_->GetFloatArrayRegion((jfloatArray) this->val.l, idx,
				len, (jfloat*) bs);
	} else if ("[D" == this->sig) {
		this->jvm->env_->GetDoubleArrayRegion((jdoubleArray) this->val.l, idx,
				len, (jdouble*) bs);
	}
}
void Object::sary(void* bs, int idx, int len) {
	if ("[Z" == this->sig) {
		this->jvm->env_->SetBooleanArrayRegion((jbooleanArray) this->val.l, idx,
				len, (jboolean*) bs);
	} else if ("[B" == this->sig) {
		this->jvm->env_->SetByteArrayRegion((jbyteArray) this->val.l, idx, len,
				(jbyte*) bs);
	} else if ("[C" == this->sig) {
		this->jvm->env_->SetCharArrayRegion((jcharArray) this->val.l, idx, len,
				(jchar*) bs);
	} else if ("[S" == this->sig) {
		this->jvm->env_->SetShortArrayRegion((jshortArray) this->val.l, idx,
				len, (jshort*) bs);
	} else if ("[I" == this->sig) {
		this->jvm->env_->SetIntArrayRegion((jintArray) this->val.l, idx, len,
				(jint*) bs);
	} else if ("[J" == this->sig) {
		this->jvm->env_->SetLongArrayRegion((jlongArray) this->val.l, idx, len,
				(jlong*) bs);
	} else if ("[F" == this->sig) {
		this->jvm->env_->SetFloatArrayRegion((jfloatArray) this->val.l, idx,
				len, (jfloat*) bs);
	} else if ("[D" == this->sig) {
		this->jvm->env_->SetDoubleArrayRegion((jdoubleArray) this->val.l, idx,
				len, (jdouble*) bs);
	}
}
Object Object::callA(string name, string vsig, string rsig, const jval * args,
		int len) {
	Method m = this->findMethod(name, vsig, rsig);
	if (m.valid) {
		return m.callA(args, len);
	} else {
		Object obj;
		obj.init(this->jvm, m.msg);
		return obj;
	}
}
Object Object::as(string name) {
	Object obj = *this;
	obj.cls = this->jvm->findClass(name);
	if (obj.cls.valid) {
		obj.sig = name;
		return obj;
	} else {
		obj.init(this->jvm, obj.cls.msg);
		return obj;
	}
}
void Object::fromc(c_object c) {
	this->fromc_(c.vmo);
	this->jvm = (JVM*) c.jvm;
	this->cls.fromc(c.cls);
	this->sig = string(c.sig);
	this->val = c.val;
}
c_object Object::toc() {
	c_object c;
	c.vmo = this->toc_();
	c.jvm = this->jvm;
	c.cls = this->cls.toc();
	this->sig.copy(c.sig, this->sig.size());
	c.sig[this->sig.size()] = 0;
	c.val = this->val;
	return c;
}
/* ------ Object End ------ */

/* ------ Method ------ */
Method::Method() {
	this->jvm = 0;
	this->mid = 0;
	this->static_ = false;
}
Object Method::callA(const jval * args, int len) {
	Object obj;
	obj.sig = this->rsig;
	char tsig = 0;
	if (!this->rsig.empty()) {
		tsig = this->rsig.at(0);
	}
	if ('[' == tsig || 'L' == tsig) {
		obj.cls = this->jvm->findClass(obj.sig);
		if (obj.cls.valid) {
			obj.init(this->jvm, "");
		} else {
			obj.init(this->jvm, obj.cls.msg);
			return obj;
		}
	} else {
		obj.cls.init(this->jvm, "INVALID");
	}
	jvalue *vals = 0;
	if (len) {
		vals = new jvalue[len];
		this->covjval(vals, args, len);
	}
	jval val;
	if (this->static_) {
		if ("V" == this->rsig) {
			this->jvm->env_->CallStaticVoidMethodA(this->cls.cls, this->mid,
					vals);
			val.typ = 'V';
		} else if ("Z" == this->rsig) {
			val.z = this->jvm->env_->CallStaticBooleanMethodA(this->cls.cls,
					this->mid, vals);
			val.typ = 'Z';
		} else if ("B" == this->rsig) {
			val.b = this->jvm->env_->CallStaticByteMethodA(this->cls.cls,
					this->mid, vals);
			val.typ = 'B';
		} else if ("C" == this->rsig) {
			val.c = this->jvm->env_->CallStaticCharMethodA(this->cls.cls,
					this->mid, vals);
			val.typ = 'C';
		} else if ("S" == this->rsig) {
			val.s = this->jvm->env_->CallStaticShortMethodA(this->cls.cls,
					this->mid, vals);
			val.typ = 'S';
		} else if ("I" == this->rsig) {
			val.i = this->jvm->env_->CallStaticIntMethodA(this->cls.cls,
					this->mid, vals);
			val.typ = 'I';
		} else if ("J" == this->rsig) {
			val.j = this->jvm->env_->CallStaticLongMethodA(this->cls.cls,
					this->mid, vals);
			val.typ = 'J';
		} else if ("F" == this->rsig) {
			val.f = this->jvm->env_->CallStaticFloatMethodA(this->cls.cls,
					this->mid, vals);
			val.typ = 'F';
		} else if ("D" == this->rsig) {
			val.d = this->jvm->env_->CallStaticDoubleMethodA(this->cls.cls,
					this->mid, vals);
			val.typ = 'D';
		} else {
			val.l = this->jvm->env_->CallStaticObjectMethodA(this->cls.cls,
					this->mid, vals);
			val.typ = 'L';
		}
	} else {
		if ("V" == this->rsig) {
			this->jvm->env_->CallVoidMethodA(this->obj.val.l, this->mid, vals);
			val.typ = 'V';
		} else if ("Z" == this->rsig) {
			val.z = this->jvm->env_->CallBooleanMethodA(this->obj.val.l,
					this->mid, vals);
			val.typ = 'Z';
		} else if ("B" == this->rsig) {
			val.b = this->jvm->env_->CallByteMethodA(this->obj.val.l, this->mid,
					vals);
			val.typ = 'B';
		} else if ("C" == this->rsig) {
			val.c = this->jvm->env_->CallCharMethodA(this->obj.val.l, this->mid,
					vals);
			val.typ = 'C';
		} else if ("S" == this->rsig) {
			val.s = this->jvm->env_->CallShortMethodA(this->obj.val.l,
					this->mid, vals);
			val.typ = 'S';
		} else if ("I" == this->rsig) {
			val.i = this->jvm->env_->CallIntMethodA(this->obj.val.l, this->mid,
					vals);
			val.typ = 'I';
		} else if ("J" == this->rsig) {
			val.j = this->jvm->env_->CallLongMethodA(this->obj.val.l, this->mid,
					vals);
			val.typ = 'J';
		} else if ("F" == this->rsig) {
			val.f = this->jvm->env_->CallFloatMethodA(this->obj.val.l,
					this->mid, vals);
			val.typ = 'F';
		} else if ("D" == this->rsig) {
			val.d = this->jvm->env_->CallDoubleMethodA(this->obj.val.l,
					this->mid, vals);
			val.typ = 'D';
		} else {
			val.l = this->jvm->env_->CallObjectMethodA(this->obj.val.l,
					this->mid, vals);
			val.typ = 'L';
		}
	}
	if (vals) {
		delete vals;
		vals = 0;
	}
	obj.val = val;
	obj.init(this->jvm, "");
	return obj;
}
Object Method::newA(const jval * args, int len) {
	Object obj;
	obj.sig = this->cls.name;
	obj.cls = this->cls;
	jvalue *vals = 0;
	if (len) {
		vals = new jvalue[len];
		this->covjval(vals, args, len);
	}
	jval val;
	val.typ = 'L';
	val.l = this->jvm->env_->NewObjectA(this->cls.cls, this->mid, vals);
	if (vals) {
		delete vals;
		vals = 0;
	}
	obj.val = val;
	obj.init(this->jvm, "");
	return obj;
}
void Method::fromc(c_method c) {
	this->fromc_(c.vmo);
	this->jvm = (JVM*) c.jvm;
	this->cls.fromc(c.cls);
	this->obj.fromc(c.obj);
	this->mid = c.mid;
	this->name = string(c.name);
	this->vsig = string(c.vsig);
	this->rsig = string(c.rsig);
	this->static_ = c.static_;
}
c_method Method::toc() {
	c_method c;
	c.vmo = this->toc_();
	c.jvm = this->jvm;
	c.cls = this->cls.toc();
	c.obj = this->obj.toc();
	c.mid = this->mid;
	this->name.copy(c.name, this->name.size());
	c.name[this->name.size()] = 0;
	this->vsig.copy(c.vsig, this->vsig.size());
	c.vsig[this->vsig.size()] = 0;
	this->rsig.copy(c.rsig, this->rsig.size());
	c.rsig[this->rsig.size()] = 0;
	c.static_ = this->static_;
	return c;
}
/* ------ Method End ------ */

/* ------ Field ------ */
Field::Field() {
	this->jvm = 0;
	this->fid = 0;
	this->static_ = false;
}
Object Field::get() {
	Object obj;
	obj.sig = this->rsig;
	obj.cls = this->jvm->findClass(obj.sig);
	if (obj.cls.valid) {
		obj.init(this->jvm, "");
	} else {
		obj.init(this->jvm, obj.cls.msg);
		return obj;
	}
	jval val;
	if (this->static_) {
		if ("Z" == this->rsig) {
			val.z = this->jvm->env_->GetStaticBooleanField(this->cls.cls,
					this->fid);
			val.typ = 'Z';
		} else if ("B" == this->rsig) {
			val.b = this->jvm->env_->GetStaticByteField(this->cls.cls,
					this->fid);
			val.typ = 'B';
		} else if ("C" == this->rsig) {
			val.c = this->jvm->env_->GetStaticCharField(this->cls.cls,
					this->fid);
			val.typ = 'C';
		} else if ("S" == this->rsig) {
			val.s = this->jvm->env_->GetStaticShortField(this->cls.cls,
					this->fid);
			val.typ = 'S';
		} else if ("I" == this->rsig) {
			val.i = this->jvm->env_->GetStaticIntField(this->cls.cls,
					this->fid);
			val.typ = 'I';
		} else if ("J" == this->rsig) {
			val.j = this->jvm->env_->GetStaticLongField(this->cls.cls,
					this->fid);
			val.typ = 'L';
		} else if ("F" == this->rsig) {
			val.f = this->jvm->env_->GetStaticFloatField(this->cls.cls,
					this->fid);
			val.typ = 'F';
		} else if ("D" == this->rsig) {
			val.d = this->jvm->env_->GetStaticDoubleField(this->cls.cls,
					this->fid);
			val.typ = 'D';
		} else {
			val.l = this->jvm->env_->GetStaticObjectField(this->cls.cls,
					this->fid);
			val.typ = 'L';
		}
	} else {
		if ("Z" == this->rsig) {
			val.z = this->jvm->env_->GetBooleanField(this->obj.val.l,
					this->fid);
			val.typ = 'Z';
		} else if ("B" == this->rsig) {
			val.b = this->jvm->env_->GetByteField(this->obj.val.l, this->fid);
			val.typ = 'B';
		} else if ("C" == this->rsig) {
			val.c = this->jvm->env_->GetCharField(this->obj.val.l, this->fid);
			val.typ = 'C';
		} else if ("S" == this->rsig) {
			val.s = this->jvm->env_->GetShortField(this->obj.val.l, this->fid);
			val.typ = 'S';
		} else if ("I" == this->rsig) {
			val.i = this->jvm->env_->GetIntField(this->obj.val.l, this->fid);
			val.typ = 'I';
		} else if ("J" == this->rsig) {
			val.j = this->jvm->env_->GetLongField(this->obj.val.l, this->fid);
			val.typ = 'j';
		} else if ("F" == this->rsig) {
			val.f = this->jvm->env_->GetFloatField(this->obj.val.l, this->fid);
			val.typ = 'F';
		} else if ("D" == this->rsig) {
			val.d = this->jvm->env_->GetDoubleField(this->obj.val.l, this->fid);
			val.typ = 'D';
		} else {
			val.l = this->jvm->env_->GetObjectField(this->obj.val.l, this->fid);
			val.typ = 'L';
		}
	}
	obj.val = val;
	return obj;
}
void Field::set(jval val) {
	if (this->static_) {
		if ("B" == this->rsig) {
			this->jvm->env_->SetStaticByteField(this->cls.cls, this->fid,
					val.z);
		} else if ("C" == this->rsig) {
			this->jvm->env_->SetStaticCharField(this->cls.cls, this->fid,
					val.c);
		} else if ("S" == this->rsig) {
			this->jvm->env_->SetStaticShortField(this->cls.cls, this->fid,
					val.s);
		} else if ("I" == this->rsig) {
			this->jvm->env_->SetStaticIntField(this->cls.cls, this->fid, val.i);
		} else if ("J" == this->rsig) {
			this->jvm->env_->SetStaticLongField(this->cls.cls, this->fid,
					val.j);
		} else if ("F" == this->rsig) {
			this->jvm->env_->SetStaticFloatField(this->cls.cls, this->fid,
					val.f);
		} else if ("D" == this->rsig) {
			this->jvm->env_->SetStaticDoubleField(this->cls.cls, this->fid,
					val.d);
		} else {
			this->jvm->env_->SetStaticObjectField(this->cls.cls, this->fid,
					val.l);
		}
	} else {
		if ("Z" == this->rsig) {
			this->jvm->env_->SetBooleanField(this->obj.val.l, this->fid, val.z);
		} else if ("B" == this->rsig) {
			this->jvm->env_->SetByteField(this->obj.val.l, this->fid, val.b);
		} else if ("C" == this->rsig) {
			this->jvm->env_->SetCharField(this->obj.val.l, this->fid, val.c);
		} else if ("S" == this->rsig) {
			this->jvm->env_->SetShortField(this->obj.val.l, this->fid, val.s);
		} else if ("I" == this->rsig) {
			this->jvm->env_->SetIntField(this->obj.val.l, this->fid, val.i);
		} else if ("J" == this->rsig) {
			this->jvm->env_->SetLongField(this->obj.val.l, this->fid, val.j);
		} else if ("F" == this->rsig) {
			this->jvm->env_->SetFloatField(this->obj.val.l, this->fid, val.f);
		} else if ("D" == this->rsig) {
			this->jvm->env_->SetDoubleField(this->obj.val.l, this->fid, val.d);
		} else {
			this->jvm->env_->SetObjectField(this->obj.val.l, this->fid, val.l);
		}
	}
}
void Field::fromc(c_field c) {
	this->fromc_(c.vmo);
	this->jvm = (JVM*) c.jvm;
	this->cls.fromc(c.cls);
	this->obj.fromc(c.obj);
	this->fid = c.fid;
	this->name = string(c.name);
	this->rsig = string(c.rsig);
	this->static_ = c.static_;
}
c_field Field::toc() {
	c_field c;
	c.vmo = this->toc_();
	c.jvm = this->jvm;
	c.cls = this->cls.toc();
	c.obj = this->obj.toc();
	c.fid = this->fid;
	this->name.copy(c.name, this->name.size());
	c.name[this->name.size()] = 0;
	this->rsig.copy(c.rsig, this->rsig.size());
	c.rsig[this->rsig.size()] = 0;
	c.static_ = this->static_;
	return c;
}
