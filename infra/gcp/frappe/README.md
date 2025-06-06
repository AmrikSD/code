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
> [!NOTE]
> This command may look like it hangs, just give it some time.

```bash
bench restore /backup/20241214_131900_20241214_131900-frontend-database.sql.gz --with-public-files /backup/20241214_131900_20241214_131900-frontend-files.tar --with-private-files
 /backup/20241214_131900_20241214_131900-frontend-private-files.tar
```

put the enryption key back
```
cd ~/frappe-bench/sites/frontend
# copy the old file
cp site_config.json site_config.json.old
# look in /backup/20241217_190303_20241217_190303-frontend-site_config_backup.json
# put the encryption_key there
#
```

do not change the password or db name... even after the restore file the creds do not go to the old ones.

If you have migrated to a newer version of frappe, you may need to run migrations before the website works again.
```
bench migrate
```

