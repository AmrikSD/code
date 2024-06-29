package uk.co.amrik.steel.user.model;

import uk.co.amrik.steel.permissions.Role;

import java.util.List;

public class User {
    private final String name;
    private final List<Role> roles;

    public User(
            String name,
            List<Role> roles
    ){
        this.name = name;
        this.roles = roles;
    }

    public String getName() {
        return name;
    }
    public List<Role> getRoles() {
        return roles;
    }

}
