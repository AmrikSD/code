package uk.co.amrik.steel.permissions;

import com.google.inject.AbstractModule;

public class PermissionsModule extends AbstractModule {
    @Override
    protected void configure() {
        bind(PermissionsService.class);
    }
}
