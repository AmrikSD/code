package uk.co.amrik.steel.middleware.filters;

import jakarta.ws.rs.container.ContainerRequestContext;
import jakarta.ws.rs.container.ContainerResponseContext;
import jakarta.ws.rs.container.ContainerResponseFilter;
import jakarta.ws.rs.core.NewCookie;

import java.io.IOException;
import java.util.Optional;

public class LoggingMiddleware implements ContainerResponseFilter {

    @Override
    public void filter(ContainerRequestContext requestContext, ContainerResponseContext responseContext) throws IOException {

        System.out.println(requestContext.getMethod() + " /" + requestContext.getUriInfo().getPath());
        // Create a new cookie
        NewCookie cookie = new NewCookie("myCookie", "cookieValue");

        // Add the cookie to the response headers
        responseContext.getHeaders().add("Set-Cookie", cookie.toString());
    }
}
