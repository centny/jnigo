package jnigo;

public class Ary {

	public static void showis(int[] is) {
		System.out.println(is);
	}

	public static void showas(A[] as) {
		System.out.println(as.length);
		for (A a : as) {
			System.out.println("---showas--->11");
			a.show();
		}
	}

	public void show(A[] as) {
		System.out.println("------->" + as);
		for (A a : as) {
			System.out.println("------->11");
			a.show();
		}
	}
}
