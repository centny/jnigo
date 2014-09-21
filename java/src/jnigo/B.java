package jnigo;

public class B {
	private A a;

	public B(A a) {
		this.a = a;
	}

	public void setA(A a) {
		this.a = a;
	}

	public A getA() {
		return this.a;
	}

	public void show() {
		System.out.println("Class B");
	}
}
