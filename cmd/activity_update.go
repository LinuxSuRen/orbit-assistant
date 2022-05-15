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
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(
		"%s-%s-%s-%s",
		r.Title, r.GitHub, r.Date, r.Type)))
}

type EventRecord struct {
	Type         string   `yaml:"type"`
	Date         string   `json:"date"`
	Title        string   `json:"title"`
	Link         string   `json:"link"`
	Participants []Member `json:"participants"`
}

func (r EventRecord) GetID(member Member) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(
		"%s-%s-%s-%s-%s",
		r.Type, r.Date, r.Title, r.Link, member.GitHub)))
}

type Member struct {
	GitHub string `json:"github"`
}

type ActivityRecords struct {
	Workspace string
	Records   []ActivityRecord `yaml:"records"`
	Events    []EventRecord    `yaml:"events"`
}

func (r ActivityRecords) ListRecords() []ActivityRecord {
	return r.Records
}

func (r ActivityRecords) ListEvents() []EventRecord {
	return r.Events
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
	for i := range records.ListRecords() {
		record := records.ListRecords()[i]

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
				Source: "orbit-assistant",
			},
		}
		orbit.CreateActivity(records.Workspace, activityPayload)
	}

	for i := range records.ListEvents() {
		event := records.ListEvents()[i]

		for j := range event.Participants {
			participant := event.Participants[j]
			activity := orbit.GetActivityByID(records.Workspace, event.GetID(participant))
			if activity != nil {
				continue
			}

			activityPayload := &client.ActivityPayload{
				Activity: client.Activity{
					Title:        event.Title,
					URL:          event.Link,
					ActivityType: event.Type,
					OccurredAt:   event.Date,
					Member: client.Member{
						GitHub: participant.GitHub,
					},
					Key: event.GetID(participant),
				},
				Identity: client.Identity{
					Source: "orbit-assistant",
				},
			}
			orbit.CreateActivity(records.Workspace, activityPayload)
		}
	}
	return
}
