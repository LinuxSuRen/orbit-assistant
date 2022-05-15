package cmd

import (
	"fmt"
	"github.com/linuxsuren/orbit-assistant/client"
	"github.com/spf13/cobra"
	"os"
)

type activityListOption struct {
	commonOption
	member string
	id     string
}

func newActivityListCommand() (cmd *cobra.Command) {
	opt := &activityListOption{}

	cmd = &cobra.Command{
		Use:  "ls",
		RunE: opt.runE,
	}

	flags := cmd.Flags()
	flags.StringVarP(&opt.id, "id", "", "", "")
	flags.StringVarP(&opt.member, "member", "m", "", "")
	flags.StringVarP(&opt.workspace, "workspace", "w", "", "")

	_ = cmd.MarkFlagRequired("workspace")
	return
}

func (o *activityListOption) runE(cmd *cobra.Command, args []string) (err error) {
	token := os.Getenv("ORBIT_TOKEN")
	orbit := client.NewOrbit(token)
	if o.member != "" {
		result := orbit.GetActivityListByMember(o.workspace, o.member)
		for _, item := range result.Data {
			fmt.Println(item.Type)
		}
	} else if o.id != "" {
		result := orbit.GetActivityByID(o.workspace, o.id)
		fmt.Println(result)
	}
	return
}
