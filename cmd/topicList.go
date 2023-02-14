package cmd

import (
	"context"
	"fmt"
	"github.com/A-Zee029/rocketmq-client-go/v2/admin"
	"github.com/A-Zee029/rocketmq-client-go/v2/primitive"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// topicListCmd represents the topicList command
var topicListCmd = &cobra.Command{
	Use:   "topicList",
	Short: "Fetch all topic",
	Long:  "Fetch all topic",
	Run: func(cmd *cobra.Command, args []string) {
		nameSrvAddrs, err := cmd.Flags().GetString("nameSrvAddrs")
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
			topicList, err := adminInstance.FetchAllTopicList(context.Background())
			if err != nil {
				log.Fatalln("List topic error:", err.Error())
			}
			for _, topicName := range topicList.TopicNameList {
				fmt.Println(topicName)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(topicListCmd)

	topicListCmd.Flags().StringP("nameSrvAddrs", "n", "127.0.0.1:9876", "NameServer endpoint in format ip:port")

}
