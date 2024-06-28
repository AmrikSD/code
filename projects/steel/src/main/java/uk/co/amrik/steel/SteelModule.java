package uk.co.amrik.steel;

import com.google.inject.AbstractModule;
import uk.co.amrik.steel.order.OrderApplicationModule;
import uk.co.amrik.steel.persistence.PersistenceModule;
import uk.co.amrik.steel.user.UserApplicationModule;

public class SteelModule extends AbstractModule {

    @Override
    protected void configure() {
        install(new PersistenceModule());
        install(new OrderApplicationModule());
        install(new UserApplicationModule());
    }


}
