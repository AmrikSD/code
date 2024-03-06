package de.amrik;

import static org.junit.Assert.assertEquals;

import org.junit.Test;
import static de.amrik.LongestSubarrayOfOnesAfterDeletingOneElement.getLengthOfSubarray;

public class LongestSubarrayOfOnesAfterDeletingOneElementTest {

    @Test
    public void justOneZero() {
        int[] input = { 1, 1, 0, 1 };
        int expected = 3;
        assertEquals(expected, getLengthOfSubarray(input));
    }

    @Test
    public void multipleZeros() {
        int[] input = { 0, 1, 1, 1, 0, 1, 1, 0, 1 };
        int expected = 5;
        assertEquals(expected, getLengthOfSubarray(input));
    }

    @Test
    public void allOnes() {
        int[] input = { 1, 1, 1 };
        int expected = 2; // Sadly, we MUST delete some element.
        assertEquals(expected, getLengthOfSubarray(input));
    }

    @Test
    public void twoZerosNextToEachOtherAndOneAlone() {
        int[] input = { 1, 1, 0, 0, 1, 1, 1, 0, 1 };
        int expected = 4;
        assertEquals(expected, getLengthOfSubarray(input));
    }

    @Test
    public void allZeros() {
        int[] input = { 0, 0, 0 };
        int expected = 0;
        assertEquals(expected, getLengthOfSubarray(input));
    }
}