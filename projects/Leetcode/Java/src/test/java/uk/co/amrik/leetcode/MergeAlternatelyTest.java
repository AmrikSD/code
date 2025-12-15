package uk.co.amrik.leetcode;

import org.junit.jupiter.api.Test;
import org.junit.platform.runner.JUnitPlatform;
import org.junit.runner.RunWith;

import static org.assertj.core.api.Assertions.assertThat;

@RunWith(JUnitPlatform.class)
public class MergeAlternatelyTest {
    @Test
    public void exampleOne(){
        String wordOne = "abc";
        String wordTwo = "pqr";
        String expected = "apbqcr"; // Base case
        assertThat(MergeAlternately.solve(wordOne, wordTwo)).isEqualTo(expected);
    }

    @Test
    public void exampleTwo(){
        String wordOne = "ab";
        String wordTwo = "pqrs";
        String expected = "apbqrs";
        assertThat(MergeAlternately.solve(wordOne, wordTwo)).isEqualTo(expected);
    }

    @Test
    public void exampleThree(){
        String wordOne = "abcd";
        String wordTwo = "pq";
        String expected = "apbqcd";
        assertThat(MergeAlternately.solve(wordOne, wordTwo)).isEqualTo(expected);
    }
}
