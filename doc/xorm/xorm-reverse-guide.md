# xorm reverse guide

`xorm` is a golang orm。 `xorm reverse` is a go struct generate tools help to generate go code from table.

## install

```bash
go get xorm.io/reverse
```

## usage

```bash
reverse -f config.yml
```

sample config file

```yml
kind: reverse
name: my_database
source:
  database: mysql
  conn_str: "user:password@tcp(127.0.0.1:3306)/my_database?charset=utf8mb4"
targets:
  - type: codes
    language: golang
    output_dir: models
```

After connect to database, it will generate struct for all tables, output to `output_dir`