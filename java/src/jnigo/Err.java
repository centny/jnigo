package jnigo;

public class Err {
	public static void err1() throws Exception {
		System.out.println("err1");
		throw new Exception("err1");
	}

	public static void err2() {
		System.out.println("err2");
		throw new RuntimeException("err2");
	}

	public static void err3() {
		System.out.println("err3");
		throw new RuntimeException("err3");
	}
}
