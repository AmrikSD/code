package uk.co.amrik.rolodex.services.watch;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import uk.co.amrik.rolodex.services.BackgroundService;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;
import java.util.stream.Stream;

public class DirectoryWatcherService extends BackgroundService {

    private static final Logger logger = LoggerFactory.getLogger(DirectoryWatcherService.class);
    private final Set<Path> seenFiles = ConcurrentHashMap.newKeySet();
    private int n = 0;

    @Override
    protected void workLoop() throws Exception {
        try(Stream<Path> paths = Files.list(Path.of("/Users/amrik/code/test-dump"))){
            paths.forEach(path -> {
                if (seenFiles.contains(path.getFileName())){
                    return;
                }
                logger.info("Found file {}", path.getFileName());
                seenFiles.add(path.getFileName());
            });
        }
    }
}
