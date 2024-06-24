package uk.co.amrik.rolodex.user;

import com.google.inject.AbstractModule;

public class UserApplicationModule extends AbstractModule {

    @Override
    protected void configure() {
        bind(UserApplication.class);
    }
}
