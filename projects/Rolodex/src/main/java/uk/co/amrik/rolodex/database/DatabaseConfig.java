package uk.co.amrik.rolodex.database;

import com.google.inject.Inject;
import com.typesafe.config.Config;

public class DatabaseConfig {

    private final String url;
    private final String username;
    private final String password;
    private final int maximumPoolSize;
    private final int connectionTimeoutMs;

    @Inject
    public DatabaseConfig(Config config){
        Config dbConfig = config.getConfig("database");

        this.url = dbConfig.getString("url");
        this.username = dbConfig.getString("username");
        this.password = dbConfig.getString("password");

        Config poolConfig = dbConfig.getConfig("pool");
        this.maximumPoolSize = poolConfig.getInt("maxSize");
        this.connectionTimeoutMs = poolConfig.getInt("connectionTimeout");

    }

    public String getUrl() {
        return url;
    }

    public String getUsername() {
        return username;
    }

    public String getPassword() {
        return password;
    }

    public int getMaximumPoolSize() {
        return maximumPoolSize;
    }

    public int getConnectionTimeoutMs() {
        return connectionTimeoutMs;
    }
}
