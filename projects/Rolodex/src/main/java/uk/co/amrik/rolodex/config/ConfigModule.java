package uk.co.amrik.rolodex.config;

import com.google.inject.AbstractModule;
import com.google.inject.Provides;
import com.google.inject.Singleton;
import com.typesafe.config.Config;
import com.typesafe.config.ConfigException;
import com.typesafe.config.ConfigFactory;
import org.slf4j.LoggerFactory;
import org.slf4j.Logger;

public class ConfigModule extends AbstractModule {
    private static final Logger logger = LoggerFactory.getLogger(ConfigModule.class);

    @Provides
    @Singleton
    public Config provideConfig() throws ConfigException {
        try {
            Config config = ConfigFactory.load("application");
            if (!validateConfig(config)){
                throw new ConfigException.Generic("Failed custom validation checks");
            };
            return config;
        }catch(ConfigException e){
            logger.error("Failed to load configuration", e);
            throw e;
        }
    }

    private boolean validateConfig(Config config){
        return true; //TODO: Actually validate lol
    }
}
