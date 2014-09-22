package jnigo;

public class Ary {

	public static void showis(int[] is) {
		System.out.println(is);
	}

	public static void showas(A[] as) {
		System.out.println(as.length);
	}

	public void show(A[] as) {
		for (A a : as) {
			a.show();
		}
	}
}
