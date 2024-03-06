package de.amrik;

/**
 * IDEA: loop over each node: Recursively visit the neighbours of the node, set
 * a visited bit to somewhere if it's a 0 return if it's a 1 then continue
 */
public class NumberOfIslands {
    static boolean[][] visited;

    public static int numIslands(char[][] grid) {
        int islands = 0;
        visited = new boolean[grid.length][grid[0].length];

        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (!visited[i][j]) {
                    if (grid[i][j] == '1') {
                        islands++;
                    }
                    visitElement(i, j, grid);
                }
            }
        }

        return islands;
    }

    private static void visitElement(int x, int y, char[][] grid) {

        if (visited[x][y]) {
            return;
        }

        visited[x][y] = true;

        if (grid[x][y] == '0') {
            return;
        }
        if (x > 0) {
            visitElement(x - 1, y, grid);
        }

        if (x < grid.length - 1) {
            visitElement(x + 1, y, grid);
        }

        if (y > 0) {
            visitElement(x, y - 1, grid);
        }

        if (y < grid[x].length - 1) {
            visitElement(x, y + 1, grid);
        }

    }

}
