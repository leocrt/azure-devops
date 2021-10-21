package workitem

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

func GetWorkItemById(client workitemtracking.Client, ctx context.Context, id int) (*workitemtracking.WorkItem, error) {
	referenceId := id
	item, err := client.GetWorkItem(ctx, workitemtracking.GetWorkItemArgs{Id: &referenceId})
	return item, err
}
