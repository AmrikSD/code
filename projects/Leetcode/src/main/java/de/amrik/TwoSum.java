package de.amrik;

import java.util.HashMap;

/**
 * https://leetcode.com/problems/two-sum/
 */
public class TwoSum {

    /***
     * Given an array of integers @param nums and an integer @param target, returns
     * indices of the two numbers such that they add up to target. You may assume
     * that each input would have exactly one solution, and you may not use the same
     * element twice. The answer may be returned in any order.
     * 
     * @param nums
     * @param target
     * @return the INDICES of the two numbers in @param nums , such that they add up
     *         to @param targert
     */
    public static int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> hm = new HashMap<Integer, Integer>();

        for (int i = 0; i < nums.length; i++) {
            hm.put(nums[i], i);
        }

        for (int i = 0; i < nums.length; i++) {
            if (hm.containsKey(target - nums[i])) {
                if (hm.get(target - nums[i]) == i) {
                    continue;
                }
                return new int[] { i, hm.get(target - nums[i]) };
            }
        }

        return new int[] {};

    }

}