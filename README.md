# JWT Authentication in Go

## Tooling

- GORM (Object Relational Model)
- Gin (REST API framework)
- bcrypt (go extended library for adaptive hashing algorithms)
- jwt-go (A go implementation of JSON Web Tokens)
- GoDotEnv (load env vars from a .env file)
- CompileDaemon (watcher for live update) similar to _air_.

## Installing deps

## Getting postgres up!

#### Step 1
```
docker pull postgres
```

#### Step 2
```
docker run -d --name AuthN -p 5432:5432 -e POSTGRES_PASSWORD=pwd123 postgres
```

- `-d` flag specifies that the container should execute in the background.
- `--name` option assigns the container’s name, i.e., “AuthN”.
- `-p` assigns the port for the container i.e. “5432:5432”.
- `-e` POSTGRES_PASSWORD” configures the password to be `pwd123`.
- `postgres` is the official Docker image:

#### Step 3

Accessing container via `bash`.

```
docker exec -it AuthN bash
```
#### Step 4

Connect to a PostgreSQL Database Server.

```
root@...:/# psql -h localhost -U postgres
```

#### Step 5

Create a PostgreSQL Database

```
CREATE DATABASE authn_jwt;
```

Confirm database creation.

```\l```

Establish a Connection With a Database

```\c authn_jwt;```

Create a table in the database

```
CREATE TABLE users (
    ID INT PRIMARY KEY NOT NULL, 
    NAME TEXT NOT NULL);
```


