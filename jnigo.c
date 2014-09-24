/*
 ============================================================================
 Name        : jnigo.c
 Author      : Centny
 Version     :
 Copyright   : Your copyright notice
 Description : Hello World in C, Ansi-style
 ============================================================================
 */

#include <stdio.h>
#include <stdlib.h>
#include <jni.h>
#include "jnigo.h"
void covjval(jvalue *vals, const jval * args, int len) {
	int i = 0;
	for (i = 0; i < len; i++) {
		switch (args[i].typ) {
		case 0:
			vals[i].z = args[i].z;
			break;
		case 1:
			vals[i].b = args[i].b;
			break;
		case 2:
			vals[i].c = args[i].c;
			break;
		case 3:
			vals[i].s = args[i].s;
			break;
		case 4:
			vals[i].i = args[i].i;
			break;
		case 5:
			vals[i].j = args[i].j;
			break;
		case 6:
			vals[i].f = args[i].f;
			break;
		case 7:
			vals[i].d = args[i].d;
			break;
		default:
			vals[i].l = args[i].l;
			break;
		}
	}
}
jclass JNIGO_FindClass(JNIEnv *env, const char *name) {
	return (*env)->FindClass(env, name);
}
jint JNIGO_DestroyJavaVM(JavaVM *vm) {
	return (*vm)->DestroyJavaVM(vm);
}
jobject JNIGO_NewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->NewObjectA(env, clazz, methodID, vals);
	} else {
		return (*env)->NewObjectA(env, clazz, methodID, 0);
	}
}
jthrowable JNIGO_ExceptionOccurred(JNIEnv *env) {
	return (*env)->ExceptionOccurred(env);
}
void JNIGO_ExceptionClear(JNIEnv *env) {
	return (*env)->ExceptionClear(env);
}
//
//
jfieldID JNIGO_GetFieldID(JNIEnv *env, jclass clazz, const char *name,
		const char *sig) {
	return (*env)->GetFieldID(env, clazz, name, sig);
}
jfieldID JNIGO_GetStaticFieldID(JNIEnv *env, jclass clazz, const char *name,
		const char *sig) {
	return (*env)->GetStaticFieldID(env, clazz, name, sig);
}
//
jobject JNIGO_GetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetObjectField(env, obj, fieldID);
}
jboolean JNIGO_GetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetBooleanField(env, obj, fieldID);
}
jbyte JNIGO_GetByteField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetByteField(env, obj, fieldID);
}
jchar JNIGO_GetCharField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetCharField(env, obj, fieldID);
}
jshort JNIGO_GetShortField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetShortField(env, obj, fieldID);
}
jint JNIGO_GetIntField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetIntField(env, obj, fieldID);
}
jlong JNIGO_GetLongField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetLongField(env, obj, fieldID);
}
jfloat JNIGO_GetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetFloatField(env, obj, fieldID);
}
jdouble JNIGO_GetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetDoubleField(env, obj, fieldID);
}
jobject JNIGO_GetStaticObjectField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetStaticObjectField(env, obj, fieldID);
}
jboolean JNIGO_GetStaticBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetStaticBooleanField(env, obj, fieldID);
}
jbyte JNIGO_GetStaticByteField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetStaticByteField(env, obj, fieldID);
}
jchar JNIGO_GetStaticCharField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetStaticCharField(env, obj, fieldID);
}
jshort JNIGO_GetStaticShortField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetStaticShortField(env, obj, fieldID);
}
jint JNIGO_GetStaticIntField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetStaticIntField(env, obj, fieldID);
}
jlong JNIGO_GetStaticLongField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetStaticLongField(env, obj, fieldID);
}
jfloat JNIGO_GetStaticFloatField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetStaticFloatField(env, obj, fieldID);
}
jdouble JNIGO_GetStaticDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID) {
	return (*env)->GetStaticDoubleField(env, obj, fieldID);
}
//
void JNIGO_SetVal(JNIEnv *env, jobject obj, jfieldID fieldID, jval val) {
	switch (val.typ) {
	case 0:
		(*env)->SetBooleanField(env, obj, fieldID, val.z);
		break;
	case 1:
		(*env)->SetByteField(env, obj, fieldID, val.b);
		break;
	case 2:
		(*env)->SetCharField(env, obj, fieldID, val.c);
		break;
	case 3:
		(*env)->SetShortField(env, obj, fieldID, val.s);
		break;
	case 4:
		(*env)->SetIntField(env, obj, fieldID, val.i);
		break;
	case 5:
		(*env)->SetLongField(env, obj, fieldID, val.j);
		break;
	case 6:
		(*env)->SetFloatField(env, obj, fieldID, val.f);
		break;
	case 7:
		(*env)->SetDoubleField(env, obj, fieldID, val.d);
		break;
	default:
		(*env)->SetObjectField(env, obj, fieldID, val.l);
		break;
	}
}
void JNIGO_SetStaticVal(JNIEnv *env, jobject obj, jfieldID fieldID, jval val) {
	switch (val.typ) {
	case 0:
		(*env)->SetStaticBooleanField(env, obj, fieldID, val.z);
		break;
	case 1:
		(*env)->SetStaticByteField(env, obj, fieldID, val.b);
		break;
	case 2:
		(*env)->SetStaticCharField(env, obj, fieldID, val.c);
		break;
	case 3:
		(*env)->SetStaticShortField(env, obj, fieldID, val.s);
		break;
	case 4:
		(*env)->SetStaticIntField(env, obj, fieldID, val.i);
		break;
	case 5:
		(*env)->SetStaticLongField(env, obj, fieldID, val.j);
		break;
	case 6:
		(*env)->SetStaticFloatField(env, obj, fieldID, val.f);
		break;
	case 7:
		(*env)->SetStaticDoubleField(env, obj, fieldID, val.d);
		break;
	default:
		(*env)->SetStaticObjectField(env, obj, fieldID, val.l);
		break;
	}
}
//
//
jmethodID JNIGO_GetMethodID(JNIEnv *env, jclass clazz, const char *name,
		const char *sig) {
	return (*env)->GetMethodID(env, clazz, name, sig);
}
jmethodID JNIGO_GetStaticMethodID(JNIEnv *env, jclass clazz, const char *name,
		const char *sig) {
	return (*env)->GetStaticMethodID(env, clazz, name, sig);
}
//
jobject JNIGO_CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallObjectMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallObjectMethodA(env, obj, methodID, 0);
	}
}
jboolean JNIGO_CallBooleanMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallBooleanMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallBooleanMethodA(env, obj, methodID, 0);
	}
}
jbyte JNIGO_CallByteMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallByteMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallByteMethodA(env, obj, methodID, 0);
	}
}
jchar JNIGO_CallCharMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallCharMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallCharMethodA(env, obj, methodID, 0);
	}
}
jshort JNIGO_CallShortMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallShortMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallShortMethodA(env, obj, methodID, 0);
	}
}
jint JNIGO_CallIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallIntMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallIntMethodA(env, obj, methodID, 0);
	}
}
jlong JNIGO_CallLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallLongMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallLongMethodA(env, obj, methodID, 0);
	}
}
jfloat JNIGO_CallFloatMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallFloatMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallFloatMethodA(env, obj, methodID, 0);
	}
}
jdouble JNIGO_CallDoubleMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallDoubleMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallDoubleMethodA(env, obj, methodID, 0);
	}
}
void JNIGO_CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		(*env)->CallVoidMethodA(env, obj, methodID, vals);
	} else {
		(*env)->CallVoidMethodA(env, obj, methodID, 0);
	}
}
//
jobject JNIGO_CallStaticObjectMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallStaticObjectMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallStaticObjectMethodA(env, obj, methodID, 0);
	}
}
jboolean JNIGO_CallStaticBooleanMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallStaticBooleanMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallStaticBooleanMethodA(env, obj, methodID, 0);
	}
}
jbyte JNIGO_CallStaticByteMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallStaticByteMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallStaticByteMethodA(env, obj, methodID, 0);
	}
}
jchar JNIGO_CallStaticCharMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallStaticCharMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallStaticCharMethodA(env, obj, methodID, 0);
	}
}
jshort JNIGO_CallStaticShortMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallStaticShortMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallStaticShortMethodA(env, obj, methodID, 0);
	}
}
jint JNIGO_CallStaticIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallStaticIntMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallStaticIntMethodA(env, obj, methodID, 0);
	}
}
jlong JNIGO_CallStaticLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallStaticLongMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallStaticLongMethodA(env, obj, methodID, 0);
	}
}
jfloat JNIGO_CallStaticFloatMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallStaticFloatMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallStaticFloatMethodA(env, obj, methodID, 0);
	}
}
jdouble JNIGO_CallStaticDoubleMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		return (*env)->CallStaticDoubleMethodA(env, obj, methodID, vals);
	} else {
		return (*env)->CallStaticDoubleMethodA(env, obj, methodID, 0);
	}
}
void JNIGO_CallStaticVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len) {
	if (len > 0) {
		jvalue vals[len];
		covjval(vals, args, len);
		(*env)->CallStaticVoidMethodA(env, obj, methodID, vals);
	} else {
		(*env)->CallStaticVoidMethodA(env, obj, methodID, 0);
	}
}
//
//
jobjectArray JNIGO_NewObjectArray(JNIEnv *env, jsize len, jclass clazz,
		jobject init) {
	return (*env)->NewObjectArray(env, len, clazz, 0);
}
jbooleanArray JNIGO_NewBooleanArray(JNIEnv *env, jsize len) {
	return (*env)->NewBooleanArray(env, len);
}
jbyteArray JNIGO_NewByteArray(JNIEnv *env, jsize len) {
	return (*env)->NewByteArray(env, len);
}
jcharArray JNIGO_NewCharArray(JNIEnv *env, jsize len) {
	return (*env)->NewCharArray(env, len);
}
jshortArray JNIGO_NewShortArray(JNIEnv *env, jsize len) {
	return (*env)->NewShortArray(env, len);
}
jintArray JNIGO_NewIntArray(JNIEnv *env, jsize len) {
	return (*env)->NewIntArray(env, len);
}
jlongArray JNIGO_NewLongArray(JNIEnv *env, jsize len) {
	return (*env)->NewLongArray(env, len);
}
jfloatArray JNIGO_NewFloatArray(JNIEnv *env, jsize len) {
	return (*env)->NewFloatArray(env, len);
}
jdoubleArray JNIGO_NewDoubleArray(JNIEnv *env, jsize len) {
	return (*env)->NewDoubleArray(env, len);
}
void JNIGO_SetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index,
		jobject val) {
	(*env)->SetObjectArrayElement(env, array, index, val);
}
void JNIGO_SetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start,
		jsize l, const jboolean *buf) {
	(*env)->SetBooleanArrayRegion(env, array, start, l, buf);
}
void JNIGO_SetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start,
		jsize len, const jbyte *buf) {
	(*env)->SetByteArrayRegion(env, array, start, len, buf);
}
void JNIGO_SetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start,
		jsize len, const jchar *buf) {
	(*env)->SetCharArrayRegion(env, array, start, len, buf);
}
void JNIGO_SetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start,
		jsize len, const jshort *buf) {
	(*env)->SetShortArrayRegion(env, array, start, len, buf);
}
void JNIGO_SetIntArrayRegion(JNIEnv *env, jintArray array, jsize start,
		jsize len, const jint *buf) {
	(*env)->SetIntArrayRegion(env, array, start, len, buf);
}
void JNIGO_SetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start,
		jsize len, const jlong *buf) {
	(*env)->SetLongArrayRegion(env, array, start, len, buf);
}
void JNIGO_SetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start,
		jsize len, const jfloat *buf) {
	(*env)->SetFloatArrayRegion(env, array, start, len, buf);
}
void JNIGO_SetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start,
		jsize len, const jdouble *buf) {
	(*env)->SetDoubleArrayRegion(env, array, start, len, buf);
}
//
jobject JNIGO_GetObjectArrayElement(JNIEnv *env, jobjectArray array,
		jsize index) {
	return (*env)->GetObjectArrayElement(env, array, index);
}
void JNIGO_GetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start,
		jsize len, jboolean *buf) {
	(*env)->GetBooleanArrayRegion(env, array, start, len, buf);
}
void JNIGO_GetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start,
		jsize len, jbyte *buf) {
	(*env)->GetByteArrayRegion(env, array, start, len, buf);
}
void JNIGO_GetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start,
		jsize len, jchar *buf) {
	(*env)->GetCharArrayRegion(env, array, start, len, buf);
}
void JNIGO_GetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start,
		jsize len, jshort *buf) {
	(*env)->GetShortArrayRegion(env, array, start, len, buf);
}
void JNIGO_GetIntArrayRegion(JNIEnv *env, jintArray array, jsize start,
		jsize len, jint *buf) {
	(*env)->GetIntArrayRegion(env, array, start, len, buf);
}
void JNIGO_GetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start,
		jsize len, jlong *buf) {
	(*env)->GetLongArrayRegion(env, array, start, len, buf);
}
void JNIGO_GetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start,
		jsize len, jfloat *buf) {
	(*env)->GetFloatArrayRegion(env, array, start, len, buf);
}
void JNIGO_GetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start,
		jsize len, jdouble *buf) {
	(*env)->GetDoubleArrayRegion(env, array, start, len, buf);
}
jsize JNIGO_GetArrayLength(JNIEnv *env, jarray array) {
	return (*env)->GetArrayLength(env, array);
}
//// void show(int a, int...) {
//
//// }
//void abc(jvalue val) {
//	printf("%d\n", val.i);
//}
//int sss(void) {
//	JavaVMOption options[1];
//	JNIEnv *env;
//	JavaVM *jvm;
//	JavaVMInitArgs vm_args;
//	long status;
//	jclass cls;
//	jmethodID mid;
//
//	char* class_name = "jni/CJava";
//	options[0].optionString = "-Djava.class.path=kjava.jar";
//	vm_args.version = JNI_VERSION_1_6;
//	vm_args.nOptions = 1;
//	vm_args.options = options;
//	vm_args.ignoreUnrecognized = JNI_TRUE;
//
//	/* Create the Java VM */
//	status = JNI_CreateJavaVM(&jvm, (void**) &env, &vm_args);
//	if (status < 0 || status == JNI_ERR) {
//		printf("Status of creating Java VM:%ld\n", status);
//		fprintf(stderr, "Failed to create Java VM!\n");
//		exit(1);
//	}
//
//	/*Find the class obj*/
//	cls = (*env)->FindClass(env, class_name);
//	if (cls != 0) {
//		/* ———————————————————- */
//		/*Test to invoke a static method*/
//
//		char* static_method_name = "add";
//		char* static_method_sign = "(II)I";
//
//		/*Get method ID*/
//		mid = (*env)->GetStaticMethodID(env, cls, static_method_name,
//				static_method_sign);
//		if (mid != 0) {
//			/*invoke static int method*/
//			jobject result = (*env)->CallObjectMethod(env, cls, mid, 5, 1);
//			jint abb = result;
//			printf("Call static method %s: %d\n", static_method_name, abb);
//		} else {
//			printf("Failed to find method %s!\n", static_method_name);
//		}
//		char* instance_method_name = "sub";
//		char* instance_method_sign = "(II)I";
//
//		jmethodID constructor_mid = (*env)->GetMethodID(env, cls, "<init>",
//				"()V"); //The construcotr method name is <init>
//		jobject jobj = (*env)->NewObject(env, cls, constructor_mid);
//
//		/*Get method ID*/
//		mid = (*env)->GetMethodID(env, cls, instance_method_name,
//				instance_method_sign);
//		if (mid != 0) {
//			/*invoke int method*/
//			jint result = (*env)->CallIntMethod(env, jobj, mid, 9, 4);
//			printf("Call instance method %s: %d\n", instance_method_name,
//					result);
//		} else {
//			printf("Failed to find method %s!\n", instance_method_name);
//		}
//	} else {
//		printf("Failed to find Class %s!\n", class_name);
//	}
//
//	/*destory JVM*/
//	(*jvm)->DestroyJavaVM(jvm);
//	return 0;
//}
