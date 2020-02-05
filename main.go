package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

func main() {

	newSemaphoreValue := 0
	userName := "IBS"
	userPassword := "IBS"
	standName := "AAAAA"
	standIP := "11.22.33.44"
	var semaphoreValue string

	/* Открыть соединение */
	db, err := sql.Open("godror", fmt.Sprintf("%s/%s@%s:1521/%s", userName, userPassword, standIP, standName))
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Поставить блокировку */
	query := "select IBS.executor.lock_open() from dual"
	_, err = db.Query(query)
	if err != nil {
		fmt.Printf("Error executing LockOpen query: %s", err)
		return
	}

	/* Прочитать значение семафора */
	semaphoreValue = getSemaphoreValue(err, db)
	fmt.Printf("Old semaphore value: '%s'\n", semaphoreValue)

	/* Установить значение семафора */
	query = fmt.Sprintf("update IBS.Z#SHB_DWH_UNLOAD_S set c_status = %d where id = 1261992816", newSemaphoreValue)
	fmt.Printf("Set semaphore value to: '%d'\n", newSemaphoreValue)
	_, err = db.Query(query)
	if err != nil {
		fmt.Printf("Error setting Semaphore query: %s", err)
		return
	}

	/* Прочитать значение семафора */
	semaphoreValue = getSemaphoreValue(err, db)
	fmt.Printf("New semaphore value: '%s'\n", semaphoreValue)

	/* Закрыть соединение */
	err = db.Close()
	if err != nil {
		fmt.Printf("Error close DB: %s", err)
	}

}

/* Вернуть значение семафора */
func getSemaphoreValue(err error, db *sql.DB) string {

	var semaphoreValue string

	rows, err := db.Query("select c_status from IBS.Z#SHB_DWH_UNLOAD_S where id = 1261992816")
	if err != nil {
		fmt.Printf("Error running Get Semaphore query: %s", err)
	} else {
		for rows.Next() {
			err = rows.Scan(&semaphoreValue)
			if err != nil {
				fmt.Printf("Error in Scan: %s", err)
			}
		}
		err = rows.Close()
		if err != nil {
			fmt.Printf("Error close rows: %s", err)
		}
	}
	return semaphoreValue
}
