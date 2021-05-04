import java.util.*;
import java.awt.Point;
import java.util.ArrayDeque;

public class Main {
  private static final int[] xdirection = {0, 1, 0, -1};
  private static final int[] ydirection = {1, 0, -1, 0};
	public static void main(String[] args) {
		Scanner sc=new Scanner(System.in);
    int r = sc.nextInt();
    int c = sc.nextInt();
    int sy = sc.nextInt() -1;
    int sx = sc.nextInt() -1;
    int gy = sc.nextInt() -1;
    int gx = sc.nextInt() -1;

    char s[][] = new char[r][c];
    for (int i = 0; i < r; i++) {
      String ss = sc.next();
      for (int j = 0; j < c; j++) {
        s[i][j] = ss.charAt(j);
      }
    }

    int dist[][] = new int[r][c];
    for (int i = 0; i < r; i++) {
      for (int j = 0; j < c; j++) {
        dist[i][j] = -1;
      }
    }

    ArrayDeque<Point> q = new ArrayDeque<>();
    q.add(new Point(sy,sx));
    dist[sy][sx] = 0;

    while(!q.isEmpty()) {
      Point p = q.removeFirst();

      for (int i = 0; i < 4; i++) {
        int i2 = p.x + xdirection[i];
        int j2 = p.y + ydirection[i];

        if ((0 > i2 || i2 >= r) && (0 > j2 || j2 >= c)) {
          continue;
        }

        if (s[p.x][p.y] == '#') {
          continue;
        }

        if (dist[i2][j2] == -1) {
          dist[i2][j2] = dist[p.x][p.y] + 1;
          q.add(new Point(i2,j2));
        }
      }
    }
    System.out.println(dist[gy][gx]);
  }
}