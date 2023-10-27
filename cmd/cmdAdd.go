package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/DanyZugz/zas/models"
	"github.com/spf13/cobra"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var cmdAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Agregar una tarea a la lista.",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("\nInserta una nueva tarea")
		fmt.Print("-> ")

		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		str = strings.Replace(str, "\r\n", "", 1)

		myTasks, _ := models.ReadFile()

		newTask := models.ToDoList{
			Task:      str,
			Done:      false,
			CreatedAt: time.Now(),
		}
		myTasks = append(myTasks, newTask)
		models.SaveFile(myTasks)

		fmt.Println("\nLa tarea ha sido agregada correctamente!")

	},
}

func init() {
	rootCmd.AddCommand(cmdAddCmd)
}
