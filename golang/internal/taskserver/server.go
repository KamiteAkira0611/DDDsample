package taskserver

import (
	"context"
	"dddsample/rpc/task"
)

// Server implements the Haberdasher service
type TaskServer struct{}

func (s *TaskServer) ListTask(ctx context.Context, req *task.ListTaskReq) (resp *task.ListTaskResp, err error) {
	return &task.ListTaskResp{}, nil
}
// func (s *TaskServer) CreateTask(ctx context.Context, req *task.ListTaskReq) (resp *task.ListTaskResp, err error) {
// 	return &task.ListTaskResp{}, nil
// }
// func (s *TaskServer) DeleteTask(ctx context.Context, req *task.ListTaskReq) (resp *task.ListTaskResp, err error) {
// 	return &task.ListTaskResp{}, nil
// }
