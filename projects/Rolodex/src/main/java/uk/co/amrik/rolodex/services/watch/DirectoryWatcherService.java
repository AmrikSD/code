package uk.co.amrik.rolodex.services.watch;

import com.google.inject.Inject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import uk.co.amrik.rolodex.services.BackgroundService;

import javax.sql.DataSource;
import java.io.FileInputStream;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;
import java.util.stream.Stream;

public class DirectoryWatcherService extends BackgroundService {

    private static final Logger logger = LoggerFactory.getLogger(DirectoryWatcherService.class);
    private final DirectoryWatcherConfig config;
    private final DataSource dataSource;

    private final Set<Path> seenFiles = ConcurrentHashMap.newKeySet();

    @Inject
    public DirectoryWatcherService(DirectoryWatcherConfig config, DataSource dataSource) {
        this.config = config;
        this.dataSource = dataSource;
    }

    @Override
    protected void workLoop() throws IOException {
        List<Path> newFiles = new ArrayList<>();

        try(Stream<Path> fileStream = Files.walk(Path.of(config.getWatchPath()))) {
            fileStream.filter(Files::isRegularFile)
                    .filter(file -> !seenFiles.contains(file))
                    .forEach(file -> {
                        logger.info("{}", file.normalize());
                        newFiles.add(file);
                        seenFiles.add(file);
                    });
        }
        if (!newFiles.isEmpty()){
            batchInsertFileRecords(newFiles);
        }
    }

    private void batchInsertFileRecords(List<Path> files) {
        String sql = """
            INSERT INTO docs (file_name, file_path, mime_type, md5_hash)
            VALUES (?,?,?,?::uuid)
            ON CONFLICT DO NOTHING
        """;

        try (Connection conn = dataSource.getConnection();
             PreparedStatement stmt = conn.prepareStatement(sql)) {

            conn.setAutoCommit(false);

            for (Path file : files) {
                String fileHash = md5(file);
                stmt.setString(1, file.getFileName().toString());
                stmt.setString(2, file.normalize().toString());
                stmt.setString(3, Files.probeContentType(file));
                stmt.setString(4, fileHash);
                stmt.addBatch();
                logger.debug("Added to batch: {} with hash: {}", file, fileHash);
            }

            int[] results = stmt.executeBatch();
            conn.commit();
            int insertedCount = Arrays.stream(results)
                    .filter(result -> (result > 0))
                    .sum();

            logger.info("Batch inserted {} new file records ({} duplicates skipped)", insertedCount, files.size() - insertedCount);

        } catch (SQLException e) {
            logger.error("Failed to batch insert file records", e);
        } catch (IOException | NoSuchAlgorithmException e) {
            throw new RuntimeException(e);
        }
    }

    private String md5(Path filePath) throws IOException, NoSuchAlgorithmException {
        MessageDigest md = MessageDigest.getInstance("MD5");

        try (FileInputStream fis = new FileInputStream(filePath.toFile())) {
            byte[] buffer = new byte[8192];
            int bytesRead;

            while ((bytesRead = fis.read(buffer)) != -1) {
                md.update(buffer, 0, bytesRead);
            }
        }

        byte[] hashBytes = md.digest();

        StringBuilder result = new StringBuilder();
        for (byte b : hashBytes) {
            result.append(String.format("%02x", b));
        }
        return result.toString();
    }
}
