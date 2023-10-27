/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DanyZugz/zas/models"
	"github.com/spf13/cobra"
)

// cmdDoneCmd represents the cmdDone command
var cmdDoneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marcar una tarea como completada.",

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

		myTasks[index-1].Done = true
		fmt.Printf("El estado de la tarea '%s' con Id = %d fue actualizado correctamente!\n", myTasks[index-1].Task, index)

		models.SaveFile(myTasks)
	},
}

func init() {
	rootCmd.AddCommand(cmdDoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmdDoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmdDoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
