package cmd

import (
	"context"
	"github.com/A-Zee029/rocketmq-client-go/v2/admin"
	"github.com/A-Zee029/rocketmq-client-go/v2/primitive"
	"github.com/spf13/cobra"
	"log"
)

// deleteTopicCmd represents the deleteTopic command
var deleteTopicCmd = &cobra.Command{
	Use:   "deleteTopic",
	Short: "delete a topic",
	Long:  "delete a topic",
	Run: func(cmd *cobra.Command, args []string) {
		clusterName, err := cmd.Flags().GetString("clusterName")
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		topic, err := cmd.Flags().GetString("topic")
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		if len(nameSrvAddrs) == 0 {
			log.Fatalln("Fail to parse nameSrvAddrs")
			return
		}

		adminInstance, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddrs)))
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
	},
}

func init() {
	rootCmd.AddCommand(deleteTopicCmd)

	deleteTopicCmd.Flags().StringP("clusterName", "c", "", "cluster to delete topic")
	deleteTopicCmd.Flags().StringP("topic", "t", "", "name of topic")
}
