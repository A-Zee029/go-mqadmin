package cmd

import (
	"context"
	"github.com/A-Zee029/rocketmq-client-go/v2/admin"
	"github.com/A-Zee029/rocketmq-client-go/v2/primitive"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// deleteTopicCmd represents the deleteTopic command
var deleteTopicCmd = &cobra.Command{
	Use:   "deleteTopic",
	Short: "delete a topic",
	Long:  "delete a topic",
	Run: func(cmd *cobra.Command, args []string) {
		nameSrvAddrs, err := cmd.Flags().GetString("nameSrvAddrs")
		clusterName, err := cmd.Flags().GetString("clusterName")
		topic, err := cmd.Flags().GetString("topic")

		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		if nameSrvAddrs != "" {
			nameSrvAddr := strings.Split(nameSrvAddrs, ";")
			adminInstance, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
			if err != nil {
				log.Fatalln(err.Error())
				return
			}

			err = adminInstance.DeleteTopic(
				context.Background(),
				admin.WithTopicDelete(topic),
				admin.WithClusterNameDelete(clusterName),
			)
			if err != nil {
				log.Fatalln("Delete topic error:", err.Error())
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(deleteTopicCmd)

	topicListCmd.Flags().StringP("nameSrvAddrs", "n", "", "NameServer endpoint in format ip:port")
	topicListCmd.Flags().StringP("clusterName", "c", "", "cluster to create or update topic")
	topicListCmd.Flags().StringP("topic", "t", "", "name of topic")
}
