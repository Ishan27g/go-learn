package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func appendTask(db *sql.DB, num int, task string) {

	str := "Insert INTO tasks(task) values ('" + task + "')"
	statement, _ := db.Prepare(str)
	statement.Exec()

}
func displayTask(db *sql.DB) {

	str := "SELECT * FROM tasks"
	rows, err := db.Query(str)
	if err != nil {
		fmt.Println("Error querying db")
	}
	var id int
	var task string
	for rows.Next() {
		rows.Scan(&id, &task)
		fmt.Println(id, task)
	}

}

func delTask(db *sql.DB, id int) {
	//todo add check for valid id before deleting
	str := "DELETE FROM tasks where id = " + strconv.Itoa(id)
	statement, _ := db.Prepare(str)
	statement.Exec()
}

func main() {

	var taskNum int = 0
	taskMap := make(map[int]string)
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		fmt.Println("error opening db")
		fmt.Println(err)
	}

	var cmdAdd = &cobra.Command{
		Use:   "add [string to print]",
		Short: "Add a task",
		Long:  `Add a task to the list`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskMap[taskNum] = strings.Join(args, " ")
			appendTask(db, taskNum, taskMap[taskNum])
			taskNum++
		},
	}

	var cmdDisplay = &cobra.Command{
		Use:   "display",
		Short: "display all tasks",
		Long:  `display all tasks created`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			displayTask(db)
		},
	}

	var cmdDel = &cobra.Command{
		Use:   "del [number]",
		Short: "delete a task",
		Long:  `delete an existing task`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			/*if len(args) > 1 {
				for _, arg := range args {
					taskList = append(taskList, strconv.Atoi(stings.arg))
				}
			}
			*/
			taskNumber, _ := strconv.Atoi(strings.Join(args, " "))
			delTask(db, taskNumber)
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdAdd, cmdDel, cmdDisplay)
	rootCmd.Execute()

}
