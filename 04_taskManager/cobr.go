package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {

	//	var taskList []int
	var taskNum int = 0

	taskMap := make(map[int]string)

	var cmdAdd = &cobra.Command{
		Use:   "add [string to print]",
		Short: "Add a task",
		Long:  `Add a task to the list`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskMap[taskNum] = strings.Join(args, " ")
			taskNum++
		},
	}
	/*
		var cmdEcho = &cobra.Command{
			Use:   "list",
			Short: "list all tasks",
			Long:  `list all tasks created`,
			Args:  cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("<<listing all task here>>")
			},
		}
	*/
	var cmdDisplay = &cobra.Command{
		Use:   "display",
		Short: "display all tasks",
		Long:  `display all tasks created`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("<<displaying all task here>>")
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
			fmt.Println(taskNumber)
		},
	}

	//cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "app"}
	//rootCmd.AddCommand(cmdPrint, cmdEcho)
	//cmdEcho.AddCommand(cmdTimes)

	rootCmd.AddCommand(cmdAdd, cmdDel, cmdDisplay)

	rootCmd.Execute()

}
