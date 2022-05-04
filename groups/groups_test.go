package groups

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/HybriStratus/test-github-groups/http/mock"
)

// Test used to test the Mock Client
type responses struct {
	method   string
	url      string
	response http.Response
}

// TestCreateTeam tests CreateTeam function of a team
func TestCreateTeam(t *testing.T) {

	createResponse := fmt.Sprintf(`{
    "name": "test_team",
    "id": 5714710,
    "node_id": "T_kwDOBc9UpM4AVzMW",
    "slug": "test_team",
    "description": "test_team",
    "privacy": "closed",
    "url": "https://api.github.com/organizations/97473700/team/5714710",
    "permission": "pull",
    "members_count": 2,
    "repos_count": 0,
    "organization": {
        "login": "HybriStratus",
        "id": 97473700,
        "node_id": "O_kgDOBc9UpA",
        "url": "https://api.github.com/orgs/HybriStratus",
        "repos_url": "https://api.github.com/orgs/HybriStratus/repos",
        "description": null,
        "email": null,
        "type": "Organization"
    },
    "parent": null
}`)
	newTeam := Team{
		Name:        "test_team",
		Description: "test_team",
		Maintainers: []string{"testuser"},
		Privacy:     "closed",
	}
	// Create your table test
	tests := []struct {
		name           string
		requestClients []responses
		payload        Team
		expected       error
	}{
		{
			name:     "Testing successful team creation",
			payload:  newTeam,
			expected: nil,
			requestClients: []responses{
				{
					method: http.MethodPost,
					url:    fmt.Sprintf("%s/%s/teams", baseURL, TestOrg),
					response: http.Response{
						StatusCode: http.StatusCreated,
						Body:       ConvertBytesToIoReadCloser([]byte(createResponse)),
					},
				},
			},
		},
		{
			name:     "Testing failure of team creation",
			payload:  newTeam,
			expected: fmt.Errorf("Error in creating a new team : test_team"),
			requestClients: []responses{
				{
					method: http.MethodPost,
					url:    fmt.Sprintf("%s/%s/teams", baseURL, TestOrg),
					response: http.Response{
						StatusCode: http.StatusBadGateway,
					},
				},
			},
		},
	}
	// Go through each of the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create the mock Client
			mockClient := mock.Client{}
			for _, response := range tt.requestClients {
				mockClient.SetResponses(response.method, response.url, response.response)
			}

			got := tt.payload.CreateTeam(mockClient)
			if got != nil && tt.expected != nil {
				if got.Error() != tt.expected.Error() {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			} else {
				if got != tt.expected {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			}
		})
	}
}

//TestGetTeamDetails tests the GetTeamDetails of a team
func TestGetTeamDetails(t *testing.T) {

	createResponse := fmt.Sprintf(`{
    "name": "test_team",
    "id": 5714710,
    "node_id": "T_kwDOBc9UpM4AVzMW",
    "slug": "test_team",
    "description": "test_team",
    "privacy": "closed",
    "url": "https://api.github.com/organizations/97473700/team/5714710",
    "permission": "pull",
    "members_count": 2,
    "repos_count": 0,
    "organization": {
        "login": "HybriStratus",
        "id": 97473700,
        "node_id": "O_kgDOBc9UpA",
        "url": "https://api.github.com/orgs/HybriStratus",
        "repos_url": "https://api.github.com/orgs/HybriStratus/repos",
        "description": null,
        "email": null,
        "type": "Organization"
    },
    "parent": null
}`)
	newTeam := Team{
		Name: "test_team",
	}
	// Create your table test
	tests := []struct {
		name           string
		requestClients []responses
		payload        Team
		expected       error
	}{
		{
			name:     "Testing successful retrival of team details",
			payload:  newTeam,
			expected: nil,
			requestClients: []responses{
				{
					method: http.MethodGet,
					url:    fmt.Sprintf("%s/%s/teams/%s", baseURL, TestOrg, newTeam.Name),
					response: http.Response{
						StatusCode: http.StatusOK,
						Body:       ConvertBytesToIoReadCloser([]byte(createResponse)),
					},
				},
			},
		},
		{
			name:     "Testing failure of team details retrival",
			payload:  newTeam,
			expected: fmt.Errorf("Error in getting team deatils : test_team"),
			requestClients: []responses{
				{
					method: http.MethodGet,
					url:    fmt.Sprintf("%s/%s/teams/%s", baseURL, TestOrg, newTeam.Name),
					response: http.Response{
						StatusCode: http.StatusConflict,
					},
				},
			},
		},
	}
	// Go through each of the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create the mock Client
			mockClient := mock.Client{}
			for _, response := range tt.requestClients {
				mockClient.SetResponses(response.method, response.url, response.response)
			}

			got := tt.payload.GetTeamDetails(mockClient)
			if got != nil && tt.expected != nil {
				if got.Error() != tt.expected.Error() {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			} else {
				if got != tt.expected {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			}
		})
	}
}

