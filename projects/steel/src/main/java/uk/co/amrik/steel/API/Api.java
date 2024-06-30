package uk.co.amrik.steel.API;

import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;

import java.util.List;
import java.util.Optional;

@Path("/")
public interface Api<T> {

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    List<T> getAll();

    @GET
    @Path("/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    Optional<T> get(@PathParam("id") Integer id);

    @DELETE
    @Path("/{id}")
    void delete(@PathParam("id") Integer id);

    @PUT
    @Produces(MediaType.APPLICATION_JSON)
    T put(T entity);

}
