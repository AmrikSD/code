package de.amrik.DataStructures;

public class TreeNode {

    public int val;
    public TreeNode left;
    public TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }

    public static TreeNode fromArray(Integer[] arr) {
        if (arr.length == 0) {
            return null;
        }
        TreeNode root = new TreeNode(arr[0]);
        return fromArray(arr, root, 0);
    }

    static TreeNode fromArray(Integer[] arr, TreeNode root, int depth) {
        if (depth < arr.length) {
            if (arr[depth] == null) {
                return null;
            }
            TreeNode temp = new TreeNode(arr[depth]);
            root = temp;
            root.left = fromArray(arr, root.left, depth * 2 + 1);
            root.right = fromArray(arr, root.right, depth * 2 + 2);
        }
        return root;
    }

}
