package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

/*

CREATE TABLE test (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(35),
    age INT,
    PRIMARY KEY (`id`)
);

 */

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	stmt,err := db.Prepare("insert into test(name,age)values(?,?)");
	if err!=nil {
		fmt.Println(err.Error());
		return;
	}
	defer stmt.Close();
	if result,err := stmt.Exec("zhangsan",20);err==nil {
		if id,err := result.LastInsertId();err==nil {
			fmt.Println("insert id : ",id);
		}
	}
	if result,err := stmt.Exec("lisi",30);err==nil {
		if id,err := result.LastInsertId();err==nil {
			fmt.Println("insert id : ",id);
		}
	}
	if result,err := stmt.Exec("wangwu",25);err==nil {
		if id,err := result.LastInsertId();err==nil {
			fmt.Println("insert id : ",id);
		}
	}

	rows,err := db.Query("select * from test");
	if err!=nil {
		fmt.Println(err.Error());
		return;
	}
	defer rows.Close();
	var id int
	var name string
	var age int
	for rows.Next() { //开始循环、像游标吗？必须rows.Next()哦
		err = rows.Scan(&id, &name, &age) //扫每一行，并把字段的值赋到id,name,age里面去
		if err == nil {
			fmt.Printf("query result, Row : id=%d name=%v age=%v\n", id, name, age)
		}
	}
}
