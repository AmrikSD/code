package uk.co.amrik.leetcode;

import java.util.ArrayDeque;
import java.util.Queue;

public class IsSubsequence {
    public static Boolean solve(String wordOne, String wordTwo){

        Queue<Integer> queue = new ArrayDeque<>();
        wordOne.chars().forEach(queue::add);

        if (queue.isEmpty()) {
            return true;
        }
        wordTwo.chars().forEach(c -> {
            if (queue.isEmpty()) {
                return;
            }
            if (queue.peek() == c) {
                queue.remove();
            }
        });

        return queue.isEmpty();
    }
}
