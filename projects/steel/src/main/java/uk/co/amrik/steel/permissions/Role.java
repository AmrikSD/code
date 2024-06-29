package uk.co.amrik.steel.permissions;

import java.util.List;

public enum Role {
    USER_WRITE(
            List.of(Permission.USER_WRITE,  Permission.USER_READ)
    ),
    USER_READ(List.of(Permission.USER_WRITE));

    public final List<Permission> permissions;

    Role(List<Permission> permissions){
        this.permissions = permissions;
    }
}