// TestUpdateTeam tests UpdateTeam function for a team
func TestUpdateTeam(t *testing.T) {

	createResponse := fmt.Sprintf(`{
    "name": "test_team",
    "id": 5714710,
    "node_id": "T_kwDOBc9UpM4AVzMW",
    "slug": "test_team",
    "description": "test_team",
    "privacy": "secret",
    "url": "https://api.github.com/organizations/97473700/team/5714710",
    "permission": "pull",
    "members_count": 2,
    "repos_count": 0,
    "organization": {
        "login": "HybriStratus",
        "id": 97473700,
        "node_id": "O_kgDOBc9UpA",
        "url": "https://api.github.com/orgs/HybriStratus",
        "repos_url": "https://api.github.com/orgs/HybriStratus/repos",
        "description": null,
        "email": null,
        "type": "Organization"
    },
    "parent": null
}`)
	newTeam := Team{
		Name:    "test_team",
		Privacy: "closed",
	}
	// Create your table test
	tests := []struct {
		name           string
		requestClients []responses
		payload        Team
		expected       error
	}{
		{
			name:     "Testing successful update of team",
			payload:  newTeam,
			expected: nil,
			requestClients: []responses{
				{
					method: http.MethodPatch,
					url:    fmt.Sprintf("%s/%s/teams/%s", baseURL, TestOrg, newTeam.Name),
					response: http.Response{
						StatusCode: http.StatusOK,
						Body:       ConvertBytesToIoReadCloser([]byte(createResponse)),
					},
				},
			},
		},
		{
			name:     "Testing failure of update of team",
			payload:  newTeam,
			expected: fmt.Errorf("Error in updating team : test_team"),
			requestClients: []responses{
				{
					method: http.MethodPatch,
					url:    fmt.Sprintf("%s/%s/teams/%s", baseURL, TestOrg, newTeam.Name),
					response: http.Response{
						StatusCode: http.StatusInternalServerError,
					},
				},
			},
		},
	}
	// Go through each of the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create the mock Client
			mockClient := mock.Client{}
			for _, response := range tt.requestClients {
				mockClient.SetResponses(response.method, response.url, response.response)
			}

			got := tt.payload.UpdateTeam(mockClient)
			if got != nil && tt.expected != nil {
				if got.Error() != tt.expected.Error() {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			} else {
				if got != tt.expected {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			}
		})
	}
}

// TestDeleteTeam tests DeleteTeam function of a team
func TestDeleteTeam(t *testing.T) {

	newTeam := Team{
		Name: "test_team",
	}
	// Create your table test
	tests := []struct {
		name           string
		requestClients []responses
		payload        Team
		expected       error
	}{
		{
			name:     "Testing successful deletion of team",
			payload:  newTeam,
			expected: nil,
			requestClients: []responses{
				{
					method: http.MethodDelete,
					url:    fmt.Sprintf("%s/%s/teams/%s", baseURL, TestOrg, newTeam.Name),
					response: http.Response{
						StatusCode: http.StatusNoContent,
					},
				},
			},
		},
		{
			name:     "Testing failure of deletion of team",
			payload:  newTeam,
			expected: fmt.Errorf("Error in deleting team : test_team"),
			requestClients: []responses{
				{
					method: http.MethodDelete,
					url:    fmt.Sprintf("%s/%s/teams/%s", baseURL, TestOrg, newTeam.Name),
					response: http.Response{
						StatusCode: http.StatusForbidden,
					},
				},
			},
		},
	}
	// Go through each of the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create the mock Client
			mockClient := mock.Client{}
			for _, response := range tt.requestClients {
				mockClient.SetResponses(response.method, response.url, response.response)
			}

			got := tt.payload.DeleteTeam(mockClient)
			if got != nil && tt.expected != nil {
				if got.Error() != tt.expected.Error() {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			} else {
				if got != tt.expected {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			}
		})
	}
}

// TestListMemebersOfTeam tests ListMemebersOfTeam function of a team
func TestListMemebersOfTeam(t *testing.T) {

	createResponse := fmt.Sprintf(`
	[
    {
        "login": "test_user1",
        "id": 90353216,
        "node_id": "MDQ6VXNlcjkwMzUzMjE2",
        "gravatar_id": "",
        "url": "https://api.github.com/users/test_user1",
        "html_url": "https://github.com/test_user1",
        "type": "User",
        "site_admin": false
    },
    {
        "login": "test_user2",
        "id": 94664144,
        "node_id": "U_kgDOBaR10A",
        "gravatar_id": "",
        "url": "https://api.github.com/users/test_user2",
        "html_url": "https://github.com/test_user2",
        "type": "User",
        "site_admin": false
    }
]`)

	newTeam := Team{
		Name: "test_team",
	}
	// Create your table test
	tests := []struct {
		name           string
		requestClients []responses
		payload        Team
		expected       error
	}{
		{
			name:     "Testing successful retrival of team members",
			payload:  newTeam,
			expected: nil,
			requestClients: []responses{
				{
					method: http.MethodGet,
					url:    fmt.Sprintf("%s/%s/teams/%s/members", baseURL, TestOrg, newTeam.Name),
					response: http.Response{
						StatusCode: http.StatusOK,
						Body:       ConvertBytesToIoReadCloser([]byte(createResponse)),
					},
				},
			},
		},
		{
			name:     "Testing failure of retival of team members",
			payload:  newTeam,
			expected: fmt.Errorf("Error in getting members of a team : test_team"),
			requestClients: []responses{
				{
					method: http.MethodGet,
					url:    fmt.Sprintf("%s/%s/teams/%s/members", baseURL, TestOrg, newTeam.Name),
					response: http.Response{
						StatusCode: http.StatusGatewayTimeout,
					},
				},
			},
		},
	}
	// Go through each of the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create the mock Client
			mockClient := mock.Client{}
			for _, response := range tt.requestClients {
				mockClient.SetResponses(response.method, response.url, response.response)
			}

			got := tt.payload.ListMemebersOfTeam(mockClient)
			if got != nil && tt.expected != nil {
				if got.Error() != tt.expected.Error() {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			} else {
				if got != tt.expected {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			}
		})
	}
}

