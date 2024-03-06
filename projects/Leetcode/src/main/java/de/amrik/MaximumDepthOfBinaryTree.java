package de.amrik;

import de.amrik.DataStructures.TreeNode;

public class MaximumDepthOfBinaryTree {

    public static int getMaxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        if (root.left == null && root.right == null) {
            return 1;
        }
        return 1 + Math.max(getMaxDepth(root.left), getMaxDepth(root.right));
    }

}
