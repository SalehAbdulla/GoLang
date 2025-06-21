package main

import (
    "errors"
    "fmt"
    "sync"
    "time"
)

// Constants & Variables
const appVersion = "1.0"
var taskID int = 1

// Struct & Methods
type Task struct {
    ID          int
    Name        string
    Description string
    Priority    int
    Done        bool
}

func (t *Task) MarkDone() {
    t.Done = true
}

// Interface
type TaskManager interface {
    AddTask(name string, desc string, priority int)
    ListTasks()
    DeleteTask(id int) error
    CompleteTask(id int) error
}

// Slice, Map, and Struct Usage
type SimpleTaskManager struct {
    tasks []Task
    taskMap map[int]*Task
}

// Constructor Function
func NewTaskManager() *SimpleTaskManager {
    return &SimpleTaskManager{
        tasks:   make([]Task, 0),
        taskMap: make(map[int]*Task),
    }
}

// Callback & Methods
func (tm *SimpleTaskManager) AddTask(name string, desc string, priority int) {
    newTask := Task{ID: taskID, Name: name, Description: desc, Priority: priority}
    tm.tasks = append(tm.tasks, newTask)
    tm.taskMap[taskID] = &tm.tasks[len(tm.tasks)-1]
    taskID++
}

func (tm *SimpleTaskManager) ListTasks() {
    fmt.Println("\nTasks:")
    printSlice(tm.tasks)
}

func (tm *SimpleTaskManager) CompleteTask(id int) error {
    if task, ok := tm.taskMap[id]; ok {
        task.MarkDone()
        return nil
    }
    return errors.New("task not found")
}

func (tm *SimpleTaskManager) DeleteTask(id int) error {
    if _, ok := tm.taskMap[id]; !ok {
        return errors.New("task not found")
    }
    delete(tm.taskMap, id)
    for i, task := range tm.tasks {
        if task.ID == id {
            tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
            break
        }
    }
    return nil
}

// Generics
func printSlice[T any](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}

// Defer, Panic, Recover
func safeRun(fn func()) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    fn()
}

// Goroutines, Channels, WaitGroup
func asyncSaveTasks(tasks []Task, wg *sync.WaitGroup, ch chan string) {
    defer wg.Done()
    time.Sleep(1 * time.Second)
    ch <- fmt.Sprintf("Saved %d tasks to disk!", len(tasks))
}

func main() {
    fmt.Println("🔧 Welcome to CLI Task Manager v" + appVersion)

    manager := NewTaskManager()
    manager.AddTask("Learn Go", "Study generics and interfaces", 1)
    manager.AddTask("Build Project", "Create a Go CLI app", 2)

    // Function Expression
    showStats := func() {
        fmt.Printf("📋 Total Tasks: %d\n", len(manager.tasks))
    }
    showStats()

    manager.ListTasks()

    // Logical, Comparison Operators, If
    if err := manager.CompleteTask(1); err != nil {
        fmt.Println("Error:", err)
    }

    // Pointer Concept
    t := &manager.tasks[0]
    fmt.Println("✅ Marked done:", t.Name, "Status:", t.Done)

    // Defer & Panic Simulation
    safeRun(func() {
        panic("Something went wrong internally")
    })

    // Goroutines + WaitGroup
    var wg sync.WaitGroup
    ch := make(chan string)
    wg.Add(1)
    go asyncSaveTasks(manager.tasks, &wg, ch)

    // Select Example
    go func() {
        wg.Wait()
        close(ch)
    }()

    select {
    case msg := <-ch:
        fmt.Println("✅", msg)
    case <-time.After(2 * time.Second):
        fmt.Println("⏳ Timeout saving tasks")
    }

    // Switch Statement
    taskStatus := "done"
    switch taskStatus {
    case "done":
        fmt.Println("📌 Great job!")
    case "pending":
        fmt.Println("🚧 Keep going!")
    default:
        fmt.Println("❓ Unknown status")
    }
}
