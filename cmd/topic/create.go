/*
Copyright © 2022 Gabriel PASCUAL <gabriel.pascual@gmail.com>
*/
package topic

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

// createTopicCmd represents the createTopic command
var createTopicCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: createTopic,
}

func init() {
	TopicCmd.AddCommand(createTopicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createTopicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTopicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createTopic(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	projectId, _ = cmd.InheritedFlags().GetString("project-id")
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		// TODO: Handle error.
	}

	// Create a new topic with the given name.
	topic, err := client.CreateTopic(ctx, topicName)
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "Impossible to create the topic %s %v", topicName, err)
		return
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Topic has been created %s", topic)
}
