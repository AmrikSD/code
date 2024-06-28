package uk.co.amrik.steel.order.api.v1;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;
import uk.co.amrik.steel.order.model.Order;

import java.util.List;

@Path("/")
public interface OrderResource {

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    List<Order> order();

}
