package uk.co.amrik.steel.middleware.filters;

import jakarta.ws.rs.container.ContainerRequestContext;
import jakarta.ws.rs.container.ContainerResponseContext;
import jakarta.ws.rs.container.ContainerResponseFilter;
import jakarta.ws.rs.core.NewCookie;
import jakarta.ws.rs.core.Response;

import java.io.IOException;
import java.util.Optional;

public class NotFoundFilter implements ContainerResponseFilter {

    @Override
    public void filter(ContainerRequestContext requestContext, ContainerResponseContext responseContext) throws IOException {
        Object entity = responseContext.getEntity();

        if (entity == null){
            responseContext.setStatusInfo(Response.Status.NOT_FOUND);
            responseContext.setEntity("Resource not found");
        }

        if (entity instanceof Optional) {
            Optional<?> optional = (Optional<?>) entity;
            if (optional.isEmpty()) {
                responseContext.setStatusInfo(Response.Status.NOT_FOUND);
                responseContext.setEntity("Resource not found");
            }
        }
    }
}
