package uk.co.amrik.steel.persistence;

import com.google.inject.Inject;
import org.flywaydb.core.Flyway;
import org.flywaydb.core.api.FlywayException;


public class Migrations {

    private final DatabaseConfig databaseConfig;

    @Inject
    public Migrations(
            DatabaseConfig databaseConfig
    ){
        this.databaseConfig = databaseConfig;
    }

    public void runMigrations() throws Exception {
        String url = "jdbc:postgresql://%s:%s/%s".formatted(
                databaseConfig.getHost(),
                databaseConfig.getPort(),
                databaseConfig.getName()
        );
        Flyway flyway = Flyway.configure()
                .dataSource(
                        url,
                        databaseConfig.getUser(),
                        databaseConfig.getPass()
                )
                .baselineOnMigrate(true)
                .locations("classpath:db/migration")
                .load();

        // Start the migration
        try {
            flyway.migrate();
        } catch (FlywayException e) {
            throw new Exception(e);
        }

    }

}
