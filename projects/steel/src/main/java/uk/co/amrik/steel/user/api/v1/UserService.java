package uk.co.amrik.steel.user.api.v1;

import uk.co.amrik.steel.user.model.User;
import java.util.List;

public class UserService implements UserResource {

    public List<User> users(){
        return List.of(
                new User("hi")
        );
    }

    public List<User> more(){
        return List.of(
                new User("hi"),
                new User("hi"),
                new User("hi"),
                new User("hi"),
                new User("hi")
        );
    }
}
