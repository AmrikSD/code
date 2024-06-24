package uk.co.amrik.steel.user;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;

import java.util.List;

@Path("/user")
public class UserService implements UserResource {
    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public List<User> user(){
        return List.of(
                new User("hi")
        );
    }
}
