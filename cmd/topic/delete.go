/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package topic

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

// deleteTopicCmd represents the deleteTopic command
var deleteTopicCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: deleteTopic,
}

func init() {
	TopicCmd.AddCommand(deleteTopicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteTopicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteTopicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteTopic(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	projectId, _ = cmd.InheritedFlags().GetString("project-id")
	fmt.Fprintf(cmd.OutOrStdout(), "Project Id: %s\n", projectId)
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "Failed to create pubsub client %v", err)
		return
	}

	topic := client.Topic(topicName)
	if err := topic.Delete(ctx); err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "Failed to delete topic %s %v", topicName, err)
		return
	}
}
