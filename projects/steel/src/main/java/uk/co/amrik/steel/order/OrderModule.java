package uk.co.amrik.steel.order;

import com.google.inject.AbstractModule;
import uk.co.amrik.steel.order.api.v1.OrderResource;
import uk.co.amrik.steel.order.api.v1.OrderService;

public class OrderModule extends AbstractModule {

    @Override
    protected void configure() {
        bind(OrderResource.class).to(OrderService.class);
    }
}
