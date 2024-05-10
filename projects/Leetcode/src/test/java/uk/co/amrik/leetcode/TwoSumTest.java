package uk.co.amrik.leetcode;

import static org.assertj.core.api.Assertions.*;

import org.junit.Test;

public class TwoSumTest{
    @Test
    public void shouldAlwaysPass(){
        int[] answer = TwoSum.Solve();
        assertThat(answer).hasSize(2);
    }

}
