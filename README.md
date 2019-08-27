# gin-skeleton

[![Build Status](https://travis-ci.org/hyperjiang/gin-skeleton.svg?branch=master)](https://travis-ci.org/hyperjiang/gin-skeleton)
[![Go Report](https://goreportcard.com/badge/github.com/hyperjiang/gin-skeleton)](https://goreportcard.com/report/github.com/hyperjiang/gin-skeleton)
[![License](https://img.shields.io/github/license/hyperjiang/gin-skeleton.svg)](https://github.com/hyperjiang/gin-skeleton)

Gin Skeleton is a simple boilerplate to kickstart a web server project based on Gin Framework.

Require go version >= 1.11, if your go version is 1.10, please use branch `go1.10`.

```
# start a web server listening on 0.0.0.0:8080
go run main.go
```

## Components

- Framework: [gin-gonic/gin](https://github.com/gin-gonic/gin)
- Database ORM: [jinzhu/gorm](https://github.com/jinzhu/gorm)
- Database migration: [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)
- Leveled logs: [golang/glog](https://github.com/golang/glog)
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
