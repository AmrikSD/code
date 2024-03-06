package de.amrik;

import org.junit.Test;
import static org.junit.Assert.assertTrue;

import java.util.Arrays;

import static de.amrik.TwoSum.twoSum;

public class TwoSumTest {

    @Test
    public void sumOfFirstTwoIndicesEqualsTarget() {
        int[] nums = new int[] { 2, 7, 11, 15 };
        int target = 9;
        int[] expectedAnswer = new int[] { 0, 1 };
        int[] reverseExpectedAnswer = new int[] { 0, 1 };
        assertTrue(Arrays.equals(twoSum(nums, target), expectedAnswer)
                || Arrays.equals(twoSum(nums, target), reverseExpectedAnswer));
    }

    @Test
    public void sumOfSecondAndThirdIndicesEqualsTarger() {
        int[] nums = new int[] { 3, 2, 4 };
        int target = 6;
        int[] expectedAnswer = new int[] { 1, 2 };
        int[] reverseExpectedAnswer = new int[] { 1, 2 };
        assertTrue(Arrays.equals(twoSum(nums, target), expectedAnswer)
                || Arrays.equals(twoSum(nums, target), reverseExpectedAnswer));
    }

    @Test
    public void sameNumberTwice() {
        int[] nums = new int[] { 3, 3 };
        int target = 6;
        int[] expectedAnswer = new int[] { 0, 1 };
        int[] reverseExpectedAnswer = new int[] { 0, 1 };
        assertTrue(Arrays.equals(twoSum(nums, target), expectedAnswer)
                || Arrays.equals(twoSum(nums, target), reverseExpectedAnswer));
    }
}
