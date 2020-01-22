package main

import (
	"database/sql"
	"fmt"
	_ "gopkg.in/goracle.v2"
)

func main()  {
	db, err := sql.Open("goracle", "scott/tiger@10.0.1.127:1521/orclpdb1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()


	rows,err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)

}

