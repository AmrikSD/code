package uk.co.amrik.steel;

import com.google.inject.Inject;
import org.eclipse.jetty.servlet.DefaultServlet;
import org.eclipse.jetty.servlet.ServletContextHandler;
import org.eclipse.jetty.servlet.ServletHolder;
import org.glassfish.jersey.server.ResourceConfig;
import org.glassfish.jersey.servlet.ServletContainer;
import uk.co.amrik.steel.middleware.filters.LoggingMiddleware;
import uk.co.amrik.steel.middleware.filters.NotFoundFilter;
import uk.co.amrik.steel.order.api.v1.OrderService;
import uk.co.amrik.steel.permissions.PermissionsService;
import uk.co.amrik.steel.persistence.Migrations;
import uk.co.amrik.steel.user.api.v1.UserService;

public class Server {

    private final Migrations migrations;
    private final PermissionsService permissionsService;
    private final UserService userService;
    private final OrderService orderService;

    @Inject
    public Server(
            Migrations migrations,
            PermissionsService permissionsService,
            UserService userService,
            OrderService orderService
    ){
        this.migrations = migrations;
        this.permissionsService = permissionsService;
        this.userService = userService;
        this.orderService = orderService;
    }

    public void run() throws Exception {
        migrations.runMigrations();

        permissionsService.bootstrapRoles();

        org.eclipse.jetty.server.Server server = new org.eclipse.jetty.server.Server(8080);
        ServletContextHandler servletContextHandler = new ServletContextHandler(ServletContextHandler.NO_SESSIONS);
        servletContextHandler.setDefaultContextPath("/api");
        servletContextHandler.addServlet(DefaultServlet.class, "/");

        addServletForApplication(servletContextHandler, orderService, "/v1/order");
        addServletForApplication(servletContextHandler, userService, "/v1/user");

        server.setHandler(servletContextHandler);
        server.start();
        server.join();


    }

    private static void addServletForApplication(ServletContextHandler context, Object resource, String pathSpec) {
        ServletHolder servletHolder = new ServletHolder(new ServletContainer(
                new ResourceConfig()
                        .register(resource)
                        .register(LoggingMiddleware.class)
        ));
        context.addServlet(servletHolder, pathSpec + "/*");
    }

}
