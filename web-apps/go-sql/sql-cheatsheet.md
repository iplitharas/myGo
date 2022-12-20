## PostgresSQL

## Docs
* [Postgresql-docs](https://www.postgresql.org/docs/current/)
* [go-sql-docs](https://pkg.go.dev/database/sql)
* [Tutorial-part-1](https://www.alexedwards.net/blog/introduction-to-using-sql-databases-in-go?hmsr=joyk.com&utm_source=joyk.com&utm_medium=referral)
* [Tutorial-part-2](https://www.alexedwards.net/blog/organising-database-access)

## Table of contents
1. [Installation](#installation)
2. [Intro](#intro)
3. [Manipulation](#manipulation)
4. [Queries](#queries)
5. [Multiple Tables](#multiple-tables)
6. [Aggregate Functions](#aggregate-functions)
7. [Create User](#create-users)

## Installation 
[homebrew](https://wiki.postgresql.org/wiki/Homebrew)

### Verify installation
```bash
psql postgres
```
### quit psql
```bash
\q
```

## list-databases
```bash
\l
```
```
postgres=# \l
                                         List of databases
   Name    |      Owner       | Encoding | Collate | Ctype |           Access privileges
-----------+------------------+----------+---------+-------+---------------------------------------
 postgres  | ioannisplitharas | UTF8     | C       | C     |
 template0 | ioannisplitharas | UTF8     | C       | C     | =c/ioannisplitharas                  +
           |                  |          |         |       | ioannisplitharas=CTc/ioannisplitharas
 template1 | ioannisplitharas | UTF8     | C       | C     | =c/ioannisplitharas                  +
           |                  |          |         |       | ioannisplitharas=CTc/ioannisplitharas
```

## Intro
### Create database
```bash
CREATE DATABASE database-name;
```

### Connect to database
```bash
\c database-name
```

### Switch back to postgres
```bash
\c postgres
```

### See current-user
```SQL
SELECT current_user;
```

### See current-database
```SQL
SELECT current_database;
```

### Drop database
```bash
DROP DATABASE database-name;
```

### Create table

The `CREATE TABLE` statement creates a new table in a database.
It allows one to specify the name of the table and the name of 
each column in the table.

```SQL
CREATE TABLE table_name (
column1 datatype,
column2 datatype,
column3 datatype
);
```

Column constraints are the rules applied to the values of individual columns:

* PRIMARY KEY constraint can be used to uniquely identify the row.
* UNIQUE columns have a different value for every row.
* NOT NULL columns must have a value.
* DEFAULT assigns a default value for the column when no value is specified.
There can be only one **PRIMARY** KEY column per table and multiple 
**UNIQUE** columns.
* 
```SQL

CREATE TABLE student (
  id INTEGER PRIMARY KEY,
  name TEXT UNIQUE,
  grade INTEGER NOT NULL,
  age INTEGER DEFAULT 10
);
```

## Manipulation

The `INSERT INTO` statement is used to add a new record (row) to a table.
It has two forms as shown:
Insert into columns in order.
Insert into columns by name.
```SQL
-- Insert into columns in order:
INSERT INTO table_name
VALUES (value1, value2);

-- Insert into columns by name:
INSERT INTO table_name (column1, column2)
VALUES (value1, value2);
```

### ALTER TABLE Statement

The `ALTER TABLE` statement is used to modify the columns of an 
existing table. When combined with the `ADD COLUMN` clause, 
it is used to add a new column.

```SQL
ALTER TABLE table_name
ADD column_name datatype;
```

### DELETE Statement

The `DELETE` statement is used to delete records (rows) in a table. 
The `WHERE` clause specifies which record or records that should be deleted. 
If the `WHERE` clause is omitted, all records will be deleted.

```SQL
DELETE FROM table_name
WHERE some_column = some_value;
```
### UPDATE Statement

The `UPDATE` statement is used to edit records (rows) in a table. 
It includes a `SET` clause that indicates the column to edit and a `WHERE` clause for 
specifying the record(s).

```SQL
UPDATE table_name
SET column1 = value1, column2 = value2
WHERE some_column = some_value;
```

## Queries

### SELECT Statement

The `SELECT` * statement returns all columns from the 
provided table in the result set. The given query will fetch all 
columns  and records (rows) from the movies table.

```SQL
SELECT *
FROM movies;
```

### WHERE Clause

The `WHERE` clause is used to filter records (rows)
that match a certain condition.
The given query will select all records where the pub_year equals 2017.

```SQL
SELECT title
FROM library
WHERE pub_year = 2017;
```

### AND Operator

The `AND` operator allows multiple conditions to be combined.
Records must match both conditions that are joined by `AND`
to be included in the result set. 
The given query will match any car that is blue and made after 2014.

```SQL
SELECT model 
FROM cars 
WHERE color = 'blue' 
  AND year > 2014;
```

### AS Clause

Columns or tables can be aliased using the `AS` clause.
This allows columns or tables to be specifically 
renamed in the returned result set.
The given query will return a result set with
the column for name renamed to movie_title.

```SQL
SELECT name AS 'movie_title'
FROM movies;
```

### OR Operator

The `OR` operator allows multiple conditions to be combined.
Records matching either condition joined by the OR are included
in the result set.
The given query will match customers whose state is either 'CA' or 'NY'.

```SQL
SELECT name
FROM customers 
WHERE state = 'CA' 
   OR state = 'NY';
```

### % Wildcard

The `%` wildcard can be used in a LIKE 
operator pattern to match zero or more unspecified character(s).
The given query will match any movie that begins with `The`,
followed by zero or more of any characters.

```SQL
SELECT name
FROM movies
WHERE name LIKE 'The%';
```

### _ Wildcard

The `_` wildcard can be used in a 
LIKE operator pattern to match any single
unspecified character. 
The given query will match any movie which begins with a 
single character, followed by ove.

```SQL
SELECT name
FROM movies
WHERE name LIKE '_ove';
```

### ORDER BY Clause

The `ORDER` BY clause can be used to sort the result set 
by a particular column either alphabetically or numerically.
It can be ordered in two ways:
* `DESC` is a keyword used to sort the results in descending order.
* `ASC` is a keyword used to sort the results in ascending order (default).

```SQL
SELECT *
FROM contacts
ORDER BY birth_date DESC;
```

### LIKE Operator

The `LIKE` operator can be used inside of a `WHERE` clause to 
match a specified pattern.
The given query will match any movie that begins with Star in its title.

```SQL
SELECT name
FROM movies
WHERE name LIKE 'Star%';
```

### DISTINCT Clause

Unique values of a column can be selected using a `DISTINCT` query.
For a table contact_details having five rows in which the city column contains Chicago, Madison, Boston, Madison, 
and Denver, the given query would return:
* Chicago
* Madison
* Boston
* Denver
* 
```SQL
SELECT DISTINCT city
FROM contact_details;
```

### BETWEEN Operator

The `BETWEEN` operator can be used to filter by a range of values.
The range of values can be text, numbers, or date data.
The given query will match any movie made between the 
years 1980 and 1990, inclusive.

```SQL
SELECT *
FROM movies
WHERE year BETWEEN 1980 AND 1990;
```

### LIMIT Clause

The `LIMIT` clause is used to narrow, or limit,
a result set to the specified number of rows.
The given query will limit the result set to 5 rows.

```SQL
SELECT *
FROM movies
LIMIT 5;
```

### NULL Values

Column values can be `NULL`, or have no value.
These records can be matched (or not matched)
using the IS NULL and IS NOT NULL operators in combination
with the WHERE clause.
The given query will match all addresses where the address 
has a value or is not NULL.

```SQL
SELECT address
FROM records
WHERE address IS NOT NULL;
```

## Multiple Tables
For example only employee can have multiple phone numbers
In Phone numbers table:  
`EMP_ID INT references TABLE_name(ID)`

### Foreign Key
A foreign key is a reference in one table’s records to the primary key of another table.
To maintain multiple records for a specific row, the use of foreign key plays a vital role.
For instance, to track all the orders of a specific customer, the table order
can contain a foreign key.

### Primary Key
A primary key column in a SQL table is used to uniquely identify each record in that table.
A primary key cannot be NULL. 
**Primary keys are often used in JOIN operations**.


### WITH Clause
The WITH clause stores the result of a query in a 
temporary table (temporary_movies) using an alias.

Multiple temporary tables can be defined with one
instance of the `WITH` keyword.
```SQL
WITH temporary_movies AS (
   SELECT *
   FROM movies
)
SELECT *
FROM temporary_movies
WHERE year BETWEEN 2000 AND 2020;
```

### UNION Clause
The UNION clause is used to combine results 
that appear from multiple SELECT statements and filter duplicates.

For example, given a first_names table with 
a column name containing rows of data “James” and “Hermione”, 
and a last_names table with a column name containing rows of data 
“James”, “Hermione” and “Cassidy”, the result of this query would contain three 
names: “Cassidy”, “James”, and “Hermione”.
```SQL
SELECT name
FROM first_names
UNION
SELECT name
FROM last_names
```

### CROSS JOIN Clause
The `CROSS JOIN` clause is used to **combine** each row from one 
table with each row from another in the result set. 
This JOIN is helpful for creating all possible combinations for the records (rows) in two tables.
The given query will select the shirt_color and pants_color columns from the result set, 
which will contain all combinations of combining the rows in the shirts and pants tables. 
If there are 3 different shirt colors in the shirts table and 5 different 
pants colors in the pants table then the result set will contain 3 x 5 = 15 rows
```SQL
SELECT shirts.shirt_color,
   pants.pants_color
FROM shirts
CROSS JOIN pants;
```

### Inner Join
The JOIN clause allows for the return of results 
from more than one table by joining them together with other results based on 
common column values specified using an ON clause.
INNER JOIN is the default `JOIN` and it will only return results **matching** the **condition** **specified** by `ON`.
```SQL
SELECT * 
FROM books
JOIN authors
  ON books.author_id = authors.id;
```

### Outer Join
An outer join will combine rows from  different tables even if the join 
condition **is not met**. In a `LEFT JOIN`, every row in the left table is 
returned  in the result set, 
and **if the join condition is not met**, then `NULL` values are used to **fill
in the columns from the right table**.


```SQL
SELECT column_name(s)
FROM table1
LEFT JOIN table2
  ON table1.column_name = table2.column_name;
```


## Aggregate Functions

### MIN() Aggregate Function

The MIN() aggregate function returns the smallest value in a column.
For instance, to find the smallest value of the amount column from the
table named transactions, the given query can be used.

```SQL
SELECT MIN(amount) 
FROM transactions;
```

### AVG() Aggregate Function

The AVG() aggregate function returns the average value in a column.
For instance, to find the average salary for the employees who have 
less than 5 years of experience, the given query can be used.
```SQL
SELECT AVG(salary)
FROM employees
WHERE experience < 5;
```

### MAX() Aggregate Function

The MAX() aggregate function takes the name of a column as an argument
and returns the largest value in a column.
The given query will return the largest value from the amount column.

```SQL
SELECT MAX(amount) 
FROM transactions;
```

### COUNT() Aggregate Function
The COUNT() aggregate function returns the total number of rows 
that match the specified criteria. For instance, to find the total 
number of employees who have less than 5 years of experience, the 
given query can be used.
Note: A column name of the table can also be used instead of *.
Unlike COUNT(*), this variation COUNT(column) will not count NULL 
values in that column.

```SQL
SELECT COUNT(*)
FROM employees
WHERE experience < 5;
```

### SUM() Aggregate Function

The SUM() aggregate function takes the name of a column as an argument and 
returns the sum of all the value in that column.

```SQL
SELECT SUM(salary)
FROM salary_disbursement;
```

### ROUND() Function

The ROUND() function will round a number value to a 
specified number of places. It takes two arguments: a number,
and a number of decimal places. It can be combined with other aggregate functions,
as shown in the given query.
This query will calculate the average rating of movies from 2015,
rounding to 2 decimal places.

```SQL
SELECT year, 
   ROUND(AVG(rating), 2) 
FROM movies 
WHERE year = 2015;
```

### Column References

The GROUP BY and ORDER BY clauses can reference the selected 
columns by number in which they appear in the SELECT statement. The example query will count the number of movies per rating, and will:
GROUP BY column 2 (rating)
ORDER BY column 1 (total_movies)

```SQL
SELECT COUNT(*) AS 'total_movies', 
   rating 
FROM movies 
GROUP BY 2 
ORDER BY 1;
```

### HAVING Clause

The `HAVING` clause is used to further filter the result set groups 
provided by the GROUP BY clause. 
`HAVING` is often used with aggregate functions to filter 
the result set groups based on an aggregate property.
The given query will select only the records (rows) from only 
years where more than 5 movies were released per year.
The HAVING clause **must always come after** a `GROUP BY` clause
but must come before any ORDER BY or LIMIT clause.

```SQL
SELECT year, 
   COUNT(*) 
FROM movies 
GROUP BY year
HAVING COUNT(*) > 5;
```

## Create users

### Details of users
```bash
\du
```

### Create user

```SQL
CREATE USER ioannis WITH PASSWORd "password";
```

### Grant privileges
```SQL
GRANT ALL PRIVILEGES ON DATABASE company TO ioannis;
```

### Revoke privileges
```SQL
REVOKE ALL PRIVILEGES ON DATABASE company TO iaonnis;
```

### ALTER 
```SQL
ALTER USER ioannis WITH SUPERUSER;
```

### REMOVE
```SQL
DROP USER ioannis;
```
