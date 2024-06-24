package uk.co.amrik.steel.user;

import com.google.inject.AbstractModule;

public class UserApplicationModule extends AbstractModule {

    @Override
    protected void configure() {
        bind(UserApplication.class);
        bind(UserResource.class).to(UserService.class);
    }
}
