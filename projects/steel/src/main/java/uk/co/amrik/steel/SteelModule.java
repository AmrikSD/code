package uk.co.amrik.steel;

import com.google.inject.AbstractModule;
import uk.co.amrik.steel.order.OrderModule;
import uk.co.amrik.steel.permissions.PermissionsModule;
import uk.co.amrik.steel.persistence.PersistenceModule;
import uk.co.amrik.steel.user.UserModule;

public class SteelModule extends AbstractModule {

    @Override
    protected void configure() {
        install(new PermissionsModule());
        install(new PersistenceModule());
        install(new OrderModule());
        install(new UserModule());
        bind(Server.class);
    }


}
