package uk.co.amrik.steel.order;

import com.google.inject.AbstractModule;
import uk.co.amrik.steel.persistence.PersistenceModule;

public class TestModule extends AbstractModule {
    @Override
    public void configure(){
        install(new PersistenceModule());
        install(new OrderModule());
    }
}
