package cmd

import (
	"context"
	"fmt"
	"github.com/A-Zee029/rocketmq-client-go/v2/admin"
	"github.com/A-Zee029/rocketmq-client-go/v2/primitive"
	"github.com/spf13/cobra"
	"log"
)

// topicListCmd represents the topicList command
var topicListCmd = &cobra.Command{
	Use:   "topicList",
	Short: "Fetch all topic",
	Long:  "Fetch all topic",
	Run: func(cmd *cobra.Command, args []string) {

		if len(nameSrvAddrs) == 0 {
			log.Fatalln("Fail to parse nameSrvAddrs")
			return
		}

		adminInstance, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddrs)))
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		topicList, err := adminInstance.FetchAllTopicList(context.Background())
		if err != nil {
			log.Fatalln("List topic error:", err.Error())
		}
		for _, topicName := range topicList.TopicNameList {
			fmt.Println(topicName)
		}
	},
}

func init() {
	rootCmd.AddCommand(topicListCmd)
}
