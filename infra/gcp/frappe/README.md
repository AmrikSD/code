# Restoring backups

First, get the backup from s3, or r2, or wherever you put it.


SCP the files onto the box.
```bash
scp 20241214_131900_20241214_131900* <url>:~/backups/
```

SSH onto to box.
```bash
ssh <url>.xyz
```

Get a shell on a container with "bench" running
```bash
cd ~/backups
docker compose run -it -v $(pwd):/backup frontend sh

whoami
#this should return frappe, just doulbe check ur actually in the container...
```

Run the restore command
```bash
bench restore /backup/20241214_131900_20241214_131900-frontend-database.sql.gz --with-public-files /backup/20241214_131900_20241214_131900-frontend-files.tar --with-private-files
 /backup/20241214_131900_20241214_131900-frontend-private-files.tar
```

> [!NOTE]
> This last command may look like it hangs, give it some time.

