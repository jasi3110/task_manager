// main.go
package main

import (
    "fmt"
    "github.com/jasi3110/proto_project/protocol_buffer"
    "github.com/jasi3110/proto_project/controllers"
)


func main() {
    // Create a new task
    task := &task_manager.Task{
        Id:          1,
        Title:       "Sample Task",
        Description: "This is a sample task.",
        Completed:   false,
    }
    controllers.CreateTask(task)

    // Get all tasks
    allTasks := controllers.GetAllTasks()
    for _, task := range allTasks {
        fmt.Println("Task ID:", task.GetId())
        fmt.Println("Title:", task.GetTitle())
        fmt.Println("Description:", task.GetDescription())
        fmt.Println("Completed:", task.GetCompleted())
        fmt.Println()
    }
}
