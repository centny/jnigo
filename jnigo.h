#ifndef JNIGO_H_
#define JNIGO_H_
//

#include <jni.h>

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
	int typ;
} jval;
//
jclass JNIGO_FindClass(JNIEnv *env, const char *name);
jint JNIGO_DestroyJavaVM(JavaVM *vm);
jobject JNIGO_NewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID,
		const jval *args, int len);
jthrowable JNIGO_ExceptionOccurred(JNIEnv *env);
void JNIGO_ExceptionClear(JNIEnv *env);
//
//
jfieldID JNIGO_GetFieldID(JNIEnv *env, jclass clazz, const char *name,
		const char *sig);
jfieldID JNIGO_GetStaticFieldID(JNIEnv *env, jclass clazz, const char *name,
		const char *sig);
//
jobject JNIGO_GetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID);
jboolean JNIGO_GetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID);
jbyte JNIGO_GetByteField(JNIEnv *env, jobject obj, jfieldID fieldID);
jchar JNIGO_GetCharField(JNIEnv *env, jobject obj, jfieldID fieldID);
jshort JNIGO_GetShortField(JNIEnv *env, jobject obj, jfieldID fieldID);
jint JNIGO_GetIntField(JNIEnv *env, jobject obj, jfieldID fieldID);
jlong JNIGO_GetLongField(JNIEnv *env, jobject obj, jfieldID fieldID);
jfloat JNIGO_GetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID);
jdouble JNIGO_GetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID);
//
jobject JNIGO_GetStaticObjectField(JNIEnv *env, jobject obj, jfieldID fieldID);
jboolean JNIGO_GetStaticBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID);
jbyte JNIGO_GetStaticByteField(JNIEnv *env, jobject obj, jfieldID fieldID);
jchar JNIGO_GetStaticCharField(JNIEnv *env, jobject obj, jfieldID fieldID);
jshort JNIGO_GetStaticShortField(JNIEnv *env, jobject obj, jfieldID fieldID);
jint JNIGO_GetStaticIntField(JNIEnv *env, jobject obj, jfieldID fieldID);
jlong JNIGO_GetStaticLongField(JNIEnv *env, jobject obj, jfieldID fieldID);
jfloat JNIGO_GetStaticFloatField(JNIEnv *env, jobject obj, jfieldID fieldID);
jdouble JNIGO_GetStaticDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID);
//
//
void JNIGO_SetVal(JNIEnv *env, jobject obj, jfieldID fieldID, jval val);
void JNIGO_SetStaticVal(JNIEnv *env, jobject obj, jfieldID fieldID, jval val);
//
//
jmethodID JNIGO_GetMethodID(JNIEnv *env, jclass clazz, const char *name,
		const char *sig);
jmethodID JNIGO_GetStaticMethodID(JNIEnv *env, jclass clazz, const char *name,
		const char *sig);
//
jobject JNIGO_CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jboolean JNIGO_CallBooleanMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jbyte JNIGO_CallByteMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jchar JNIGO_CallCharMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jshort JNIGO_CallShortMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jint JNIGO_CallIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jlong JNIGO_CallLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jfloat JNIGO_CallFloatMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jdouble JNIGO_CallDoubleMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
void JNIGO_CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
//
jobject JNIGO_CallStaticObjectMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len);
jboolean JNIGO_CallStaticBooleanMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len);
jbyte JNIGO_CallStaticByteMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jchar JNIGO_CallStaticCharMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jshort JNIGO_CallStaticShortMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len);
jint JNIGO_CallStaticIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jlong JNIGO_CallStaticLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
jfloat JNIGO_CallStaticFloatMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len);
jdouble JNIGO_CallStaticDoubleMethodA(JNIEnv *env, jobject obj,
		jmethodID methodID, const jval *args, int len);
void JNIGO_CallStaticVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID,
		const jval *args, int len);
//
//
jobjectArray JNIGO_NewObjectArray(JNIEnv *env, jsize len, jclass clazz,
		jobject init);
jbooleanArray JNIGO_NewBooleanArray(JNIEnv *env, jsize len);
jbyteArray JNIGO_NewByteArray(JNIEnv *env, jsize len);
jcharArray JNIGO_NewCharArray(JNIEnv *env, jsize len);
jshortArray JNIGO_NewShortArray(JNIEnv *env, jsize len);
jintArray JNIGO_NewIntArray(JNIEnv *env, jsize len);
jlongArray JNIGO_NewLongArray(JNIEnv *env, jsize len);
jfloatArray JNIGO_NewFloatArray(JNIEnv *env, jsize len);
jdoubleArray JNIGO_NewDoubleArray(JNIEnv *env, jsize len);
//
//
void JNIGO_SetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index,
		jobject val);
void JNIGO_SetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start,
		jsize l, const jboolean *buf);
void JNIGO_SetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start,
		jsize len, const jbyte *buf);
void JNIGO_SetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start,
		jsize len, const jchar *buf);
void JNIGO_SetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start,
		jsize len, const jshort *buf);
void JNIGO_SetIntArrayRegion(JNIEnv *env, jintArray array, jsize start,
		jsize len, const jint *buf);
void JNIGO_SetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start,
		jsize len, const jlong *buf);
void JNIGO_SetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start,
		jsize len, const jfloat *buf);
void JNIGO_SetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start,
		jsize len, const jdouble *buf);
//
jobject JNIGO_GetObjectArrayElement(JNIEnv *env, jobjectArray array,
		jsize index);
void JNIGO_GetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start,
		jsize len, jboolean *buf);
void JNIGO_GetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start,
		jsize len, jbyte *buf);
void JNIGO_GetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start,
		jsize len, jchar *buf);
void JNIGO_GetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start,
		jsize len, jshort *buf);
void JNIGO_GetIntArrayRegion(JNIEnv *env, jintArray array, jsize start,
		jsize len, jint *buf);
void JNIGO_GetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start,
		jsize len, jlong *buf);
void JNIGO_GetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start,
		jsize len, jfloat *buf);
void JNIGO_GetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start,
		jsize len, jdouble *buf);
//
//
jsize JNIGO_GetArrayLength(JNIEnv *env, jarray array);
#endif
