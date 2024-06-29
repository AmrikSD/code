package uk.co.amrik.steel.user.api.v1;

import com.google.inject.Inject;
import uk.co.amrik.steel.permissions.PermissionsService;
import uk.co.amrik.steel.permissions.Role;
import uk.co.amrik.steel.user.model.User;
import java.util.List;

public class UserService implements UserResource {

    private final PermissionsService permissionsService;

    @Inject
    public UserService(
            PermissionsService permissionsService
    ){
        this.permissionsService = permissionsService;
    }

    public List<User> users(){
        List<Role> roles = permissionsService.getRoles();
        return List.of(
                new User("hi", roles)
        );
    }

    public List<User> more(){
        return List.of(
                new User("another", List.of()),
                new User("another", List.of()),
                new User("another", List.of()),
                new User("another", List.of()),
                new User("another", List.of())
        );
    }
}
