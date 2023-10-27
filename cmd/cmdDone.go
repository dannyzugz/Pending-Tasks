package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DanyZugz/zas/models"
	"github.com/spf13/cobra"
)

var cmdDoneCmd = &cobra.Command{
	Use:   "done [command]",
	Short: "Marcar una tarea como completada.",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("\n\tDebes proporcionar un Id para cambiar el estado de la tarea. Ej: zas done 2")
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

		myTasks[index-1].Done = true
		fmt.Printf("El estado de la tarea '%s' con Id = %d fue actualizado correctamente!\n", myTasks[index-1].Task, index)

		models.SaveFile(myTasks)
	},
}

func init() {
	rootCmd.AddCommand(cmdDoneCmd)
}
