# gin-skeleton

[![CI](https://github.com/hyperjiang/gin-skeleton/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/hyperjiang/gin-skeleton/actions/workflows/ci.yml)
[![Go Report](https://goreportcard.com/badge/github.com/hyperjiang/gin-skeleton)](https://goreportcard.com/report/github.com/hyperjiang/gin-skeleton)
[![License](https://img.shields.io/github/license/hyperjiang/gin-skeleton.svg)](https://github.com/hyperjiang/gin-skeleton)

Gin Skeleton is a simple boilerplate to kickstart a web server project based on Gin Framework.

Require go version >= 1.18, if your go version is lower, please use legacy branches,
there are quite a lot of incompatible changes between master and legacy branches.

```
# start a web server listening on 0.0.0.0:8080
go run main.go
```

## Components

- Framework: [gin-gonic/gin](https://github.com/gin-gonic/gin)
- Database ORM: [go-gorm/gorm](https://github.com/go-gorm/gorm)
- Database migration: [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)
- Zero Allocation JSON Logger: [rs/zerolog](https://github.com/rs/zerolog)
- YAML support: [go-yaml/yaml](https://github.com/go-yaml/yaml)
- Testing toolkit: [stretchr/testify](https://github.com/stretchr/testify)

## Configuration

Edit the `config.yml` with your own config

## Database Migration

**Create the database first**

```
CREATE DATABASE IF NOT EXISTS `gin` DEFAULT CHARACTER SET utf8mb4;
```

**Migrates the database to the most recent version available**

```
./migrate.sh up
```

**Undo a database migration**

```
./migrate.sh down
```

**Show migration status**

```
./migrate.sh status
```

**Create a new migration**

```
./migrate.sh new a_new_migration
```

## Available endpoints

See [router](https://github.com/hyperjiang/gin-skeleton/blob/master/router/router.go)

- Home page: http://localhost:8080/

- Api version: http://localhost:8080/api/version

- Sign up: http://localhost:8080/signup

- Login: http://localhost:8080/login

- Get user info: http://localhost:8080/user/1

- Test jwt: http://localhost:8080/auth/hello (only user "admin" can see this page)
