package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/linuxsuren/orbit-assistant/client"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type activityUpdateOption struct {
	commonOption
	member string
	id     string
	file   string
}

func newActivityUpdateCommand() (cmd *cobra.Command) {
	opt := &activityUpdateOption{}
	cmd = &cobra.Command{
		Use:  "update",
		RunE: opt.runE,
	}
	flags := cmd.Flags()
	flags.StringVarP(&opt.file, "file", "f", "", "")
	return
}

type ActivityRecord struct {
	Title  string `json:"title" yaml:"title"`
	GitHub string `json:"github" yaml:"github"`
	Date   string `json:"date" yaml:"date"`
	Type   string `json:"type" yaml:"type"`
}

func (r ActivityRecord) GetID() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s/%s", r.GitHub, r.Date)))
}

type ActivityRecords struct {
	Workspace string
	Records   []ActivityRecord `yaml:"records"`
}

func (r ActivityRecords) List() []ActivityRecord {
	return r.Records
}

func (o *activityUpdateOption) runE(cmd *cobra.Command, args []string) (err error) {
	if o.file == "" {
		err = fmt.Errorf("not support without file")
		return
	}

	var data []byte
	if data, err = ioutil.ReadFile(o.file); err != nil {
		return
	}

	records := &ActivityRecords{}
	if err = yaml.Unmarshal(data, records); err != nil {
		return
	}

	token := os.Getenv("ORBIT_TOKEN")
	orbit := client.NewOrbit(token)
	for i := range records.List() {
		record := records.List()[i]

		activity := orbit.GetActivityByID(records.Workspace, record.GetID())
		if activity != nil {
			continue
		}

		activityPayload := &client.ActivityPayload{
			Activity: client.Activity{
				Title:        record.Title,
				URL:          "http://github.com",
				ActivityType: record.Type,
				OccurredAt:   record.Date,
				Member: client.Member{
					GitHub: record.GitHub,
				},
				Key: record.GetID(),
			},
			Identity: client.Identity{
				Source: record.GitHub,
			},
		}
		orbit.CreateActivity(records.Workspace, activityPayload)
	}
	return
}
