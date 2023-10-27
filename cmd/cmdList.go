package cmd

import (
	"fmt"

	"github.com/DanyZugz/zas/models"
	"github.com/spf13/cobra"
)

var cmdListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas las tareas en la lista.",
	Long:  ` `,
	Run: func(cmd *cobra.Command, args []string) {

		myTasks, _ := models.ReadFile()

		if len(myTasks) != 0 {
			fmt.Println("")
			fmt.Printf("\t%v\t%v\t\t\t\t%v\t\t%v\n", "ID", "TASK", "STATUS", "CREATED AT")
			fmt.Println("----------------------------------------------------------------------------------------------------")

			for i, todo := range myTasks {
				status := "Undone"
				if todo.Done {
					status = "Done"
				}
				fmt.Printf("\t%v\t%v\t\t\t\t%v\t\t%v-%v-%v\n ", i+1, todo.Task, status, todo.CreatedAt.Day(), todo.CreatedAt.Month(), todo.CreatedAt.Year())
			}
		} else {
			fmt.Println("\n\t Su lista de tareas está vacía!")
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdListCmd)

}
