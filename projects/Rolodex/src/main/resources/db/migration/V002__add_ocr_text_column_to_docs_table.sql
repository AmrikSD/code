ALTER table docs
    ADD COLUMN ocr_text text;

CREATE TABLE ocr_jobs (
    id SERIAL PRIMARY KEY,
    doc_id INTEGER NOT NULL REFERENCES docs(id) ON DELETE CASCADE,
    reserved_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_ocr_jobs_doc_id ON ocr_jobs(doc_id);
