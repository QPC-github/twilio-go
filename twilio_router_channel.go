package twilio

import (
	"context"
	"net/url"
)

const TaskChannelPathPart = "TaskChannels"

type TaskChannelService struct {
	client       *Client
	workspaceSid string
}

type TaskChannel struct {
	AccountSid              string            `json:"account_sid"`
	DateCreated             TwilioTime        `json:"date_created"`
	DateUpdated             TwilioTime        `json:"date_updated"`
	FriendlyName            string            `json:"friendly_name"`
	Sid                     string            `json:"sid"`
	UniqueName              string            `json:"unique_name"`
	WorkspaceSid            string            `json:"workspace_sid"`
	ChannelOptimizedRouting bool              `json:"channel_optimized_routing"`
	URL                     string            `json:"url"`
	Links                   map[string]string `json:"links"`
}

type TaskChannelPage struct {
	Page
	TaskChannels []*Worker `json:"task_channels"`
}

func (r *TaskChannelService) Get(ctx context.Context, sid string) (*TaskChannel, error) {
	taskChannel := new(TaskChannel)
	err := r.client.GetResource(ctx, "Workspaces/"+r.workspaceSid+"/"+TaskChannelPathPart, sid, taskChannel)
	return taskChannel, err
}

func (r *TaskChannelService) Create(ctx context.Context, data url.Values) (*TaskChannel, error) {
	taskChannel := new(TaskChannel)
	err := r.client.CreateResource(ctx, "Workspaces/"+r.workspaceSid+"/"+TaskChannelPathPart, data, taskChannel)
	return taskChannel, err
}

// Delete deletes a Workspace.
//
// See https://www.twilio.com/docs/taskrouter/api/workers#action-delete for more.
func (r *TaskChannelService) Delete(ctx context.Context, sid string) error {
	return r.client.DeleteResource(ctx, "Workspaces/"+r.workspaceSid+"/"+TaskChannelPathPart, sid)
}

// // Update updates a Workspace.
// //
// // See https://www.twilio.com/docs/taskrouter/api/workspaces#action-update for more.
func (r *TaskChannelService) Update(ctx context.Context, sid string, data url.Values) (*Workspace, error) {
	worker := new(Workspace)
	err := r.client.UpdateResource(ctx, "Workspaces/"+r.workspaceSid+"/"+TaskChannelPathPart, sid, data, worker)
	return worker, err
}
