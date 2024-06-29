package uk.co.amrik.steel.persistence;

import com.google.inject.AbstractModule;

public class PersistenceModule extends AbstractModule {

    @Override
    protected void configure(){
        bind(DatabaseConfig.class);
        bind(DatabaseService.class);
    }
}
