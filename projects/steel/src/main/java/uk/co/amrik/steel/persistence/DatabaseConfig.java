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

        this.host = System.getProperty("STEEL_DB_HOST", env.getOrDefault("STEEL_DB_HOST", "localhost"));
        this.port = System.getProperty("STEEL_DB_PORT", env.getOrDefault("STEEL_DB_PORT", "5432"));
        this.name = System.getProperty("STEEL_DB_NAME", env.getOrDefault("STEEL_DB_NAME", "steel"));
        this.user = System.getProperty("STEEL_DB_USER", env.getOrDefault("STEEL_DB_USER", "postgres"));
        this.pass = System.getProperty("STEEL_DB_PASS", env.getOrDefault("STEEL_DB_PASS", "postgres"));

    }

    public DatabaseConfig(String host, String port, String name, String  user, String pass){
        this.host = host;
        this.port = port;
        this.name = name;
        this.user = user;
        this.pass = pass;
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

    class Builder {

        private String host;
        private String port;
        private String name;
        private String user;
        private String pass;

        public Builder setHost(String host){
            this.host = host;
            return this;
        }

    }

}
