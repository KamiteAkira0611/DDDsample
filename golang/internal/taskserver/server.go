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
	tasks := []Task{}
	q := datastore.NewQuery("Task")

	_, err = n.GetAll(q, &tasks)
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, err.Error())
	}

	var items []*task.Task

	for _, s := range tasks {
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
	// n := goon.FromContext(ctx)
	// body := req.GetBody()

	// g := &Task{Body: body}

	// key, err := n.Put(g)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprintf(w, "invalid request. %s", err.Error())
	// 	return
	// }
	
	return &task.Task{
		Id: 111,
		Body: "createTask " + req.GetBody(),
	}, nil
}

func (s *TaskServer) DeleteTask(ctx context.Context, req *task.DeleteTaskReq) (resp *task.Task, err error) {
	return &task.Task{
		Id: req.GetId(),
		Body: "deleteTask",
	}, nil
}
