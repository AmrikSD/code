package de.amrik;

import de.amrik.DataStructures.TreeNode;

import static org.junit.Assert.assertEquals;
import static de.amrik.MaximumDepthOfBinaryTree.getMaxDepth;
import org.junit.Test;

public class MaximumDepthOfBinaryTreeTest {

    @Test
    public void DepthOfThreeWithTwoLeafsMissingFromOneChild() {
        Integer[] arr = { 3, 9, 20, null, null, 15, 7 };
        TreeNode root = TreeNode.fromArray(arr);
        int expected = 3;
        assertEquals(expected, getMaxDepth(root));
    }

    @Test
    public void nodeWithOneChild() {
        Integer[] arr = { 1, null, 2 };
        TreeNode root = TreeNode.fromArray(arr);
        int expected = 2;
        assertEquals(expected, getMaxDepth(root));
    }

    @Test
    public void emptyArray() {
        Integer[] arr = {};
        TreeNode root = TreeNode.fromArray(arr);
        int expected = 0;
        assertEquals(expected, getMaxDepth(root));
    }

    @Test
    public void singleNode() {
        Integer[] arr = { 0 };
        TreeNode root = TreeNode.fromArray(arr);
        int expected = 1;
        assertEquals(expected, getMaxDepth(root));
    }
}
