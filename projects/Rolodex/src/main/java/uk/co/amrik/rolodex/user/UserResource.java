package uk.co.amrik.rolodex.user;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;

@Path("/user")
public class UserResource {
    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public User user(){
        return new User("user");
    }
    }
