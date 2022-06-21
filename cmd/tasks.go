package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/labstack/echo/v4"
)

func init() {
	rootCmd.AddCommand(tasksCmd)
}

var tasksCmd = &cobra.Command{
	Use:   "rest-server",
	Short: "Starting rest server",
	Run: func(cmd *cobra.Command, args []string) {
		srv := asynq.NewServer(
			asynq.RedisClientOpt{Addr: ":6379"},
			asynq.Config{
				// Specify how many concurrent workers to use
				Concurrency: 10,
				// Optionally specify multiple queues with different priority.
				Queues: map[string]int{
					"endpoint": 6,
				},
				// See the godoc for other configuration options
			},
		)

		// mux maps a type to a handler
		mux := asynq.NewServeMux()
		mux.HandleFunc(tasks.ImageProcessor)
		mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor())
		// ...register other handlers...

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}

