/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var task string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task",
	Long:  "add task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(task) == 0 {
			return errors.New("taskname is empty")
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot get user home dir path: %s", err.Error())
		}

		dir := fmt.Sprintf("%s/desktop/go/cli/.todo", home)
		if s, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.Mkdir(dir, 0777); err != nil {
				return fmt.Errorf("cannot create directory: %s", err.Error())
			} else if !s.IsDir() {
				return fmt.Errorf("%s is not directory", dir)
			}
		}

		f := fmt.Sprintf("%s/task.csv", dir)
		fp, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("cannot create or open file: %s", err.Error())
		}
		defer fp.Close()

		w := csv.NewWriter(fp)
		w.Write([]string{task})
		w.Flush()
		if err := w.Error(); err != nil {
			return err
		}

		fmt.Println("Success!!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&task, "task", "", "task name")
	addCmd.MarkFlagRequired("task")
}
