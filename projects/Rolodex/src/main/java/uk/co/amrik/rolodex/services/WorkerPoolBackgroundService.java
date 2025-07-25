package uk.co.amrik.rolodex.services;

import org.slf4j.LoggerFactory;
import org.slf4j.Logger;

import java.util.concurrent.*;


public abstract class WorkerPoolBackgroundService<T> extends BackgroundService {
    private final Logger logger = LoggerFactory.getLogger(this.getClass());
    private final ExecutorService workerPool;
    private final BlockingQueue<T> workQueue = new LinkedBlockingQueue<>();
    private final int workerCount;

    public WorkerPoolBackgroundService(int workerCount) {
        super();
        this.workerCount = workerCount;
        this.workerPool = Executors.newFixedThreadPool(workerCount, r -> {
            Thread t = new Thread(r);
            t.setDaemon(true);
            t.setName(this.getClass().getSimpleName() + "-worker-" + t.threadId());
            return t;
        });
    }

    @Override
    public void start() {
        // Start worker threads
        for (int i = 0; i < workerCount; i++) {
            final int workerId = i;
            workerPool.submit(() -> {
                logger.info("Worker {} starting", workerId);
                while (!Thread.currentThread().isInterrupted()) {
                    try {
                        T work = workQueue.take(); // Blocks until work available
                        processWork(work);
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        break;
                    } catch (Exception e) {
                        logger.error("Error processing work in worker {}", workerId, e);
                    }
                }
                logger.info("Worker {} shutting down", workerId);
            });
        }

        // Start the main service loop
        super.start();
    }

    @Override
    public void stop() {
        workerPool.shutdownNow();
        try {
            if (!workerPool.awaitTermination(5, TimeUnit.SECONDS)) {
                logger.warn("Worker pool took too long to shutdown");
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        super.stop();
    }

    @Override
    protected final void workLoop() throws Exception {
        // This method discovers work and adds it to the queue
        T work = discoverWork();
        if (work != null) {
            workQueue.offer(work);
        } else {
            // No work found, sleep a bit to avoid busy waiting
            Thread.sleep(1000);
        }
    }

    /**
     * Discover work to be done. Return null if no work is available.
     * This runs in a single thread.
     */
    protected abstract T discoverWork() throws Exception;

    /**
     * Process a unit of work. This runs in the worker thread pool.
     */
    protected abstract void processWork(T work) throws Exception;
}
