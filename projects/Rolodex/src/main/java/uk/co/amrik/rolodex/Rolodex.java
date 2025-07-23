package uk.co.amrik.rolodex;

import com.google.inject.servlet.GuiceFilter;
import org.eclipse.jetty.ee10.servlet.ServletContextHandler;
import org.eclipse.jetty.server.Server;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Rolodex {

    private static final Logger logger = LoggerFactory.getLogger(Rolodex.class);
    private static final int PORT = 8080;

    public static void main(String ...args) throws Exception {
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
