import java.util.*;
import java.awt.Point;

public class Main {
	public static void main(String[] args) {
		Scanner sc=new Scanner(System.in);
    int N = sc.nextInt();
    int M = sc.nextInt();
    
    char A[][] = new char[N][M];
    for (int i = 0; i < N; i++) {
      String ss = sc.next();
      for (int j = 0; j < M; j++) {
        A[i][j] = ss.charAt(j);
      }
    }

    //int group[][] = new int[11][];
    ArrayList<Point>[] group = new ArrayList[11];
    for (int i = 0; i < 11; i++) {
      group[i] = new ArrayList<Point>();
    }

    for (int i = 0; i < N; i++) {
      for (int j = 0; j < M; j++) {
        int n;
        if (A[i][j] == 'S') {
          n = 0;
        } else if (A[i][j] == 'G') {
          n = 10;
        } else {
          n = Character.getNumericValue(A[i][j]);
        }
        //System.out.println(n);
        group[n].add(new Point(i,j));
      }
    }

    long cost[][] = new long[N][M];
    int INF = Integer.MAX_VALUE;
    for (int i = 0; i < N; i++) {
      for (int j = 0; j < M; j++) {
        cost[i][j] = INF;
      }
    }

    int si = group[0].get(0).x;
    int sj = group[0].get(0).y;
    cost[si][sj] = 0;

    for (int n = 1; n < 11; n++) {
      for (Point p: group[n]) {
        int i = p.x;
        int j = p.y;
        for (Point p2: group[n-1]) {
          int i2 = p2.x;
          int j2 = p2.y;
          cost[i][j] = Math.min(cost[i][j], cost[i2][j2] + Math.abs(i-i2) + Math.abs(j-j2));
        }
      }
    }

    int gi = group[10].get(0).x;
    int gj = group[10].get(0).y;
    long ans = cost[gi][gj];
    if (ans == INF) {
      ans = -1;
    }
    
    System.out.println(ans);
  }
}