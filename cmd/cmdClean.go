/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/DanyZugz/zas/models"
	"github.com/spf13/cobra"
)

// cmdCleanCmd represents the cmdClean command
var cmdCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Borrar todas las tareas de la lista.",
	Long:  `Borra todas las tareas de la lista de tareas actual de forma definitiva. Para borrar solo una tarea use el comando "delete"`,
	Run: func(cmd *cobra.Command, args []string) {

		err := models.Clean()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(" Se eliminó la lista de tareas correctamente! ")

	},
}

func init() {
	rootCmd.AddCommand(cmdCleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmdCleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmdCleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
