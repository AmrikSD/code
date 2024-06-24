package uk.co.amrik.rolodex;

import com.google.inject.AbstractModule;
import uk.co.amrik.rolodex.order.OrderApplicationModule;
import uk.co.amrik.rolodex.user.UserApplicationModule;

public class RolodexModule extends AbstractModule {

    @Override
    protected void configure() {
        install(new OrderApplicationModule());
        install(new UserApplicationModule());
    }


}
