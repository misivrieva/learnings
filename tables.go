package hello

import "github.com/coopernurse/gorp"

type Person struct {
    Id      int64    
    Created int64
    Updated int64
    FName   string
    LName   string
}

// connect to db using standard Go database/sql API
// use whatever database/sql driver you wish
db, err := sql.Open("mymysql", "tcp:localhost:3306*mydb/myuser/mypassword")

// construct a gorp DbMap
dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

table := dbmap.AddTable(Person{}).SetKeys(true, "Id")
table := dbmap.AddTableWithName(Person{}, "People").SetKeys(true, "Id")

// Adding people to the table:

person := &Person{
    Created: time.Now().UnixNow(), 
    Updated: time.Now().UnixNow(),
    FName: "Joe",
    LName: "Smith"
}
err := dbmap.Insert(person)

