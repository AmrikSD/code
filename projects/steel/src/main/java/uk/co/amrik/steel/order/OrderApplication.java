package uk.co.amrik.steel.order;

import org.glassfish.jersey.server.ResourceConfig;
import uk.co.amrik.steel.order.api.v1.OrderService;

public class OrderApplication extends ResourceConfig {

    public OrderApplication(){
        register(OrderService.class);
    }

}
