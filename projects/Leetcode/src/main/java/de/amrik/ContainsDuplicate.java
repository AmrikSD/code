package de.amrik;

import java.util.HashSet;
import java.util.Set;

public class ContainsDuplicate {

    public static boolean containsDuplicate(int[] arr) {
        Set<Integer> hs = new HashSet<>(arr.length);
        for (int i : arr) {
            if (hs.contains(i)) {
                return true;
            }
            hs.add(i);
        }
        return false;
    }

}
