import java.util.*;
import java.awt.Point;
import java.util.ArrayDeque;

public class Main {
  private static Scanner sc=new Scanner(System.in);
  private static final int[] xdirection = {0, 1, 0, -1};
  private static final int[] ydirection = {1, 0, -1, 0};
  private static int h = sc.nextInt();;
  private static int w = sc.nextInt();
  private static char s[][] = new char[h][w];
  private static boolean visited[][] = new boolean[h][w];
	public static void main(String[] args) {
    for (int i = 0; i < h; i++) {
      String ss = sc.next();
      for (int j = 0; j < w; j++) {
        s[i][j] = ss.charAt(j);
      }
    }

    int si = 0, sj = 0, gi = 0, gj = 0;
    for (int i = 0; i < h; i++) {
      for (int j = 0; j < w; j++) {
        if (s[i][j] == 's') {
          si = i;
          sj = j;
        }
        if (s[i][j] == 'g') {
          gi = i;
          gj = j;
        }
      }
    }

    for (int i = 0; i < h; i++) {
      for (int j = 0; j < w; j++) {
        visited[i][j] = false;
      }
    }

    dfs(si, sj);

    if (visited[gi][gj]) {
      System.out.println("Yes");
    } else {
      System.out.println("No");
    }
  }

  private static void dfs(int i, int j) {
    visited[i][j] = true;

    for (int k = 0; k < 4; k++) {
      int i2 = i + xdirection[k];
      int j2 = j + ydirection[k];

      if ((0 > i2 || i2 >= h) || (0 > j2 || j2 >= w)) {
        continue;
      }

      if (s[i][j] == '#') {
        continue;
      }

      if (visited[i2][j2] != true) {
        dfs(i2,j2);
      }
    }
  }

}