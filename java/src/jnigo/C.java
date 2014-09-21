package jnigo;

public class C {
	private A[] as;

	public A[] getAs() {
		return as;
	}

	public void setAs(A[] as) {
		this.as = as;
	}

	public void setAss(int[] ss) {

	}

	public void showas() {
		for (A a : as) {
			a.show();
		}
	}
}
