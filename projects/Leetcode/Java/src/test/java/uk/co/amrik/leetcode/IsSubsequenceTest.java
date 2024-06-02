package uk.co.amrik.leetcode;

import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.Assert.assertEquals;

public class IsSubsequenceTest {

    @Test
    public void exampleOne(){
        String wordOne = "abc";
        String wordTwo = "ahbgdc";
        Boolean expected = true;
        assertThat(IsSubsequence.solve(wordOne, wordTwo)).isEqualTo(expected);
    }

    @Test
    public void exampleTwo(){
        String wordOne = "axc";
        String wordTwo = "ahbgdc";
        Boolean expected = false;
        assertThat(IsSubsequence.solve(wordOne, wordTwo)).isEqualTo(expected);
    }

    @Test
    public void exampleThree(){
        String wordOne = "abc";
        String wordTwo = "ahbgdc";
        Boolean expected = true;
        assertThat(IsSubsequence.solve(wordOne, wordTwo)).isEqualTo(expected);
    }
    @Test
    public void exampleFour(){
        String wordOne = "";
        String wordTwo = "ahbgdc";
        Boolean expected = true;
        assertThat(IsSubsequence.solve(wordOne, wordTwo)).isEqualTo(expected);
    }

}
