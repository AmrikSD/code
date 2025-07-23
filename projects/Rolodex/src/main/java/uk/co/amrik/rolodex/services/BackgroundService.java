package uk.co.amrik.rolodex.services;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public abstract class BackgroundService {

    private final Logger logger = LoggerFactory.getLogger(this.getClass());
    private final ExecutorService executor = Executors.newSingleThreadExecutor(r -> {
                Thread t = new Thread(r);
                t.setDaemon(true);
                t.setName(this.getClass().getSimpleName()); // Uses actual implementation class name
                return t;
            }
    );
    private volatile boolean running = true;

    public void start() {
        Runtime.getRuntime().addShutdownHook(new Thread(this::stop));

        executor.submit(() -> {
            while (running && !Thread.currentThread().isInterrupted()) {
                try {
                    workLoop();
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                } catch (Exception e) {
                    logger.error("Unhandled exception in service loop", e);
                }
            }
            logger.info("shutting down");
        });
    }

    public void stop() {
        running = false;
        executor.shutdown();
        try {
            if (!executor.awaitTermination(5, TimeUnit.SECONDS)) {
                logger.warn("Took too long to shutdown, forcing shutdown.");
                executor.shutdownNow(); // Force if needed
            }
        } catch (InterruptedException e) {
            executor.shutdownNow();
            Thread.currentThread().interrupt();
        }
    }

    /** Your background loop goes here. Should return quickly. */
    protected abstract void workLoop() throws Exception;
}
