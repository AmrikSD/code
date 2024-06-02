package uk.co.amrik.leetcode;

import static org.assertj.core.api.Assertions.*;

import org.junit.Test;

public class TwoSumTest{
    @Test
    public void exampleOne(){
        int[] nums = {2, 7, 11, 15};
        int target = 9;
        int[] expected = {0, 1};

        assertThat(TwoSum.solve(nums, target))
            .hasSize(2)
            .containsExactlyInAnyOrder(expected);
    }
    @Test
    public void exampleTwo(){
        int[] nums = {3, 2, 4};
        int target = 6;
        int[] expected = {1, 2};

        assertThat(TwoSum.solve(nums, target))
            .hasSize(2)
            .containsExactlyInAnyOrder(expected);
    }
    @Test
    public void exampleThree(){
        int[] nums = {3, 3};
        int target = 6;
        int[] expected = {0, 1};
        assertThat(TwoSum.solve(nums, target))
            .hasSize(2)
            .containsExactlyInAnyOrder(expected);
    }
}
