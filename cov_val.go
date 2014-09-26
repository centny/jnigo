package jnigo

import (
	"reflect"
	"unsafe"
)

type Byte byte
type Char byte

func NewAry(sig string, ptr unsafe.Pointer, l int) (string, Object, error) {
	obj, err := JNIGO_newAry(sig, l)
	if err != nil {
		return "", obj, err
	}
	if l > 0 {
		JNIGO_sary(obj, ptr, 0, l)
	}
	return obj.Sig(), obj, nil
}
func NewAryv(sig string, ptr unsafe.Pointer, l int) (string, Val, error) {
	sig, obj, err := NewAry(sig, ptr, l)
	if err == nil {
		return sig, obj.Val(), nil
	} else {
		return "", Val{}, err
	}
}
func CovArgs(args ...interface{}) (string, []Val, error) {
	vals := []Val{}
	sigs := ""
	for _, arg := range args {
		sig, val, err := covval(arg)
		if err != nil {
			return "", nil, err
		} else {
			sigs += sig
			vals = append(vals, val)
		}
	}
	return sigs, vals, nil
}
func covval(arg interface{}) (string, Val, error) {
	var _bval_ Byte
	var _cval_ Char
	var _oval_ Object
	var __oval__ *Object
	ptype := reflect.TypeOf(arg)
	switch ptype {
	case reflect.TypeOf(_bval_):
		return "B", Val{b: byte(arg.(Byte)), typ: 'Z'}, nil
	case reflect.TypeOf(_cval_):
		return "C", Val{c: byte(arg.(Char)), typ: 'C'}, nil
	case reflect.TypeOf(_oval_):
		obj := arg.(Object)
		return obj.Sig(), obj.Val(), nil
	case reflect.TypeOf(__oval__):
		obj := arg.(*Object)
		return obj.Sig(), obj.Val(), nil
	}
	switch ptype.Kind() {
	case reflect.Bool:
		return "Z", Val{z: arg.(bool), typ: 'Z'}, nil
	case reflect.Uint8:
		return "B", Val{b: byte(arg.(uint8)), typ: 'B'}, nil
	case reflect.Int16:
		return "S", Val{s: arg.(int16), typ: 'S'}, nil
	case reflect.Int32:
		return "I", Val{i: int(arg.(int32)), typ: 'I'}, nil
	case reflect.Int:
		return "I", Val{i: arg.(int), typ: 'I'}, nil
	case reflect.Int64:
		return "J", Val{j: arg.(int64), typ: 'J'}, nil
	case reflect.Float32:
		return "F", Val{f: arg.(float32), typ: 'F'}, nil
	case reflect.Float64:
		return "D", Val{d: arg.(float64), typ: 'D'}, nil
	case reflect.Slice:
		return covary(arg)
	case reflect.String:
		return covstr(arg.(string))
	default:
		return "", Val{}, Err("invalid type:%s", ptype.Kind().String())
	}
}

func covary(arg interface{}) (string, Val, error) {
	if arg == nil {
		return "", Val{}, Err("arg is nil")
	}
	pval := reflect.ValueOf(arg)
	ptype := reflect.TypeOf(arg)
	if ptype.Kind() != reflect.Slice {
		return "", Val{}, Err("not slice for:%v", arg)
	}
	var _bval_ Byte
	var _cval_ Char
	switch ptype.Elem() {
	case reflect.TypeOf(_bval_):
		l := pval.Len()
		if l < 1 {
			return NewAryv("B", nil, 0)
		}
		vals := []jbyte{}
		for i := 0; i < l; i++ {
			vals = append(vals, jbyte(pval.Index(i).Interface().(Byte)))
		}
		return NewAryv("B", unsafe.Pointer(&vals[0]), l)
	case reflect.TypeOf(_cval_):
		l := pval.Len()
		if l < 1 {
			return NewAryv("C", nil, l)
		}
		vals := []jchar{}
		for i := 0; i < l; i++ {
			vals = append(vals, jchar(pval.Index(i).Interface().(Char)))
		}
		return NewAryv("C", unsafe.Pointer(&vals[0]), l)
	}
	//
	switch ptype.Elem().Kind() {
	case reflect.Bool:
		vals := []jboolean{}
		for _, b := range arg.([]bool) {
			if b {
				vals = append(vals, jboolean(JNI_TRUE))
			} else {
				vals = append(vals, jboolean(JNI_FALSE))
			}
		}
		l := len(vals)
		if l < 1 {
			return NewAryv("Z", nil, l)
		} else {
			return NewAryv("Z", unsafe.Pointer(&vals[0]), l)
		}
	case reflect.Uint8:
		vals := []jbyte{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, jbyte(pval.Index(i).Interface().(uint8)))
		}
		l := len(vals)
		if l < 1 {
			return NewAryv("B", nil, l)
		} else {
			return NewAryv("B", unsafe.Pointer(&vals[0]), l)
		}
	case reflect.Int16:
		vals := []jshort{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, jshort(pval.Index(i).Interface().(int16)))
		}
		l := len(vals)
		if l < 1 {
			return NewAryv("S", nil, l)
		} else {
			return NewAryv("S", unsafe.Pointer(&vals[0]), l)
		}
	case reflect.Int32:
		vals := []jint{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, jint(pval.Index(i).Interface().(int32)))
		}
		l := len(vals)
		if l < 1 {
			return NewAryv("I", nil, l)
		} else {
			return NewAryv("I", unsafe.Pointer(&vals[0]), l)
		}
	case reflect.Int:
		vals := []jint{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, jint(pval.Index(i).Interface().(int)))
		}
		l := len(vals)
		if l < 1 {
			return NewAryv("I", nil, l)
		} else {
			return NewAryv("I", unsafe.Pointer(&vals[0]), l)
		}
	case reflect.Int64:
		vals := []jlong{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, jlong(pval.Index(i).Interface().(int64)))
		}
		l := len(vals)
		if l < 1 {
			return NewAryv("J", nil, l)
		} else {
			return NewAryv("J", unsafe.Pointer(&vals[0]), l)
		}
	case reflect.Float32:
		vals := []jfloat{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, jfloat(pval.Index(i).Interface().(float32)))
		}
		l := len(vals)
		if l < 1 {
			return NewAryv("F", nil, l)
		} else {
			return NewAryv("F", unsafe.Pointer(&vals[0]), l)
		}
	case reflect.Float64:
		vals := []jdouble{}
		for i := 0; i < pval.Len(); i++ {
			vals = append(vals, jdouble(pval.Index(i).Interface().(float64)))
		}
		l := len(vals)
		if l < 1 {
			return NewAryv("D", nil, l)
		} else {
			return NewAryv("D", unsafe.Pointer(&vals[0]), l)
		}
	case reflect.String:
		l := pval.Len()
		obj, _ := JNIGO_newAry("Ljava/lang/String;", l)
		for i := 0; i < l; i++ {
			str := JNIGO_newS(pval.Index(i).Interface().(string))
			obj.Set(i, str.Val())
		}
		return obj.Sig(), obj.Val(), obj.Err()
	default:
		return "", Val{}, Err("invalid type:%s", ptype.Elem().Kind().String())
	}
}

func covstr(arg string) (string, Val, error) {
	obj := JNIGO_newS(arg)
	return "Ljava/lang/String;", obj.Val(), nil
}
