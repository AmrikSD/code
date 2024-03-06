package de.amrik;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

import static de.amrik.MaximumNumberOfBalloons.maximumNumberOfBalloons;

public class MaximumNumberOfBalloonsTest {

    @Test
    public void balloonLettersAppearEachOnce() {
        String input = "nlaebolko";
        int expected = 1;
        assertEquals(expected, maximumNumberOfBalloons(input));
    }

    @Test
    public void balloonLettersAppearEachTwiceOnceInOrder() {
        String input = "loonbalxballpoon";
        int expected = 2;
        assertEquals(expected, maximumNumberOfBalloons(input));
    }

    @Test
    public void notAllBaloonLettersAppear() {
        String input = "leetcode";
        int expected = 0;
        assertEquals(expected, maximumNumberOfBalloons(input));
    }
}
