package uk.co.amrik.leetcode;

public class MergeAlternately {

    /**
     * You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
     * If a string is longer than the other, append the additional letters onto the end of the merged string.
     *
     * Return the merged string.
     */
    public static String solve(String wordOne, String wordTwo){
        StringBuilder sb = new StringBuilder();
        int maxLength = Math.max(wordOne.length(), wordTwo.length());
        for (int i = 0; i < maxLength; i++) {
            if (i < wordOne.length()) {
                sb.append(wordOne.charAt(i));
            }
            if (i < wordTwo.length()) {
                sb.append(wordTwo.charAt(i));
            }
        }
        return sb.toString();
    }
}
