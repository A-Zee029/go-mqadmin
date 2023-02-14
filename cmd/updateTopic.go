package cmd

import (
	"context"
	"github.com/A-Zee029/rocketmq-client-go/v2/admin"
	"github.com/A-Zee029/rocketmq-client-go/v2/primitive"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// updateTopicCmd represents the updateTopic command
var updateTopicCmd = &cobra.Command{
	Use:   "updateTopic",
	Short: "Create or update a topic",
	Long:  "Create or update a topic",
	Run: func(cmd *cobra.Command, args []string) {
		nameSrvAddrs, err := cmd.Flags().GetString("nameSrvAddrs")
		brokerAddr, err := cmd.Flags().GetString("brokerAddr")
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

			err = adminInstance.CreateTopic(
				context.Background(),
				admin.WithTopicCreate(topic),
				admin.WithBrokerAddrCreate(brokerAddr),
				admin.WithClusterNameCreate(clusterName),
			)
			if err != nil {
				log.Fatalln("Create topic error:", err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateTopicCmd)

	topicListCmd.Flags().StringP("nameSrvAddrs", "n", "", "NameServer endpoint in format ip:port")
	topicListCmd.Flags().StringP("brokerAddr", "b", "", "Broker endpoint in format ip:port")
	topicListCmd.Flags().StringP("clusterName", "c", "", "cluster to create or update topic")
	topicListCmd.Flags().StringP("topic", "t", "", "name of topic")
}
