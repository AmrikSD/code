package uk.co.amrik.steel.user.model;

import uk.co.amrik.steel.permissions.Role;

import java.util.List;

public class User {
    private final Integer id;
    private final String name;

    public User(
            Integer id,
            String name
    ){
        this.id = id;
        this.name = name;
    }

    public Integer getId() {
        return id;
    }
    public String getName() {
        return name;
    }

}
