import java.util.*;
import java.awt.Point;

public class Main {

	public static void main(String[] args) {
		Scanner sc=new Scanner(System.in);
    int N = sc.nextInt();
    ArrayList<Integer>[] child = new ArrayList[N];
    
    for (int i = 0; i < N; i++) {
      child[i] = new ArrayList<Integer>();
    }

    //child[0].add(0);
    for (int i = 1; i < N; i++) {
      int b = sc.nextInt();
      child[b-1].add(i);
    }
    System.out.println(dfs(0,child));
  }

  public static int dfs(int i, ArrayList<Integer>[] child) {
    if (child[i].size() == 0) {
      return 1;
    } else {
      ArrayList<Integer> values = new ArrayList<Integer>();
      for (int j : child[i]) {
        values.add(dfs(j,child));
      }
      return Collections.max(values) + Collections.min(values) + 1;
    }
  }
}