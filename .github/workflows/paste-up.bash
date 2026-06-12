#!/bin/bash
# This is a sample script for uploading a sqlitedb into a remote droplet

# Save wal trans 
sqlite3 ~/Source/Repos/Polybub/.db/env/env_sqlitedb 'PRAGMA wal_checkpoint(FULL);'
# Exit wal mod   
sqlite3 ~/Source/Repos/Polybub/.db/env/env_sqlitedb 'PRAGMA journal_mode=DELETE;'
# Paste it up     
rsync -av ~/Source/Repos/Polybub/.db/env/env_sqlitedb droplet1:/var/www/app/Polybub/.db/sqlitedb

# Re-perm the new db 
ssh droplet1 "sudo chown github_agent:github_agent /var/www/app/Polybub -R"
ssh droplet1 "sudo chmod 600 /var/www/app/Polybub/.db/sqlitedb"

# MANUAL restart via Github Actions deploy.yml