-- varchar is not long enough and causes bugs
ALTER TABLE docs
    ALTER COLUMN file_path TYPE TEXT;