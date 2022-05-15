package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Orbit struct {
	token string
}

func NewOrbit(token string) *Orbit {
	return &Orbit{
		token: token,
	}
}

func (o *Orbit) CreateActivity(workspace string, activityPayload *ActivityPayload) {
	url := fmt.Sprintf("https://app.orbit.love/api/v1/%s/activities", workspace)

	payload := activityPayload.GetPayload()
	req, _ := http.NewRequest(http.MethodPost, url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	o.setAuthHeader(req)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func (o *Orbit) GetActivityListByMember(workspace, member string) (activityResponse *ActivityListResponse) {
	url := fmt.Sprintf("https://app.orbit.love/api/v1/%s/members/%s/activities", workspace, member)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Accept", "application/json")
	o.setAuthHeader(req)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	activityResponse = &ActivityListResponse{}
	err := json.Unmarshal(body, activityResponse)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (o *Orbit) GetActivityByID(workspace, id string) (activityResponse *ActivityResponse) {
	url := fmt.Sprintf("https://app.orbit.love/api/v1/%s/activities/%s", workspace, id)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Accept", "application/json")
	o.setAuthHeader(req)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	activityResponse = &ActivityResponse{}
	err := json.Unmarshal(body, activityResponse)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (o *Orbit) setAuthHeader(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", o.token))
}
