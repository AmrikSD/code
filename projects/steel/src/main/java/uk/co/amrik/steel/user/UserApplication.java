package uk.co.amrik.steel.user;

import org.glassfish.jersey.server.ResourceConfig;
import uk.co.amrik.steel.user.api.v1.UserService;

public class UserApplication extends ResourceConfig {

    public UserApplication(){
        register(UserService.class);
    }

}
