package de.amrik;

import java.util.Arrays;

public class LongestSubarrayOfOnesAfterDeletingOneElement {

    /**
     * Given a binary array nums, you should delete one element from it.
     * 
     * @return the size of the longest non-empty subarray containing only 1's in the
     *         resulting array. Return 0 if there is no such subarray.
     */
    public static int getLengthOfSubarray(int[] nums) {

        // If we have all 0s in the array, return 0.
        if (Arrays.stream(nums).sum() == 0) {
            return 0;
        }
        if (Arrays.stream(nums).sum() >= nums.length - 1) {
            return nums.length - 1;
        }
        int maxSeen = 0;
        int left = 0;
        int right = 0;

        for (int i = 0; i < nums.length; i++) {
            left += nums[i];

            if (nums[i] == 0) {
                for (int j = i + 1; j < nums.length; j++) {
                    right += nums[j];
                    if (nums[j] == 0 || j == nums.length - 1) {
                        i = j - 1;
                        maxSeen = Math.max(maxSeen, left + right);
                        left = right;
                        right = 0;
                        break;
                    }
                }
            }

        }

        return maxSeen;
    }
}
