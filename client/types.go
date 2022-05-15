package client

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

type ActivityListResponse struct {
	Data []Data `json:"data"`
}
type ActivityResponse struct {
	Data     Data       `json:"data"`
	Included []Included `json:"included"`
}
type Properties struct {
}
type ActivityType struct {
	Data Data `json:"data"`
}
type User struct {
	Data Data `json:"data"`
}
type Relationships struct {
	ActivityType ActivityType `json:"activity_type"`
	Member       Member       `json:"member"`
	User         User         `json:"user"`
	Identities   Identities   `json:"identities"`
}
type Data struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type Attributes struct {
	ActivitiesCount         int           `json:"activities_count"`
	ActivitiesScore         int           `json:"activities_score"`
	AvatarURL               string        `json:"avatar_url"`
	Bio                     interface{}   `json:"bio"`
	Birthday                string        `json:"birthday"`
	Company                 string        `json:"company"`
	Title                   string        `json:"title"`
	CreatedAt               time.Time     `json:"created_at"`
	DeletedAt               interface{}   `json:"deleted_at"`
	FirstActivityOccurredAt time.Time     `json:"first_activity_occurred_at"`
	LastActivityOccurredAt  time.Time     `json:"last_activity_occurred_at"`
	Location                string        `json:"location"`
	Name                    string        `json:"name"`
	Pronouns                string        `json:"pronouns"`
	Reach                   int           `json:"reach"`
	ShippingAddress         string        `json:"shipping_address"`
	Slug                    string        `json:"slug"`
	Source                  string        `json:"source"`
	TagList                 []interface{} `json:"tag_list"`
	Tags                    []interface{} `json:"tags"`
	Teammate                bool          `json:"teammate"`
	Tshirt                  string        `json:"tshirt"`
	UpdatedAt               time.Time     `json:"updated_at"`
	MergedAt                interface{}   `json:"merged_at"`
	URL                     string        `json:"url"`
	OrbitURL                string        `json:"orbit_url"`
	Created                 bool          `json:"created"`
	ID                      string        `json:"id"`
	OrbitLevel              interface{}   `json:"orbit_level"`
	Love                    string        `json:"love"`
	Twitter                 string        `json:"twitter"`
	Github                  string        `json:"github"`
	Discourse               interface{}   `json:"discourse"`
	Email                   string        `json:"email"`
	Devto                   interface{}   `json:"devto"`
	Linkedin                string        `json:"linkedin"`
	Discord                 interface{}   `json:"discord"`
	GithubFollowers         int           `json:"github_followers"`
	TwitterFollowers        int           `json:"twitter_followers"`
	Topics                  []string      `json:"topics"`
	Languages               []string      `json:"languages"`
}
type Identities struct {
	Data []Data `json:"data"`
}
type Included struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    Attributes    `json:"attributes,omitempty"`
	Relationships Relationships `json:"relationships,omitempty"`
}

type Activity struct {
	Title        string `json:"title"`
	URL          string `json:"url"`
	OccurredAt   string `json:"occurred_at"`
	ActivityType string `json:"activity_type"`
	Weight       string `json:"weight"`
	Key          string `json:"key"`
	Member       Member `json:"member"`
}

type Identity struct {
	Source string
}

type ActivityPayload struct {
	Activity Activity `json:"activity"`
	Identity Identity `json:"identity"`
}

type Member struct {
	GitHub  string `json:"github"`
	Twitter string `json:"twitter"`
}

func (p *ActivityPayload) GetPayload() io.Reader {
	data, err := json.Marshal(p)
	if err == nil {
		return strings.NewReader(string(data))
	}
	return nil
}
