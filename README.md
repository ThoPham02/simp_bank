-- Lecture 1: Create database schema with dbdiagram.io --

Follow `dbdiagram.io` support to write database schema
and export file sql to use in Lecture 2

--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 2: Docker + PostgreSQL + Table Plus --

use `Docker` to config PostgreSQL image and create Container PostgreSQL
use `TablePlus` for database management

--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 3: Database migration --

In Go, have many packages to migrate database, example: 
 - dbmate
 - golang-migrate
 - goose
 - sql-migrate
 - tern

We use `migrate`. Following docs `https://github.com/golang-migrate/migrate`.
 - Install Migrate
 - To create files up and down use command `migrate create -ext sql -dir db/migration -seq name_of_migration`.
    Explain flags: -ext : to create up and down migrations.
                -dir : to save the migrations in the folder
                -seq : to sequential version number for the migration file.
    See more use `migrate -help` command
 - Add sql to create database in up file and drop in down file.
 - To run migration file up and down, see command in `Makefile`.

--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 4: Generate CRUD Golang code use sqlc --

Some features of SQLC:
 - Very fast and easy to use
 - Automatic code generation
 - Catch SQL error before generating codes
 - Support PostgreSQL, MySQL and SQLite databases

Install sqlc:
For Ubuntu use command: `sudo snap install sqlc`

I use WSL2 and I got an error: 
cannot communicate with server: Post http://localhost/v2/snaps/sqlc: dial unix /run/snapd.socket: connect: no such file or directory
To fix this error: See more in `https://github.com/microsoft/WSL/issues/5126`
 - First, install ubuntu-wsl2-systemd-script:
    git clone https://github.com/DamionGans/ubuntu-wsl2-systemd-script.git
    cd ubuntu-wsl2-systemd-script/
    bash ubuntu-wsl2-systemd-script.sh
    systemctl
 - Now, use command: sudo snap install sqlc.
 - After this, i got another error: snap/bin is not PATH. to fix
    nano ~/.bashrc
    add export PATH="$PATH:/snap/bin"  below the bottom. Ctrl + X, Y, Enter to save.
 - Now, everything is excellent.

To use sqlc: See more in `https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html`
 - First, run command: `sqlc init`
 - Now, we are having file sqlc.yaml, custom sqlc.yaml.
 - write query in db/query following docs.
 - To generate code, use command: `sqlc generate`

--------------------------------------------------------------------------------------------------------------------------------



