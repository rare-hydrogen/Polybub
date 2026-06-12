#!/bin/bash
# This is a sample script for copying a sqlitedb out of a remote droplet

# Kill the app   
ssh droplet1 sudo systemctl stop Polybub.service
# Save wal trans 
ssh droplet1 "sqlite3 /var/www/app/Polybub/.db/sqlitedb 'PRAGMA wal_checkpoint(FULL);'"
# Exit wal mod   
ssh droplet1 "sqlite3 /var/www/app/Polybub/.db/sqlitedb 'PRAGMA journal_mode=DELETE;'"
# Copy it down   
rsync -av droplet1:/var/www/app/Polybub/.db/sqlitedb ~/Source/Repos/Polybub/.db/env/env_sqlitedb

# You may need to disconnect and reconnect