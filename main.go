package main

import (
	"fmt"

	"github.com/HybriStratus/test-github-groups/groups"
	"github.com/HybriStratus/test-github-groups/http/net"
)

func main() {
	fmt.Println("Performing CRUD operations on Github teams")

	newTeam := groups.Team{
		Name:        "iac-platfom-test-team",
		Description: "Created a new test team",
		Privacy:     "secret",
	}
	// Create HTTP client
	client := net.Client{}
	fmt.Printf("Creating a new Team under %s Org %v\n\n", groups.TestOrg, newTeam)
	err := newTeam.CreateTeam(client)
	if err != nil {
		fmt.Printf(err.Error())
	}

	newTeam.Privacy = "closed"
	newTeam.Description = "Updated the description"
	fmt.Printf("Updating a %s under Org %v\n\n", groups.TestOrg, newTeam)
	err = newTeam.UpdateTeam(client)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("Get team details for team %s\n\n", newTeam.Name)
	err = newTeam.GetTeamDetails(client)
	if err != nil {
		fmt.Printf(err.Error())
	}

	addMembers := []string{"meav", "emuthusa", "rajpa", "daspano", "pkasargo", "jdwidari"}
	for _, mem := range addMembers {
		fmt.Printf("Adding %s to the team %s\n\n", mem, newTeam.Name)
		err = groups.AddMemeberToTeam(client, newTeam.Name, mem, groups.DefaultRoleType)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	}

	fmt.Printf("List memebers of %s\n\n", newTeam.Name)
	err = newTeam.ListMemebersOfTeam(client)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	delMembers := []string{"meav", "rajpa", "jdwidari"}
	for _, mem := range delMembers {
		fmt.Printf("Deleting %s from the team %s\n\n", mem, newTeam.Name)
		err = groups.DeleteMemberFromTeam(client, newTeam.Name, mem)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	}

	fmt.Printf("List memebers of %s\n\n", newTeam.Name)
	err = newTeam.ListMemebersOfTeam(client)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Printf("Delete the team %s Org %s\n\n", newTeam.Name, groups.TestOrg)
	newTeam.DeleteTeam(client)

}
