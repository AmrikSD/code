package uk.co.amrik.steel.API;

import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;

import java.util.List;
import java.util.Optional;

@Path("/")
public interface Api<Request, Response> {

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    List<Response> getAll();

    @GET
    @Path("/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    Optional<Response> get(@PathParam("id") Integer id);

    @DELETE
    @Path("/{id}")
    int delete(@PathParam("id") Integer id);

    @PUT
    @Produces(MediaType.APPLICATION_JSON)
    Optional<Response> put(Request entity);

}
