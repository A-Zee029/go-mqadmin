package cmd

import (
	"context"
	"github.com/A-Zee029/rocketmq-client-go/v2/admin"
	"github.com/A-Zee029/rocketmq-client-go/v2/primitive"
	"github.com/spf13/cobra"
	"log"
)

// updateTopicCmd represents the updateTopic command
var updateTopicCmd = &cobra.Command{
	Use:   "updateTopic",
	Short: "Create or update a topic",
	Long:  "Create or update a topic",
	Run: func(cmd *cobra.Command, args []string) {

		brokerAddr, err := cmd.Flags().GetString("brokerAddr")
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

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

		err = adminInstance.CreateTopic(
			context.Background(),
			admin.WithTopicCreate(topic),
			admin.WithBrokerAddrCreate(brokerAddr),
			admin.WithClusterNameCreate(clusterName),
		)
		if err != nil {
			log.Fatalln("Create topic error:", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(updateTopicCmd)

	updateTopicCmd.Flags().StringP("brokerAddr", "b", "", "Broker endpoint in format ip:port")
	updateTopicCmd.Flags().StringP("clusterName", "c", "", "cluster to create or update topic")
	updateTopicCmd.Flags().StringP("topic", "t", "", "name of topic")
}
