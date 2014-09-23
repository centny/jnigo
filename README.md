JNI-GO
======
> library for calling java by golang through JNI

===
#Install
* install jdk (install mingw with gcc on window)
* go get:

linux/unix/osx:

 ```
export CGO_CFLAGS="-I$JAVA_HOME/include/ -I$JAVA_HOME/include/<darwin/linux>"

export CGO_LDFLAGS="-I$JAVA_HOME/jre/lib/server -ljvm"

go get github.com/Centny/jnigo	
```

win32:

```
set CGO_CFLAGS="-I%JAVA_HOME%/include/ -I%JAVA_HOME%/include/win32"

set CGO_LDFLAGS="-I%JAVA_HOME%/lib -ljvm"

go get github.com/Centny/jnigo	
```


#Example

```go

func ExampleJvm() {
	Init("-Djava.class.path=java/bin")
	str_a, err := GVM.New("java.lang.String")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(str_a.CallInt("length"))
	//
	str_b, err := GVM.New("java.lang.String", "abc")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(str_b.CallInt("length"))
	fmt.Println(str_b.AsString())
	fmt.Println("---->")
}

func ExampleCovertType() {
	clsa := GVM.FindClass("Ljnigo/A;")
	obja, _ := clsa.New()
	objv, _ := obja.As("Ljava/lang/Object;")
	objs, _ := GVM.NewAry("Ljava/lang/Object;", 1)
	objs.SetObject(0, objv)
	fmt.Println(obja.CallVoid("show", objv))
	fmt.Println(obja.CallVoid("show", objs))
	show, err := obja.CallObject("show", "Ljava/lang/String;",
		true,              //for java boolean
		Byte(1),           //for java byte
		Char(1),           //for java char
		int16(1),          //for java short
		1,                 //for java int
		int64(1),          //for java long
		float32(1),        //for java float
		float64(1),        //for java double
		objv,              //for java Object
		"jjjjj",           //for java String
		[]bool{false},     //for java boolean[]
		[]Byte{1, 2},      //for java byte[]
		[]Char{3, 4},      //for java char[]
		[]int16{11, 12},   //for java short[]
		[]int{21, 22},     //for java int[]
		[]int64{31, 32},   //for java long[]
		[]float32{41, 42}, //for java float[]
		[]float64{51, 52}, //for java double[]
		objs,              //for java Object[]
		[]string{"aaa"},   //for java String[]
	)
	if err != nil {
		t.Error(err.Error())
		return
	}
}
```