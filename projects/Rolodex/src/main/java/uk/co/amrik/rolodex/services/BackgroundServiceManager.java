package uk.co.amrik.rolodex.services;

import com.google.inject.Inject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Set;

public class BackgroundServiceManager {

    private static final Logger logger = LoggerFactory.getLogger(BackgroundServiceManager.class);
    private final Set<BackgroundService> services;

    @Inject
    public BackgroundServiceManager(Set<BackgroundService> services){
        this.services = services;
    }

    public void startAll() {
        logger.info("Starting background services...");
        services.forEach(BackgroundService::start);
    }

    public void stopAll() {
        logger.info("Stopping background services...");
        services.forEach(BackgroundService::stop);
    }
}