// TestAddMemeberToTeam tests AddMemeberToTeam function of a team
func TestAddMemeberToTeam(t *testing.T) {

	createResponse := fmt.Sprintf(`
	{
    "state": "active",
    "role": "maintainer",
    "url": "https://api.github.com/organizations/97473711/team/5714700/memberships/test_user3"
	}
`)

	teamName := "test_team"
	userName := "test_user3"
	roleType := "maintainer"
	// Create your table test
	tests := []struct {
		name           string
		requestClients []responses
		team           string
		user           string
		role           string
		expected       error
	}{
		{
			name:     "Testing successful addition on new memeber to team",
			team:     teamName,
			user:     userName,
			role:     roleType,
			expected: nil,
			requestClients: []responses{
				{
					method: http.MethodPut,
					url:    fmt.Sprintf("%s/%s/teams/%s/memberships/%s", baseURL, TestOrg, teamName, userName),
					response: http.Response{
						StatusCode: http.StatusOK,
						Body:       ConvertBytesToIoReadCloser([]byte(createResponse)),
					},
				},
			},
		},
		{
			name:     "Testing failure of addition of new member to team",
			team:     teamName,
			user:     userName,
			role:     roleType,
			expected: fmt.Errorf("Error in adding test_user3 to team test_team"),
			requestClients: []responses{
				{
					method: http.MethodPut,
					url:    fmt.Sprintf("%s/%s/teams/%s/memberships/%s", baseURL, TestOrg, teamName, userName),
					response: http.Response{
						StatusCode: http.StatusUnprocessableEntity,
					},
				},
			},
		},
	}
	// Go through each of the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create the mock Client
			mockClient := mock.Client{}
			for _, response := range tt.requestClients {
				mockClient.SetResponses(response.method, response.url, response.response)
			}

			got := AddMemeberToTeam(mockClient, tt.team, tt.user, tt.role)
			if got != nil && tt.expected != nil {
				if got.Error() != tt.expected.Error() {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			} else {
				if got != tt.expected {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			}
		})
	}
}

// TestDeleteMemberFromTeam tests DeleteMemberFromTeam function of a team
func TestDeleteMemberFromTeam(t *testing.T) {

	teamName := "test_team"
	userName := "test_user3"
	// Create your table test
	tests := []struct {
		name           string
		requestClients []responses
		team           string
		user           string
		expected       error
	}{
		{
			name:     "Testing successful deletion of member from a team",
			team:     teamName,
			user:     userName,
			expected: nil,
			requestClients: []responses{
				{
					method: http.MethodDelete,
					url:    fmt.Sprintf("%s/%s/teams/%s/memberships/%s", baseURL, TestOrg, teamName, userName),
					response: http.Response{
						StatusCode: http.StatusNoContent,
					},
				},
			},
		},
		{
			name:     "Testing failure of deletion of member from a team",
			team:     teamName,
			user:     userName,
			expected: fmt.Errorf("Error in deleting test_user3 from team test_team"),
			requestClients: []responses{
				{
					method: http.MethodDelete,
					url:    fmt.Sprintf("%s/%s/teams/%s/memberships/%s", baseURL, TestOrg, teamName, userName),
					response: http.Response{
						StatusCode: http.StatusConflict,
					},
				},
			},
		},
	}
	// Go through each of the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create the mock Client
			mockClient := mock.Client{}
			for _, response := range tt.requestClients {
				mockClient.SetResponses(response.method, response.url, response.response)
			}

			got := DeleteMemberFromTeam(mockClient, tt.team, tt.user)
			if got != nil && tt.expected != nil {
				if got.Error() != tt.expected.Error() {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			} else {
				if got != tt.expected {
					t.Errorf("wanted %v, got %v", tt.expected, got.Error())
				}
			}
		})
	}
}

// Converts an bytes into an io.ReadCloser, which is used as the body of an API call
func ConvertBytesToIoReadCloser(objectByte []byte) io.ReadCloser {
	objectReader := bytes.NewReader(objectByte)
	return ioutil.NopCloser(objectReader)
}
