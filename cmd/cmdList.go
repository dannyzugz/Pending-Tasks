/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/DanyZugz/zas/models"
	"github.com/spf13/cobra"
)

// cmdListCmd represents the cmdList command
var cmdListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas las tareas en la lista.",
	Long:  ` `,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("")
		fmt.Printf("\t%v\t%v\t\t\t\t%v\t\t%v\n", "ID", "TASK", "STATUS", "CREATED AT")
		fmt.Println("----------------------------------------------------------------------------------------------------")

		myTasks, _ := models.ReadFile()

		for i, todo := range myTasks {
			status := "Undone"
			if todo.Done {
				status = "Done"
			}
			fmt.Printf("\t%v\t%v\t\t\t\t%v\t\t%v-%v-%v\n ", i+1, todo.Task, status, todo.CreatedAt.Day(), todo.CreatedAt.Month(), todo.CreatedAt.Year())
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdListCmd)

}
