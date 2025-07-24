package uk.co.amrik.rolodex;

import com.google.inject.Guice;
import com.google.inject.Injector;
import com.google.inject.servlet.GuiceFilter;
import org.eclipse.jetty.ee10.servlet.ServletContextHandler;
import org.eclipse.jetty.server.Server;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import uk.co.amrik.rolodex.database.MigrationService;
import uk.co.amrik.rolodex.services.BackgroundServiceManager;

public class Rolodex {

    private static final Logger logger = LoggerFactory.getLogger(Rolodex.class);
    private static final int PORT = 8080;

    public static void main(String ...args) throws Exception {

        Injector workerInjector = Guice.createInjector(
                new WorkerModule()
        );

        MigrationService migrationService = workerInjector.getInstance(MigrationService.class);
        migrationService.runMigrations();

        BackgroundServiceManager manager = workerInjector.getInstance(BackgroundServiceManager.class);
        manager.startAll();

        Server server = new Server(PORT);

        ServletContextHandler context = new ServletContextHandler(ServletContextHandler.SESSIONS);
        context.setContextPath("/");
        server.setHandler(context);

        context.addFilter(GuiceFilter.class, "/*", null);
        context.addEventListener(new RolodexServletConfig());

        logger.info("Starting Rolodex server on port {}", PORT);
        server.start();
        server.join();
    }

}
