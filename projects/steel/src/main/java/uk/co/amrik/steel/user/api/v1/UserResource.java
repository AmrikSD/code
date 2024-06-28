package uk.co.amrik.steel.user.api.v1;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;
import uk.co.amrik.steel.user.model.User;

import java.util.List;

@Path("/")
public interface UserResource {

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    List<User> user();
}
