package jnigo;

public class A {
	private static int icount = 0;
	private int idx = icount++;

	public void show() {
		System.out.println("Class A " + this.idx);
	}
}
