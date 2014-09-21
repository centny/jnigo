package jnigo;

public class ObjectM {
	private int a, b;

	public ObjectM(int a, int b) {
		this.a = a;
		this.b = b;
	}

	public int getA() {
		return a;
	}

	public void setA(int a) {
		this.a = a;
	}

	public int getB() {
		return b;
	}

	public void setB(int b) {
		this.b = b;
	}

	// //////

	public void showz(boolean a) {
		System.out.println("A z:" + a);
	}

	public void showb(byte a) {
		System.out.println("A b:" + a);
	}

	public void showc(char a) {
		System.out.println("A c:" + a);
	}

	public void shows(short a) {
		System.out.println("A s:" + a);
	}

	public void showi(int a) {
		System.out.println("A i:" + a);
	}

	public void showj(long a) {
		System.out.println("A j:" + a);
	}

	public void showf(float a) {
		System.out.println("A f:" + a);
	}

	public void showd(double a) {
		System.out.println("A d:" + a);
	}

	public void getv() {

	}

	public boolean getz() {
		return false;
	}

	public byte getb() {
		return 1;
	}

	public char getc() {
		return 1;
	}

	public short gets() {
		return 1;
	}

	public int geti() {
		return 1;
	}

	public long getj() {
		return 1;
	}

	public float getf() {
		return 1;
	}

	public double getd() {
		return 1;
	}

}
