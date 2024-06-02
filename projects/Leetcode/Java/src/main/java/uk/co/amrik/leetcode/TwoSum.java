package uk.co.amrik.leetcode;

import java.util.HashMap;

public class TwoSum {

    /**
     *
     * Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
     *
     * You may assume that each input would have exactly one solution, and you may not use the same element twice.
     * You can return the answer in any order.
     **/
    public static int[] solve(int[] arr, int target){
        HashMap<Integer, Integer>  hm = new HashMap<>();
        int[] answer = {};
        for (int i = 0; i < arr.length; i++) {
            if (hm.containsKey(target - arr[i])) {
                answer = new int[]{i, hm.get(target - arr[i])};
            }
            hm.put(arr[i], i);
        }
        return answer;
    }

}
