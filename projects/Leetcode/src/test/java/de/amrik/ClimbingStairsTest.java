package de.amrik;

import static org.junit.Assert.assertEquals;
import static de.amrik.ClimbingStairs.waysToClimbStairs;

import org.junit.Test;

public class ClimbingStairsTest {
    @Test
    public void nequals3() {
        int input = 3;
        int expected = 3;
        assertEquals(expected, waysToClimbStairs(input));
    }

    @Test
    public void nequals2() {
        int input = 2;
        int expected = 2;
        assertEquals(expected, waysToClimbStairs(input));
    }
}
