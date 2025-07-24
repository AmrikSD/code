package uk.co.amrik.rolodex;

import com.google.inject.AbstractModule;
import com.google.inject.multibindings.Multibinder;
import uk.co.amrik.rolodex.services.BackgroundService;
import uk.co.amrik.rolodex.services.watch.DirectoryWatcherModule;
import uk.co.amrik.rolodex.services.watch.DirectoryWatcherService;

public class WorkerModule extends AbstractModule {

    @Override
    protected void configure(){

        install(new DirectoryWatcherModule());          // Provides DirectoryWatcherConfig

        Multibinder<BackgroundService> serviceBinder = Multibinder.newSetBinder(binder(), BackgroundService.class);
        serviceBinder.addBinding().to(DirectoryWatcherService.class).asEagerSingleton();
    }
}
