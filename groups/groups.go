package groups

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	h "net/http"
	"os"

	"github.com/HybriStratus/test-github-groups/http"
)

const baseURL = "https://api.github.com/orgs"
const TestOrg = "HybriStratus"
const DefaultRoleType = "member"

type Team struct {
	Name         string   `json:"name"`
	Description  string   `json:"description,omitempty"`
	Maintainers  []string `json:"maintainers,omitempty"`
	Repos        []string `json:"repo_names,omitempty"`
	Privacy      string   `json:"privacy,omitempty"`
	ParentTeamID int      `json:"parent_team_id,omitempty"`
}

func sendHTTPRequest(client http.Client, method string, url string, body io.Reader) (response *h.Response, err error) {
	req, err := h.NewRequest(method, url, body)
	if err != nil {
		err = fmt.Errorf("Error occurred while creating http request " + err.Error())
		return
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("AUTH_TOKEN"))
	req.Header.Set("Content-Type", "application/vnd.github.v3+json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	// Make the API call
	response, err = client.Do(req)
	if err != nil {
		err = fmt.Errorf("Error occurred while calling github API: " + err.Error())
		return
	}
	return
}

// CreateTeam creates team in Github
func (team *Team) CreateTeam(client http.Client) (err error) {

	url := fmt.Sprintf("%s/%s/teams", baseURL, TestOrg)
	// Convert the json body object to bytes
	jsonValue, err := json.Marshal(team)
	if err != nil {
		return fmt.Errorf("Error in marshalling the request payload")
	}

	response, err := sendHTTPRequest(client, "POST", url, bytes.NewBuffer(jsonValue))
	fmt.Printf("Return status code of the request: %d\n", response.StatusCode)
	if err != nil {
		return
	}
	if response.StatusCode != h.StatusCreated {
		err = fmt.Errorf("Error in creating a new team : %s", team.Name)
		return
	}
	defer response.Body.Close()
	// Read the bytes from the response body
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Error in reading response from API response%s", err.Error())
		return
	}

	var teamResponse interface{}
	// Convert the bytes in interface
	err = json.Unmarshal(responseBodyBytes, &teamResponse)
	if err != nil {
		err = fmt.Errorf("Error in unmarshalling response from API response %s", err.Error())
		return
	}
	fmt.Printf("Response: %v\n\n", teamResponse)
	return
}

// GetTeamDetails gets GitHub team details
func (team *Team) GetTeamDetails(client http.Client) (err error) {

	url := fmt.Sprintf("%s/%s/teams/%s", baseURL, TestOrg, team.Name)

	response, err := sendHTTPRequest(client, "GET", url, nil)
	fmt.Printf("Return status code of the request: %d\n", response.StatusCode)
	if err != nil {
		return
	}

	if response.StatusCode != h.StatusOK {
		err = fmt.Errorf("Error in getting team deatils : %s", team.Name)
		return
	}
	defer response.Body.Close()
	// Read the bytes from the response body
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Error in reading response from API response%s", err.Error())
		return
	}

	var teamResponse interface{}
	// Convert the bytes in interface
	err = json.Unmarshal(responseBodyBytes, &teamResponse)
	if err != nil {
		err = fmt.Errorf("Error in unmarshalling response from API response %s", err.Error())
		return
	}
	fmt.Printf("Response: %v\n\n", teamResponse)
	return
}

// UpdateTeam updates GitHub team
func (team *Team) UpdateTeam(client http.Client) (err error) {

	url := fmt.Sprintf("%s/%s/teams/%s", baseURL, TestOrg, team.Name)
	// Convert the json body object to bytes
	jsonValue, err := json.Marshal(team)
	if err != nil {
		return fmt.Errorf("Error in marshalling the request payload")
	}
	response, err := sendHTTPRequest(client, "PATCH", url, bytes.NewBuffer(jsonValue))
	fmt.Printf("Return status code of the request: %d\n", response.StatusCode)
	if err != nil {
		return
	}

	if response.StatusCode != h.StatusOK {
		err = fmt.Errorf("Error in updating team : %s", team.Name)
		return
	}
	defer response.Body.Close()
	// Read the bytes from the response body
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Error in reading response from API response%s", err.Error())
		return
	}

	var teamResponse interface{}
	// Convert the bytes in interface
	err = json.Unmarshal(responseBodyBytes, &teamResponse)
	if err != nil {
		err = fmt.Errorf("Error in unmarshalling response from API response %s", err.Error())
		return
	}
	fmt.Printf("Response: %v\n\n", teamResponse)
	return
}

