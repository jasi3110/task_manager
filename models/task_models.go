// task_model.go
package models

import "task_manager"

type Task struct {
    ID          int64
    Title       string
    Description string
    Completed   bool
}

// Create a new Task model instance from the protobuf message
func NewTaskFromProto(protoTask *task_manager.Task) *Task {
    return &Task{
        ID:          protoTask.GetId(),
        Title:       protoTask.GetTitle(),
        Description: protoTask.GetDescription(),
        Completed:   protoTask.GetCompleted(),
    }
}

// Convert Task model to protobuf message
func (t *Task) ToProto() *task_manager.Task {
    return &task_manager.Task{
        Id:          t.ID,
        Title:       t.Title,
        Description: t.Description,
        Completed:   t.Completed,
    }
}
