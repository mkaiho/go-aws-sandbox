package main

import (
	"os"
	"sync"
	"time"

	"github.com/mkaiho/go-ecs-batch-sample/logging"
	"github.com/spf13/cobra"
)

var (
	command *cobra.Command
)

func init() {
	logging.InitLoggerWithZap()
	command = newCommand()
}

func main() {
	if err := command.Execute(); err != nil {
		logging.GetLogger().Error(err, "error has occured")
		os.Exit(1)
	}
}

func newCommand() *cobra.Command {
	command := cobra.Command{
		Use:   "echo args...",
		Short: "display args",
		Long:  "Display arguments on stdout.",
		RunE:  handle,
	}

	return &command
}

func handle(cmd *cobra.Command, args []string) error {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() error {
			time.Sleep(time.Millisecond * 2000)
			defer wg.Done()
			return r()
		}()
	}
	wg.Wait()

	logging.GetLogger().Info("Done")
	return nil
}

func r() error {
	f, err := os.Open("abc")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString("hello world")
	return err
}
