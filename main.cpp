#ifdef __DEV__
#include <iostream>
#include "jnigo.h"
#include "jnigo_c.h"

using namespace std;
void jnigo_c() {
	JNIGO_addOption("-Djava.class.path=src/java/bin");
	cout << JNIGO_init() << endl;
	struct c_class cls = JNIGO_findClass("Ljnigo/M;");
	if (!cls.vmo.valid) {
		cout << cls.vmo.msg << endl;
		JNIGO_destory();
		return;
	}
	cout << "000" << endl;
	struct c_method ms_v = JNIGO_findClsMethod(cls, "ms_v", "", "Ljava/lang/String;",
			true);
	if (!ms_v.vmo.valid) {
		cout << ms_v.vmo.msg << endl;
		JNIGO_destory();
		return;
	}
	cout << "001" << endl;
	struct c_object msv_o = JNIGO_callA(ms_v, 0, 0);
	if (!msv_o.vmo.valid) {
		cout << msv_o.vmo.msg << endl;
		JNIGO_destory();
		return;
	}
	cout << "010" << endl;
	struct c_method newm = JNIGO_findClsMethod(cls, "<init>", "", "V", false);
	if (!newm.vmo.valid) {
		cout << newm.vmo.msg << endl;
		return;
	}
	cout << "020" << endl;
	struct c_object obj = JNIGO_newA(newm, 0, 0);
	if (!obj.vmo.valid) {
		cout << obj.vmo.msg << endl;
		return;
	}
	cout << "030" << endl;
	struct c_method m_v = JNIGO_findObjMethod(obj, "m_v", "", "I");
	if (!m_v.vmo.valid) {
		cout << m_v.vmo.msg << endl;
		JNIGO_destory();
		return;
	}
	cout << "040" << endl;
	struct c_object mv_o = JNIGO_callA(m_v, 0, 0);
	if (!mv_o.vmo.valid) {
		cout << mv_o.vmo.msg << endl;
		JNIGO_destory();
		return;
	}
	cout << mv_o.val.typ << "  " << mv_o.val.i << endl;
	cout << "050" << endl;
	struct c_method m_v2 = JNIGO_findObjMethod(obj, "m_v2", "IJ", "I");
	if (!m_v2.vmo.valid) {
		cout << m_v2.vmo.msg << endl;
		JNIGO_destory();
		return;
	}
	cout << "060" << endl;
	struct jval vals[2];
	vals[0].i = 100;
	vals[0].typ = 'I';
	vals[1].j = 20000;
	vals[1].typ = 'J';
	struct c_object mv_o2 = JNIGO_callA(m_v2, vals, 2);
	if (!mv_o2.vmo.valid) {
		cout << mv_o2.vmo.msg << endl;
		JNIGO_destory();
		return;
	}
	cout << mv_o2.val.typ << "  " << mv_o2.val.i << endl;
	cout << "070" << endl;
	JNIGO_destory();
}
int main() {
	jnigo_c();
	return 0;
}

#endif
