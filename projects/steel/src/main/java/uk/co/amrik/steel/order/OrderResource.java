package uk.co.amrik.steel.order;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;

@Path("/order")
public class OrderResource {
    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public Order order(){
        return new Order("thingy");
    }
    }
