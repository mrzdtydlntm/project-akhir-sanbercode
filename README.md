# Project Akhir Sanbercode

## Step by Step

1. Run docker compose inside root folder
2. Get inside the database container
```
$ sudo docker exec -it postgres-container bash
```
3. Migrate database file to your table
```
# psql -d sanber -U postgres < /home/db.sql
```
4. Do the rest by yourself!

Good luck!!!!!