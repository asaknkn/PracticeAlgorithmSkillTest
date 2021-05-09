import java.util.*;

public class Main {

	public static void main(String[] args) {
		Scanner sc=new Scanner(System.in);
    int N = sc.nextInt();
    
    int h[] = new int[N];
    for (int i = 0; i < N; i++) {
      h[i] = sc.nextInt();
    }

    int cost[] = new int[N];
    cost[0] = 0;
    cost[1] = cost[0] + Math.abs(h[0]-h[1]);

    for (int i = 2; i < N; i++) {
      cost[i] = Math.min(cost[i-1]+Math.abs(h[i-1]-h[i]), cost[i-2]+Math.abs(h[i-2]-h[i]));
    }

    System.out.println(cost[N-1]);
  }
}