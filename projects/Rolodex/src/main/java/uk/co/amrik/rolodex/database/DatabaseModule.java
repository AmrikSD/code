package uk.co.amrik.rolodex.database;

import com.google.inject.AbstractModule;
import com.google.inject.Provides;
import com.google.inject.Singleton;
import com.typesafe.config.Config;
import com.zaxxer.hikari.HikariConfig;
import com.zaxxer.hikari.HikariDataSource;

import javax.sql.DataSource;

public class DatabaseModule extends AbstractModule {


    @Override
    public void configure(){
        bind(MigrationService.class).asEagerSingleton();
    }

    @Provides
    @Singleton
    public DatabaseConfig provideDatabaseConfig(Config config) {
        return new DatabaseConfig(config);
    }

    @Provides
    @Singleton
    public DataSource provideDataSource(DatabaseConfig config) {
        HikariConfig hikariConfig = new HikariConfig();
        hikariConfig.setJdbcUrl(config.getUrl());
        hikariConfig.setUsername(config.getUsername());
        hikariConfig.setPassword(config.getPassword());
        hikariConfig.setMaximumPoolSize(config.getMaximumPoolSize());
        hikariConfig.setConnectionTimeout(config.getConnectionTimeoutMs());
        return new HikariDataSource(hikariConfig);
    }
}
