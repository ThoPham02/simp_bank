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

To use sqlc: See more in `https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html`
 - First, run command: `sqlc init`
 - Now, we are having file sqlc.yaml, custom sqlc.yaml.
 - write query in db/query following docs.
 - To generate code, use command: `sqlc generate`
--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 5: Write Unit Tests --

3 Rules for writing unit tests:
 - file test must end with _test.go
 - function test must start with Test(with uppercase T letter)
 - function test must accept t *testing.t as an agument

There are many lib to write unit tests, we use `testify`.
Testify have 4 sub-packages:
 - `asert`  : return boolean
 - `require`: same global function as assert package, 
   but instead of returning a boolean result they terminate current test. 
 - `mock`   :
 - `suite`  :
READ MORE: `https://github.com/stretchr/testify`

First, we write TestMain function to connect to database
Next, we write unit tests for all functions.
--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 6: Implement Database Transaction --

Database Transaction:
 - A single unit of work
 - Often made up of multiple database operations
For Example: Simple Bank To Transfer 10 USD from Acc1 to Acc2:
 - Create a transfer recode with amount = 10
 - Create an account entry for account 1 with amount = -10
 - Create an account entry for account 2 with amount = 10
 - Subtract 10 from the balance of account 1
 - Add 10 to the balance of account 2
2 Main Reason for using Database Transaction:
 - To provide a reliable and consistent unit of work, even in case of system failure.
 - To provide isolation between programs that access the database concurrently.
ACID Properties:
 - A (Atomicity) Either all operations complete successfully or the transaction fails and the database is unchanged.
 - C (Consistency) The db state must be valid after the transaction. All constraints must be satisfied.
 - I (Isolation) Concurrent transactions must not affect each other.
 - D (Durability) Data written by a successful transaction must be recorded in persistent storage.
--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 7: DB Transaction lock and Handle deadlock --

This lecture, we will learn about deadlock.
Deadlock occurs because postgres is afraid that we will update the value in the foreign key field.
To solve this problem, we can remove the foreign key(bad way) or add the Select statement `FOR NO KEY UPDATE`.
--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 8: How to avoid deadlock --

The best way to avoid deadlock is make sure that `application always acquire locks in a consistent order`
--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 9: Insolation level in Mysql and PostgreSQL --

There are 4 levels of isolation:
 - Read Uncommitted
 - Read Committed
 - Repeatable Read
 - Serializable
Read More: 
Mysql: `https://dev.mysql.com/doc/refman/8.0/en/innodb-transaction-isolation-levels.html`
PostgreSQL: `https://www.postgresql.org/docs/current/transaction-iso.html`
--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 10: Setup Github Action to run automated test --

We can trigger workflow y three ways event, scheduled or manually.
A workflow consists of one or more jobs.
A Job is composed of a single or multiple steps.
Each step can have one or more actions
All job inside workflow normally run in parallel unless they depend on each other.That case they run serially.
Each Job will be run separaely by a specific runner.
The runners will report progress locks and results of the jobs back to the github.

ReadMore: 
Github Action: 
--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 11: RESTful API with Go Gin --
Read More: `https://github.com/gin-gonic/gin`
--------------------------------------------------------------------------------------------------------------------------------

--Lecture 12: Read config with Viper --
Read More: `https://github.com/spf13/viper`
--------------------------------------------------------------------------------------------------------------------------------

-- Lecture 13: Use Mock API to Test API --
Change sqlc.yml file: `emit_interface: true` to generate Querier Interface. The Querier Interface has all functions generated by queries.
Change Store to SQLStore and add Store interface. SQLStore will implement Store interface.
Use mockgen to generate file store.go(mock store).

Write code to know how to setup Test API.