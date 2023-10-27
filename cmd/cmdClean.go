package cmd

import (
	"fmt"

	"github.com/DanyZugz/zas/models"
	"github.com/spf13/cobra"
)

var cmdCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Borrar todas las tareas de la lista.",
	Long:  `Borra todas las tareas de la lista de tareas actual de forma definitiva. Para borrar solo una tarea use el comando "delete"`,
	Run: func(cmd *cobra.Command, args []string) {

		err := models.Clean()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\n\tSe eliminó la lista de tareas correctamente! ")

	},
}

func init() {
	rootCmd.AddCommand(cmdCleanCmd)
}
