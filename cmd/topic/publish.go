/*
Copyright © 2022 Gabriel PASCUAL <gabriel.pascual@gmail.com>
*/
package topic

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"

	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: publish,
}

var message string

func init() {
	TopicCmd.AddCommand(publishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")
	publishCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "message content")
	
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func publish(cmd *cobra.Command, args []string) {
	projectId, _ = cmd.InheritedFlags().GetString("project-id")
	fmt.Println("Publisher")
	fmt.Fprintf(cmd.OutOrStdout(), "Message %s\n", message)
	fmt.Fprintf(cmd.OutOrStdout(), "Topic Id:%s\n",topicName)
	fmt.Fprintf(cmd.OutOrStdout(), "Project Id:%s\n",projectId)
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "pubsub.NewClient: %v", err)
		return
	}
	defer client.Close()

	var wg sync.WaitGroup
	t := client.Topic(topicName)
	fmt.Fprintf(cmd.OutOrStdout(), "Topic Name:%s\n",t.ID())

	messageBytes := json.RawMessage(message)
	
	result := t.Publish(ctx, &pubsub.Message{
		Data: messageBytes,
	})

	wg.Add(1)
	go func(res *pubsub.PublishResult) {
		defer wg.Done()
		// The Get method blocks until a server-generated ID or
		// an error is returned for the published message.
		id, err := res.Get(ctx)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Failed to publish: %v\n", err)
			return
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Published message; msg ID: %v\n", id)

	}(result)

	wg.Wait()
}
