# Golang MySQL driver test

```shell
docker compose down -v && docker compose up -d 

go run db.go 
```

### Versions:

    docker compose exec notWorkingMySql8 mysql --user=root --password=root -e "SELECT version()"
    +-----------+
    | version() |
    +-----------+
    | 8.0.25    |
    +-----------+

    docker compose exec workingMySql5 mysql --user=root --password=rootpwd -e "SELECT version()"
    +-----------+
    | version() |
    +-----------+
    | 5.7.34    |
    +-----------+

----
### Discrepancies:

```
MySQL 5: valid time: result is {1, Me, 2021-06-02 21:05:26.445211}
MySQL 5: year 2 AD: result is {1, Me, 2021-06-02 21:05:26.445211}
MySQL 5: year 1 AD: result is {1, Me, 2021-06-02 21:05:26.445211}
MySQL 5: uninitialised date: result is {1, Me, 2021-06-02 21:05:26.445211}
MySQL 5: null time: record not found

MySQL 8: valid time: result is {1, Me, 2021-06-02 21:05:30.289505}
MySQL 8: year 2 AD: result is {1, Me, 2021-06-02 21:05:30.289505}
MySQL 8: year 1 AD: record not found
MySQL 8: uninitialised date: record not found
MySQL 8: null time: record not found
```

----

### Findings:

* Driver treats "uninitialised" as per `nil` for MySQL >= 8.0.22, similiarly for any year `< 2`
* 8.0.16 >= MySQL >= 8.0.21 returns `Incorrect DATETIME value: '0000-00-00'`
* MySQL <= 8.0.15 works identically to MySQL 5.7
