package jnigo;

public class M {
	public static int fs_i = 0;

	public static String ms_v() {
		return "ms_v";
	}

	public int f_i = 1000;

	public int m_v() {
		return 1000;
	}

	public int m_v2(int a, long b) {
		System.out.println(a + "->" + b);
		return 2000;
	}
}