// DeleteTeam deletes GitHub team
func (team *Team) DeleteTeam(client http.Client) (err error) {
	url := fmt.Sprintf("%s/%s/teams/%s", baseURL, TestOrg, team.Name)

	response, err := sendHTTPRequest(client, "DELETE", url, nil)
	fmt.Printf("Return status code of the request: %d\n\n", response.StatusCode)
	if err != nil {
		return
	}

	if response.StatusCode != h.StatusNoContent {
		err = fmt.Errorf("Error in deleting team : %s", team.Name)
		return
	}

	return
}

// ListMemebersOfTeam gets all memebers part of the Github team
func (team *Team) ListMemebersOfTeam(client http.Client) (err error) {

	url := fmt.Sprintf("%s/%s/teams/%s/members", baseURL, TestOrg, team.Name)

	response, err := sendHTTPRequest(client, "GET", url, nil)
	fmt.Printf("Return status code of the request: %d\n\n", response.StatusCode)
	if err != nil {
		return
	}

	if response.StatusCode != h.StatusOK {
		err = fmt.Errorf("Error in getting members of a team : %s", team.Name)
		return
	}
	defer response.Body.Close()
	// Read the bytes from the response body
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Error in reading response from API response %s", err.Error())
		return
	}

	var teamResponse []interface{}
	// Convert the bytes in interface
	err = json.Unmarshal(responseBodyBytes, &teamResponse)
	if err != nil {
		err = fmt.Errorf("Error in unmarshalling response from API response %s", err.Error())
		return
	}
	fmt.Printf("Response: %v\n\n", teamResponse)
	return
}

// AddMemeberToTeam adds memeber to a GitHub team
func AddMemeberToTeam(client http.Client, teamName, userName, roleType string) (err error) {

	url := fmt.Sprintf("%s/%s/teams/%s/memberships/%s", baseURL, TestOrg, teamName, userName)

	type memeberRole struct {
		Role string `json:"role,omitempty"`
	}
	memRole := memeberRole{}
	if roleType == "" {
		memRole.Role = DefaultRoleType
	} else {
		memRole.Role = roleType
	}

	// Convert the json body object to bytes
	jsonValue, err := json.Marshal(memRole)
	if err != nil {
		return fmt.Errorf("Error in marshalling the request payload")
	}

	response, err := sendHTTPRequest(client, "PUT", url, bytes.NewBuffer(jsonValue))
	fmt.Printf("Return status code of the request: %d\n", response.StatusCode)
	if err != nil {
		return
	}
	if response.StatusCode != h.StatusOK {
		err = fmt.Errorf("Error in adding %s to team %s", userName, teamName)
		return
	}
	defer response.Body.Close()
	// Read the bytes from the response body
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Error in reading response from API response%s", err.Error())
		return
	}

	var teamResponse interface{}
	// Convert the bytes in interface
	err = json.Unmarshal(responseBodyBytes, &teamResponse)
	if err != nil {
		err = fmt.Errorf("Error in unmarshalling response from API response %s", err.Error())
		return
	}
	fmt.Printf("Response: %v\n\n", teamResponse)
	return
}

// DeleteMemberFromTeam deletes memeber from GitHub team
func DeleteMemberFromTeam(client http.Client, teamName, userName string) (err error) {

	url := fmt.Sprintf("%s/%s/teams/%s/memberships/%s", baseURL, TestOrg, teamName, userName)

	response, err := sendHTTPRequest(client, "DELETE", url, nil)
	fmt.Printf("Return status code of the request: %d\n", response.StatusCode)
	if err != nil {
		return
	}

	if response.StatusCode != h.StatusNoContent {
		err = fmt.Errorf("Error in deleting %s from team %s", userName, teamName)
		return
	}

	return
}
