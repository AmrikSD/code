package uk.co.amrik.rolodex.database;

import com.google.inject.Inject;
import org.flywaydb.core.Flyway;
import org.flywaydb.core.api.FlywayException;
import org.flywaydb.core.api.MigrationInfo;
import org.flywaydb.core.api.MigrationInfoService;
import org.flywaydb.core.api.exception.FlywayValidateException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.sql.DataSource;

public class MigrationService {

    private static final Logger logger = LoggerFactory.getLogger(MigrationService.class);
    private final DataSource dataSource;
    private final DatabaseConfig config;

    @Inject
    public MigrationService(DataSource dataSource, DatabaseConfig config) {
        this.dataSource = dataSource;
        this.config = config;
    }

    public void runMigrations() {
        logger.info("Starting database migrations...");

        try {
            Flyway flyway = Flyway.configure()
                    .dataSource(dataSource)
                    .locations("classpath:db/migration")
                    .baselineOnMigrate(true)
                    .validateOnMigrate(true)
                    .load();

            // Get migration info before running
            MigrationInfoService infoService = flyway.info();
            MigrationInfo[] pendingMigrations = infoService.pending();

            if (pendingMigrations.length == 0) {
                logger.info("No pending database migrations found");
            } else {
                logger.info("Found {} pending migration(s)", pendingMigrations.length);
                for (MigrationInfo migration : pendingMigrations) {
                    logger.info("  - {}: {}", migration.getVersion(), migration.getDescription());
                }
            }

            // Run migrations
            int migrationsExecuted = flyway.migrate().migrationsExecuted;

            if (migrationsExecuted > 0) {
                logger.info("Successfully executed {} database migration(s)", migrationsExecuted);
            } else {
                logger.info("Database is up to date, no migrations executed");
            }

        } catch (FlywayException e) {
            logger.error("Failed to run database migrations", e);
        }
    }

}
