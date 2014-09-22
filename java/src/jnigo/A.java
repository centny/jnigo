package jnigo;

public class A {
	private static int icount = 0;
	private int idx = icount++;

	public void show() {
		System.out.println("Class A " + this.idx);
	}

	public void show(Object t) {
		System.out.println(t);
	}

	public void show(Object[] ts) {
		System.out.println(ts);
	}

	public String show(boolean z) {
		return "abbb";
	}

	public String show(boolean z, byte b) {
		return "abbb";
	}

	public String show(boolean z, byte b, char c, short s, int i, long j,
			float f, double d) {
		return "jjj";
	}

	public String show(boolean z, byte b, char c, short s, int i, long j,
			float f, double d, Object l, String str, boolean[] zs, byte[] bs,
			char[] cs, short[] ss, int[] is, long[] js, float[] fs,
			double[] ds, Object[] ls, String[] strs) {
		return "" + z + b + c + s + i + j + f + d + l + str + zs + bs
				+ new String(cs) + ss + is + js + fs + ds + ls + strs;
	}

	@Override
	public String toString() {
		return super.toString();
	}

	public String ts() {
		return "Tssss";
	}

	public static String tss() {
		return "Tasss";
	}
}
