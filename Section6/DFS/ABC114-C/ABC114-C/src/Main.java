import java.util.*;

public class Main {
  private static long N;
  private static int ans = 0;
	public static void main(String[] args) {
    Scanner sc=new Scanner(System.in);
    N = Integer.parseInt(sc.next());
    dfs(0, false, false, false);
    System.out.println(ans);
  }

  private static void dfs(long n, Boolean use3, Boolean use5, Boolean use7) {
    if (n > N) {
      return;
    }

    if (use3 && use5 && use7) {
      ans++;
    }

    dfs(10*n+3,true,use5,use7);
    dfs(10*n+5,use3,true,use7);
    dfs(10*n+7,use3,use5,true);
  }
}