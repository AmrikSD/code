package de.amrik;

import static org.junit.Assert.assertEquals;
import static de.amrik.ContainsDuplicate.containsDuplicate;

import org.junit.Test;

public class ContainsDuplicateTest {

    @Test
    public void noDuplicates() {
        int[] nums = { 1, 2, 3, 4 };
        boolean expected = false;
        assertEquals(expected, containsDuplicate(nums));
    }

    @Test
    public void ContainsOneTwice() {
        int[] nums = { 1, 2, 3, 1 };
        boolean expected = true;
        assertEquals(expected, containsDuplicate(nums));
    }

    @Test
    public void ContainsOneTripleAndOneDuplicate() {
        int[] nums = { 1, 2, 3, 1 };
        boolean expected = true;
        assertEquals(expected, containsDuplicate(nums));
    }

}
