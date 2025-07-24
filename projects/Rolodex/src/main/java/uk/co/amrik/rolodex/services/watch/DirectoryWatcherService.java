package uk.co.amrik.rolodex.services.watch;

import com.google.inject.Inject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import uk.co.amrik.rolodex.services.BackgroundService;

import java.lang.module.Configuration;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
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
    protected void workLoop() throws Exception {
        try(
            Stream<Path> paths = Files.list(Path.of(config.getWatchPath()))
        ){
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
