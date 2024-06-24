package uk.co.amrik.rolodex.order;

import org.glassfish.jersey.server.ResourceConfig;

public class OrderApplication extends ResourceConfig {

    public OrderApplication(){
        packages("uk.co.amrik.steel.order");
    }

}
