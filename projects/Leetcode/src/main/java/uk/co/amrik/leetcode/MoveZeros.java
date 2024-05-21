package uk.co.amrik.leetcode;

public class MoveZeros {

    public static int[] Solve(int[] input){
        int left = 0;
        int right = 0;
        for (int i = 0; i < input.length; i++) {
            if (input[right] != 0){
                int tmp = input[left];
                input[left] = input[right];
                input[right] = tmp;
                left++;
            }
            right++;
        }

        return input;
    }
}
