package uk.co.amrik.steel.user.api.v1;

import uk.co.amrik.steel.user.model.User;
import java.util.List;

public class UserService implements UserResource {

    public List<User> user(){
        return List.of(
                new User("hi")
        );
    }
}
