package uk.co.amrik.steel.persistence;

import java.util.Map;

public class DatabaseConfig {

    private final String host;
    private final String port;
    private final String name;
    private final String user;
    private final String pass;

    public DatabaseConfig(){
        Map<String, String> env = System.getenv();

        this.host = env.getOrDefault("STEEL_DB_HOST", "localhost");
        this.port = env.getOrDefault("STEEL_DB_PORT", "5432");
        this.name = env.getOrDefault("STEEL_DB_NAME", "steel");
        this.user = env.getOrDefault("STEEL_DB_USER", "postgres");
        this.pass = env.getOrDefault("STEEL_DB_PASS", "postgres");

    }

    public String getHost() {
        return host;
    }

    public String getPort() {
        return port;
    }

    public String getName() {
        return name;
    }

    public String getUser() {
        return user;
    }

    public String getPass() {
        return pass;
    }
}
