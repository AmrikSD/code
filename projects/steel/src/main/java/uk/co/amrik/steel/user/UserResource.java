package uk.co.amrik.steel.user;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;

import java.util.List;

@Path("/")
public interface UserResource {
    @GET
    @Produces(MediaType.APPLICATION_JSON)
    List<User> user();
}
