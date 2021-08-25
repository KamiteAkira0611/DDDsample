package taskserver

import (
	"context"
	"dddsample/rpc/task"

	"github.com/mjibson/goon"
	"github.com/twitchtv/twirp"
	"google.golang.org/appengine/datastore"
)

// Server implements the Haberdasher service
type TaskServer struct{}

type Task struct {
	ID int64 `datastore:"-" goon:"id"`
	Body string 
}

// var w http.ResponseWriter
// var r *http.Request

func (s *TaskServer) ListTask(ctx context.Context, req *task.ListTaskReq) (resp *task.ListTaskResp, err error) {

	n := goon.FromContext(ctx)
	q_tasks := []Task{}
	q := datastore.NewQuery("Task")

	_, err = n.GetAll(q, &q_tasks)
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, err.Error())
	}

	var items []*task.Task

	for _, s := range q_tasks {
		item := task.Task{
			Id: s.ID,
			Body: s.Body,
		}
		items = append(items, &item)
	}

	result := &task.ListTaskResp{Tasks: items}
	return result, nil
}

func (s *TaskServer) CreateTask(ctx context.Context, req *task.CreateTaskReq) (resp *task.Task, err error) {
	n := goon.FromContext(ctx)
	body := req.GetBody()
	if body == "" {
		return nil, twirp.NewError(twirp.Internal, "body is undefined")
	}

	q_task := &Task{Body: body}

	_, err = n.Put(q_task)
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, err.Error())
	}
	
	return &task.Task{
		Id: q_task.ID,
		Body: q_task.Body,
	}, nil
}

func (s *TaskServer) DeleteTask(ctx context.Context, req *task.DeleteTaskReq) (resp *task.Task, err error) {
	n := goon.FromContext(ctx)
	id := req.GetId()

	q_task := &Task{ID: id}
	getErr := n.Get(q_task)
	if getErr != nil {
		return nil, twirp.NewError(twirp.Internal, getErr.Error())
	}

	deleteErr := n.Delete(q_task)
	if deleteErr != nil {
		return nil, twirp.NewError(twirp.Internal, deleteErr.Error())
	}

	return &task.Task{
		Id: q_task.ID,
		Body: q_task.Body,
	}, nil
}
