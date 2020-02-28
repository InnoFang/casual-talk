# casual-talk

Build a forum in Go. Learn from [gwp](https://github.com/sausheong/gwp) and make some change.

[![](https://img.shields.io/badge/license-MIT-yellowgreen)](./LICENSE) ![](https://img.shields.io/badge/database-MySQL-blue)


## How to run

Clone this project under `$GOPATH/src`, then create a database (it's ok whatever you want to named) and create some tables 
by following [`data/setup.sql`](./data/setup.sql). After that, configure your MySql username, password and database name 
in [`data/sql.json`](./data/sql.json).

A MySql driver is required

```
$ go get github.com/go-sql-driver/mysql
```

then you can run it
```
$ cd $GOPATH/src/casual-talk

$ go run main.go 
```

BTW, if you want to run this project in Docker directly, you need to link this container with MySQL container, that is to 
say you have been pull MySQL image and run MySQL with docker before you run this.

```
$ docker build -t casual-talk .
$ docker run --link mysql:mysql -p 8080:8080 casual-talk
```
