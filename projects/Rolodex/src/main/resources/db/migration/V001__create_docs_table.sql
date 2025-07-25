CREATE TABLE docs (
                      id SERIAL PRIMARY KEY,
                      file_name VARCHAR(255) NOT NULL,
                      file_path VARCHAR(500) NOT NULL,
                      mime_type VARCHAR(100),
                      md5_hash uuid, -- stores MD5 as UUID (128 bits)
                      indexed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_docs_hash on docs(md5_hash);
