// task_controller.go
package controllers

import (
    "task_manager"
    "task_manager/models"
)

var tasks []*models.Task

// CreateTask adds a new task to the list
func CreateTask(task *task_manager.Task) {
    newTask := models.NewTaskFromProto(task)
    tasks = append(tasks, newTask)
}

// GetAllTasks returns a list of all tasks
func GetAllTasks() []*task_manager.Task {
    var taskList []*task_manager.Task
    for _, task := range tasks {
        taskList = append(taskList, task.ToProto())
    }
    return taskList
}

// ...
// Implement other CRUD methods as needed (UpdateTask, DeleteTask, GetTaskByID, etc.)
