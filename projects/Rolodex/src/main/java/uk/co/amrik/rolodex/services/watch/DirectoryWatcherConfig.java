package uk.co.amrik.rolodex.services.watch;

import com.google.inject.Inject;
import com.typesafe.config.Config;

public class DirectoryWatcherConfig {

    private final String watchPath;

    @Inject
    public DirectoryWatcherConfig(Config config) {
        Config watcherConfig = config.getConfig("directory.watcher");
        this.watchPath = watcherConfig.getString("path");
    }

    public String getWatchPath() {
        return watchPath;
    }

}
