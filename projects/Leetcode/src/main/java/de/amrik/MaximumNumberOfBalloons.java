package de.amrik;

import java.util.HashMap;

/**
 * https://leetcode.com/problems/maximum-number-of-balloons/
 */
public class MaximumNumberOfBalloons {
    /**
     * @param text 1< text.length <= 10^4. Only consists of lower case English
     *             letters.
     * @return the maximum number of times the word "balloon" can be created from
     *         the characters that make up the string {@link text}
     */
    public static int maximumNumberOfBalloons(String text) {
        HashMap<Character, Integer> hm = new HashMap<Character, Integer>();
        String balloon = "balloon";

        for (char c : text.toCharArray()) {
            if (hm.containsKey(c)) {
                hm.put(c, hm.get(c) + 1);
            } else {
                hm.put(c, 1);
            }
        }

        int balloons = 0;
        while (true) {
            for (char c : balloon.toCharArray()) {
                if (!hm.containsKey(c)) {
                    return balloons;
                }
                if (hm.get(c) < 1) {
                    return balloons;
                }
                hm.put(c, hm.get(c) - 1);
            }
            balloons++;
        }

    }
}
