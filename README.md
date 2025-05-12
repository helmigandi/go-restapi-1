# go-restapi-1

Learn Go Rest API from Programmer Zaman Now

## Database

- SQL Queries

```bash
mysql -h 127.0.0.1 -P 3306 -u root -p

# Create Table
create table category(
  id integer primary key auto_increment,
  name varchar(255) not null
) engine = InnoDB;

# Select all categories
select * from category;
```

## Source

- [MySql driver repository](https://github.com/go-sql-driver/mysql)
- [Http router repository](https://github.com/julienschmidt/httprouter)
- [Go Validator repository](https://github.com/go-playground/validator)
- [Material slide](https://docs.google.com/presentation/d/1CWDLPYNslBY44Krtzbt-YMntxqocNgMgvrObt5p2uxM/edit?slide=id.p#slide=id.p)
- [Source repository](https://github.com/ProgrammerZamanNow/belajar-golang-restful-api)
