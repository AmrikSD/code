package uk.co.amrik.leetcode;

import org.junit.jupiter.api.Test;
import org.junit.platform.runner.JUnitPlatform;
import org.junit.runner.RunWith;
import static org.assertj.core.api.Assertions.assertThat;

@RunWith(JUnitPlatform.class)
public class MoveZerosTest {

    @Test
    public void exampleOne(){
        int[] input = {0,1,0,3,12};
        int[] expected = {1,3,12,0,0};
        int[] actual = MoveZeros.Solve(input);

        assertThat(actual).isEqualTo(expected);
    }

    @Test
    public void exampleTwo(){
        int[] input = {0};
        int[] expected = {0};
        int[] actual = MoveZeros.Solve(input);

        assertThat(actual).isEqualTo(expected);
    }
}
