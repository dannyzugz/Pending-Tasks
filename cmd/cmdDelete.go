package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DanyZugz/zas/models"
	"github.com/spf13/cobra"
)

var cmdDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Eliminar una tarea de la lista.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("\n\tDebes proporcionar un Id para eliminar la tarea. Ej: zas delete 2")
			os.Exit(1)
		}
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
		fmt.Printf("La tarea '%s' con Id = %d ha sido eliminada correctamente!\n", myTasks[index-1].Task, index)
		myTasks = append(myTasks[:index-1], myTasks[index:]...)

		models.SaveFile(myTasks)
	},
}

func init() {
	rootCmd.AddCommand(cmdDeleteCmd)
}
