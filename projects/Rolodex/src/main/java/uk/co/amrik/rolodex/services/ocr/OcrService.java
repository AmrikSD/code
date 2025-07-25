package uk.co.amrik.rolodex.services.ocr;

import com.google.inject.Inject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import uk.co.amrik.rolodex.services.WorkerPoolBackgroundService;

import javax.sql.DataSource;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.Optional;

public class OcrService extends WorkerPoolBackgroundService<Optional<OcrJobReservation>> {

    private final Logger logger = LoggerFactory.getLogger(OcrService.class);
    private final DataSource dataSouce;

    @Inject
    public OcrService(DataSource dataSource){
        super(4);
        this.dataSouce = dataSource;
    }

    @Override
    protected Optional<OcrJobReservation> discoverWork() throws Exception {
        return getPathFromDB();
    }

    @Override
    protected void processWork(Optional<OcrJobReservation> work) throws Exception {
        if (work.isEmpty()) {
            return;
        }

        OcrJobReservation reservation = work.get();
        logger.info("Processing file: {}", reservation.filePath());

        String deleteSql = "DELETE FROM ocr_jobs WHERE doc_id = ?";
        String updateSql = "UPDATE docs SET ocr_text = ? WHERE id = ?";

        try (
                Connection conn = dataSouce.getConnection();
                PreparedStatement updateStmt = conn.prepareStatement(updateSql);
                PreparedStatement deleteStmt = conn.prepareStatement(deleteSql)
        ) {
            conn.setAutoCommit(false); // Start transaction

            updateStmt.setString(1, "processed by worker");
            updateStmt.setLong(2, reservation.docId());
            updateStmt.executeUpdate();

            deleteStmt.setLong(1, reservation.docId());
            deleteStmt.executeUpdate();

            conn.commit(); // Commit both changes atomically
        } catch (SQLException e) {
            logger.error("Failed to process OCR job atomically", e);
            throw e;
        }
    }

    private Optional<OcrJobReservation> getPathFromDB(){
        String sql = """
            WITH doc AS (
                SELECT id, file_path
                FROM docs
                WHERE ocr_text IS NULL
                AND NOT EXISTS (
                    SELECT 1 FROM ocr_jobs WHERE ocr_jobs.doc_id = docs.id
                )
                LIMIT 1
                FOR UPDATE SKIP LOCKED
            )
            INSERT INTO ocr_jobs (doc_id, reserved_at)
            SELECT id, NOW() FROM doc
            RETURNING doc_id, (SELECT file_path FROM doc)
        """;
        try (
                Connection conn = dataSouce.getConnection();
                PreparedStatement stmt = conn.prepareStatement(sql);
                java.sql.ResultSet rs = stmt.executeQuery()
        ) {
            if (rs.next()) {
                long docId = rs.getLong(1);
                String filePath = rs.getString(2);
                return Optional.of(new OcrJobReservation(rs.getLong(1), rs.getString(2)));
            }
        } catch (SQLException e) {
            logger.error("Failed to reserve and fetch OCR job", e);
        }
        return Optional.empty();
    }

}
