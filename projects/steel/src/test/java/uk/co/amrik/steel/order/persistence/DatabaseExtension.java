package uk.co.amrik.steel.order.persistence;

import com.google.inject.Guice;
import com.google.inject.Inject;
import com.google.inject.Injector;
import org.junit.jupiter.api.extension.AfterEachCallback;
import org.junit.jupiter.api.extension.BeforeAllCallback;
import org.junit.jupiter.api.extension.BeforeEachCallback;
import org.junit.jupiter.api.extension.ExtensionContext;
import org.testcontainers.containers.wait.strategy.Wait;
import uk.co.amrik.steel.persistence.Migrations;

import org.testcontainers.containers.PostgreSQLContainer;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;


public class DatabaseExtension implements BeforeAllCallback, BeforeEachCallback, AfterEachCallback {

    private static PostgreSQLContainer<?> postgresqlContainer;

    @Inject
    private Migrations migrations;

    @Override
    public void beforeAll(ExtensionContext extensionContext) {
        if (postgresqlContainer == null) {
            postgresqlContainer = new PostgreSQLContainer<>("postgres:latest")
                    .withUsername("postgres")
                    .withPassword("postgres")
                    .withDatabaseName("steel")
                    .waitingFor(Wait.forLogMessage(".*database system is ready to accept connections.*\\n", 1));

            postgresqlContainer.start();

            System.setProperty("STEEL_DB_HOST", postgresqlContainer.getHost());
            System.setProperty("STEEL_DB_PORT", String.valueOf(postgresqlContainer.getFirstMappedPort()));
            System.setProperty("STEEL_DB_NAME", postgresqlContainer.getDatabaseName());
            System.setProperty("STEEL_DB_USER", postgresqlContainer.getUsername());
            System.setProperty("STEEL_DB_PASS", postgresqlContainer.getPassword());

            waitForDatabaseConnection(postgresqlContainer, 5, 1000);
        }

    }

    @Override
    public void beforeEach(ExtensionContext extensionContext) throws Exception {
        Injector injector = Guice.createInjector(binder -> binder.bind(Migrations.class));
        injector.injectMembers(this);
        if (postgresqlContainer != null) {
            migrations.runMigrations();
        }
    }

    @Override
    public void afterEach(ExtensionContext extensionContext) {
        if (postgresqlContainer != null) {
            String jdbcUrl = postgresqlContainer.getJdbcUrl();
            String username = postgresqlContainer.getUsername();
            String password = postgresqlContainer.getPassword();
            try (Connection conn = DriverManager.getConnection(jdbcUrl, username, password)) {
                if (conn != null) {
                    Statement stmt = conn.createStatement();
                    // Drop and recreate the public schema to ensure a fresh start
                    stmt.execute("DROP SCHEMA public CASCADE");
                    stmt.execute("CREATE SCHEMA public");
                    stmt.close();
                }
            } catch (SQLException e) {
                throw new RuntimeException("Interrupted while waiting for database connection", e);
            }
        }
    }

    private void waitForDatabaseConnection(PostgreSQLContainer<?> container, int maxAttempts, long delayBetweenAttempts) {
        String jdbcUrl = container.getJdbcUrl();
        String username = container.getUsername();
        String password = container.getPassword();

        for (int attempt = 1; attempt <= maxAttempts; attempt++) {
            try (Connection conn = DriverManager.getConnection(jdbcUrl, username, password)) {
                if (conn != null) {
                    System.out.println("Successfully connected to the database.");
                    return;
                }
            } catch (SQLException e) {
                try {
                    Thread.sleep(delayBetweenAttempts);
                } catch (InterruptedException ie) {
                    Thread.currentThread().interrupt();
                    throw new RuntimeException("Interrupted while waiting for database connection", ie);
                }
            }
        }
        throw new RuntimeException("Failed to connect to database after " + maxAttempts + " attempts.");
    }

}
