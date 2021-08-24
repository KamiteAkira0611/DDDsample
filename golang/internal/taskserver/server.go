package taskserver

import (
	"context"
	"dddsample/rpc/task"
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

	// n := goon.NewGoon(r)
	// tasks := []Task{}
	// q := datastore.NewQuery("Task")

	// aaa, err := n.GetAll(q, &tasks)
	// if err != nil {
	// 	return nil, twirp.NewError(twirp.Internal, err.Error())
	// }

	// fmt.Fprintf(w, "invalid request. %s", aaa)
	// fmt.Fprintf(w, "invalid request. %s", err.Error())

	return &task.ListTaskResp{}, nil
}

func (s *TaskServer) CreateTask(ctx context.Context, req *task.CreateTaskReq) (resp *task.Task, err error) {
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
