# Exercise 1

The first step of using a datastore is connecting it to your software program. In this exercise we will be using generic go program and creating a database connection to it.
Using a sqLite database driver and this [main.go](/ex-1-connection/main.go) add a connection to the database of your choosing. After you have connected and verified your connection, explore the database. Make sure to query the database's users table and handle the error. Try running `SELECT`, `INSERT`, and `UPDATE` statements

Follow up questions:
*what kind of package organization would make sense for organizing your database logic?
*what data base driver did you pick?
\*did the data persist?

_NOTE:_ You are free to use any db driver. You can run many datbases locally using docker. Below is an example of running postgres locally in a docker container.

```
docker pull postgres
docker run -e POSTGRES_PASSWORD=my_password -e POSTGRES_USERNAME=user1 -p 5431:5432 postgres
After you get docker running in your local environment setup your database. You will need to CREATE your tables and INSERT data into the table. You can do this in your Go app or via a sql script editor. psql is postgres's command line tool.
```

Solution [here](/databases/ex-1-connection/solution.go)
