package uk.co.amrik.steel.order;

import org.glassfish.jersey.server.ResourceConfig;

public class OrderApplication extends ResourceConfig {

    public OrderApplication(){
        packages("uk.co.amrik.steel.order");
    }

}
