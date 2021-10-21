package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

func main() {
	organizationUrl := "https://lanlink.visualstudio.com/"
	personalAccessToken := "anxxu6hpt3y6cuyu4hnekrb3w7qxognky4avmg65glfdly4aej5q"

	//Create a connection with the DevOps organization
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx := context.Background()

	workItemClient, err := workitemtracking.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	// item, err := workitem.GetWorkItemById(workItemClient, ctx, 8763)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fields, err := json.Marshal(item.Fields)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Print(string(fields))

	var op webapi.Operation = "add"
	path := "/fields/System.Title"
	var from string
	value := "Sample task"

	jsonPatch := webapi.JsonPatchOperation{
		Op:    &op,
		Path:  &path,
		From:  &from,
		Value: &value,
	}

	patchs := []webapi.JsonPatchOperation{jsonPatch}
	project := "Business Transformation"
	itemType := "Bug"

	createWorkItemArgs := workitemtracking.CreateWorkItemArgs{
		Document: &patchs,
		Project:  &project,
		Type:     &itemType,
	}

	createdWorkItem, err := workItemClient.CreateWorkItem(ctx, createWorkItemArgs)
	if err != nil {
		log.Fatal(err)
	}

	fields, err := json.Marshal(createdWorkItem.Fields)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(fields))
}
