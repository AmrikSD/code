package uk.co.amrik.rolodex.services.watch;

import com.google.inject.Inject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import uk.co.amrik.rolodex.services.BackgroundService;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;
import java.util.stream.Stream;

public class DirectoryWatcherService extends BackgroundService {

    private static final Logger logger = LoggerFactory.getLogger(DirectoryWatcherService.class);
    private final DirectoryWatcherConfig config;

    private final Set<Path> seenFiles = ConcurrentHashMap.newKeySet();

    @Inject
    public DirectoryWatcherService(DirectoryWatcherConfig config) {
        this.config = config;
    }

    @Override
    protected void workLoop() throws IOException {
        try(
                Stream<Path> fileStream = Files.walk(Path.of(config.getWatchPath()));
        ) {
            fileStream.forEach(file -> {

                if (file.toString().equals(config.getWatchPath())) {
                    return;
                }
                if (seenFiles.contains(file)){
                   return;
                }

                logger.info("{}", file.normalize());
                seenFiles.add(file);

            });
        }
    }
}
