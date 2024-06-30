package uk.co.amrik.steel.user;

import com.google.inject.AbstractModule;
import uk.co.amrik.steel.permissions.PermissionsModule;
import uk.co.amrik.steel.user.api.v1.UserService;

public class UserModule extends AbstractModule {

    @Override
    protected void configure() {
        install(new PermissionsModule());
        bind(UserService.class);
    }
}
