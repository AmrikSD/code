package uk.co.amrik.rolodex;

import com.google.inject.Scopes;
import com.google.inject.servlet.ServletModule;
import uk.co.amrik.rolodex.health.HealthCheckServlet;

public class RolodexServletModule extends ServletModule {
    @Override
    protected void configureServlets() {
        bind(HealthCheckServlet.class).in(Scopes.SINGLETON);
        serve("/healthcheck").with(HealthCheckServlet.class);
    }
}
