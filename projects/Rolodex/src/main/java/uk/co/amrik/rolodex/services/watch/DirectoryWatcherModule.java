package uk.co.amrik.rolodex.services.watch;

import com.google.inject.AbstractModule;
import com.google.inject.Provides;
import com.google.inject.Singleton;
import com.typesafe.config.Config;
import uk.co.amrik.rolodex.config.ConfigModule;

public class DirectoryWatcherModule extends AbstractModule {

    @Provides
    @Singleton
    public DirectoryWatcherConfig provideDirectoryWatcherConfig(Config config) {
        DirectoryWatcherConfig watcherConfig = new DirectoryWatcherConfig(config);
        return watcherConfig;
    }

    @Override
    protected void configure() {
        install(new ConfigModule());
    }


}
