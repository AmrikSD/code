package uk.co.amrik.steel.persistence;

import com.google.inject.Inject;
import com.google.inject.Singleton;
import org.jooq.DSLContext;
import org.jooq.Field;
import org.jooq.Record;
import org.jooq.SelectSelectStep;
import org.jooq.impl.DSL;

import java.nio.file.attribute.FileTime;

@Singleton
public class DatabaseService {

    private final DSLContext dsl;

    @Inject
    public DatabaseService(
            DatabaseConfig databaseConfig
    ){
        String url = "jdbc:postgresql://%s:%s/%s".formatted(
                databaseConfig.getHost(),
                databaseConfig.getPort(),
                databaseConfig.getName()
        );
        this.dsl = DSL.using(
                url,
                databaseConfig.getUser(),
                databaseConfig.getPass()
        );
    }

    public DSLContext getDsl() {
        return dsl;
    }
}
