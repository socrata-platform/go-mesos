package mesos

import (
    "strings"
    "fmt"
)


type Framework struct {
    Id               string      `json:"id"`
    Name             string      `json:"name"`
    Active           bool        `json:"active"`
    HostName         string      `json:"hostname"`
    Resources        *Resource   `json:"resources"`
    OfferedResources *Resource   `json:"offered_resources"`
    UsedResources    *Resource   `json:"used_resources"`
    Tasks            []*Task     `json:"tasks"`
    CompletedTasks   []*Task     `json:"completed_tasks"`
    Executors        []*Executor `json:"executors"`
}


type Executor struct {
    Id        string    `json:"id"`
    Name      string    `json:"name"`
    SlaveId   string    `json:"slave_id"`
    State     string    `json:"state"`
    Container string    `json:"container"`
    Resources *Resource `json:"resources"`
    Tasks     []*Task   `json:"tasks"`
}


type Task struct {
    Id               string    `json:"id"`
    Name             string    `json:"name"`
    Resources        *Resource `json:"resources"`
    SlaveId          string    `json:"slave_id"`
    State            string    `json:"state"`
    Statuses         []*Status `json:"statuses"`
}

type Status struct {
    State       string    `json:"state"`
    Timestamp   float64   `json:"timestamp"`
}

type Resource struct {
    Cpus  float32    `json:"cpus"`
    Disk  float32    `json:"disk"`
    Mem   float32    `json:"mem"`
    Ports string     `json:"ports"`
}


func (e *Executor) RegisteredContainerName(t *Task) (string) {
    return fmt.Sprintf("mesos-%s.%s", t.SlaveId, e.Container)
}


func (t *Task) AppId() (string) {
    idParts := strings.Split(t.Name, ".")
    for i, j := 0, len(idParts)-1; i < j; i, j = i+1, j-1 {
        idParts[i], idParts[j] = idParts[j], idParts[i]
    }
    id := strings.Join(idParts, "/")
    return fmt.Sprintf("/%s", id)
}
