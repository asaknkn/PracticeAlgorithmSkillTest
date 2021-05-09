import java.util.*;

public class Main {

	public static void main(String[] args) {
		Scanner sc=new Scanner(System.in);
    int N = sc.nextInt();
    int W = sc.nextInt();

    int ws[] = new int[N+1];
    long vs[] = new long[N+1];
    ws[0] = 0;
    vs[0] = 0;
    
    for (int i = 0; i < N; i++) {
      ws[i+1] = sc.nextInt();
      vs[i+1] = sc.nextInt();
    }    
    
    long value[][] = new long[N+1][W+1];
    for (int i = 0; i < N+1; i++) {
      for (int w = 0; w < W+1; w++) {
        value[i][w] = 0;
      }
    }

    for (int i = 1; i < N+1; i++) {
      for (int w = 0; w < W+1; w++) {
        value[i][w] = Math.max(value[i][w], value[i-1][w]);

        if (w-ws[i] >= 0) {
          value[i][w] = Math.max(value[i][w], value[i-1][w-ws[i]]+vs[i]);
        }
      }
    }

    long ans = 0;
    for (int w = 0; w < W; w++) {
      ans = Math.max(value[N][w],value[N][w+1]);
    }
    
    System.out.println(ans);
  }
}