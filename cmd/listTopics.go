/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

// listTopicsCmd represents the listTopics command
var listTopicsCmd = &cobra.Command{
	Use:   "list-topics",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: listTopics,

	PreRun: func(cmd *cobra.Command, args []string) {
		projectId, _ = cmd.InheritedFlags().GetString("project-id")
	},
}

func init() {
	RootCmd.AddCommand(listTopicsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listTopicsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listTopicsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listTopics(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "pubsub.NewClient: %v", err)
		return
	}
	defer client.Close()

	var topics []*pubsub.Topic

	it := client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Next: %v", err)
			return
		}
		topics = append(topics, topic)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Topics: %s", topics)
}
