package uk.co.amrik.steel;

import com.google.inject.Guice;
import com.google.inject.Injector;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.servlet.DefaultServlet;
import org.eclipse.jetty.servlet.ServletContextHandler;
import org.eclipse.jetty.servlet.ServletHolder;
import org.glassfish.jersey.server.ResourceConfig;
import org.glassfish.jersey.servlet.ServletContainer;
import uk.co.amrik.steel.order.OrderApplication;
import uk.co.amrik.steel.user.UserApplication;

public class Steel {

    public static void main(String ...args) throws Exception {

        Injector injector = Guice.createInjector(new SteelModule());

        OrderApplication orderApplication = injector.getInstance(OrderApplication.class);
        UserApplication userApplication = injector.getInstance(UserApplication.class);


        Server server = new Server(8080);
        ServletContextHandler servletContextHandler = new ServletContextHandler(ServletContextHandler.NO_SESSIONS);
        servletContextHandler.setContextPath("/");
        servletContextHandler.addServlet(DefaultServlet.class, "/");

        addServletForApplication(servletContextHandler, orderApplication, "/api/order/*");
        addServletForApplication(servletContextHandler, userApplication, "/api/another/*");

        server.setHandler(servletContextHandler);
        server.start();
        server.join();

        System.out.println("Hello");
    }

    private static void addServletForApplication(ServletContextHandler context, ResourceConfig app, String pathSpec) {
        ServletHolder sh = context.addServlet(ServletContainer.class, pathSpec);
        sh.setInitParameter("jakarta.ws.rs.Application", app.getClass().getName());
    }

}
