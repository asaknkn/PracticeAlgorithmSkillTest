import java.util.*;

public class Main {

	public static void main(String[] args) {
		Scanner sc=new Scanner(System.in);
    int N = sc.nextInt();

    int ps[] = new int[N+1];
    ps[0] = 0;
    
    int P = 0;
    for (int i = 0; i < N; i++) {
      ps[i+1] = sc.nextInt();
      P += ps[i+1];
    }

    boolean exist[][] = new boolean[N+1][P+1];
    for (int i = 0; i < N+1; i++) {
      for (int s = 0; s < P+1; s++) {
        exist[i][s] = false;
      }
    }
    exist[0][0] = true;

    for (int i = 1; i < N+1; i++) {
      for (int s = 0; s < P+1; s++) {
        if (exist[i-1][s]) {
          exist[i][s] = true;
        }
        if (s >= ps[i] && exist[i-1][s-ps[i]]) {
          exist[i][s] = true;
        }
      }
    }

    int ans = 0;
    for (int s = 0; s < P+1; s++) {
      if (exist[N][s]) {
        ans++;
      }
    }
    System.out.println(ans);
  }
}