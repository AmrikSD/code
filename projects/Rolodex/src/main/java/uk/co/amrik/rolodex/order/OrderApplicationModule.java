package uk.co.amrik.rolodex.order;

import com.google.inject.AbstractModule;

public class OrderApplicationModule extends AbstractModule {

    @Override
    protected void configure() {
        bind(OrderApplication.class);
    }
}
