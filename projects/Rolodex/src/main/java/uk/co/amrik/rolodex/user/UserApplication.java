package uk.co.amrik.rolodex.user;

import org.glassfish.jersey.server.ResourceConfig;

public class UserApplication extends ResourceConfig {

    public UserApplication(){
        packages("uk.co.amrik.steel.user");
    }

}
