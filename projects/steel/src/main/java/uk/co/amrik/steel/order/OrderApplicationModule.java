package uk.co.amrik.steel.order;

import com.google.inject.AbstractModule;

public class OrderApplicationModule extends AbstractModule {

    @Override
    protected void configure() {
        bind(OrderApplication.class);
    }
}
