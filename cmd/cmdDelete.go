package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DanyZugz/zas/models"
	"github.com/spf13/cobra"
)

// cmdDeleteCmd represents the cmdDelete command
var cmdDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Eliminar una tarea de la lista.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("")
		myTasks, _ := models.ReadFile()
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}
		if index > len(myTasks) {
			fmt.Println(" El indice indicado no existe en la lista")
			os.Exit(1)
		}
		// var newTasks []models.ToDoList
		fmt.Printf("La tarea '%s' con Id = %d ha sido eliminada correctamente!\n", myTasks[index-1].Task, index)
		myTasks = append(myTasks[:index-1], myTasks[index:]...)

		models.SaveFile(myTasks)
	},
}

func init() {
	rootCmd.AddCommand(cmdDeleteCmd)
}
