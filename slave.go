package mesos

import (
    "time"
)

type Slave struct {
    Id             string            `json:"id"`
    HostName       string            `json:"hostname"`
    Attributes     map[string]string `json:"attributes"`
    Resources      Resource          `json:"resources"`
    Frameworks     []*Framework      `json:"frameworks"`
    StateLoadError bool
    Stats          *SlaveStats
    StateLastUpdated time.Time
    StatsLastUpdated time.Time
}

type SlaveStats struct {
    FailedTasks              int32   `json:"failed_tasks"`
    FinishedTasks            int32   `json:"finished_tasks"`
    KilledTasks              int32   `json:"killed_tasks"`
    LostTasks                int32   `json:"lost_tasks"`
    CpusPercent              float64 `json:"slave/cpus_percent"`
    CpusTotal                int32   `json:"slave/cpus_total"`
    CpusUsed                 float64 `json:"slave/cpus_used"`
    DiskPercent              float64 `json:"slave/disk_percent"`
    DiskTotal                int32   `json:"slave/disk_total"`
    DiskUsed                 int32   `json:"slave/disk_used"`
    ExecutorsRegistering     int32   `json:"slave/executors_registering"`
    ExecutorsRunning         int32   `json:"slave/executors_running"`
    ExecutorsTerminated      int32   `json:"slave/executors_terminated"`
    ExecutorsTerminating     int32   `json:"slave/executors_terminating"`
    FrameworksActive         int32   `json:"slave/frameworks_active"`
    InvalidFrameworkMessages int32   `json:"slave/invalid_framework_messages"`
    InvalidStatusUpdates     int32   `json:"slave/invalid_status_updates"`
    MemPercent               float64 `json:"slave/mem_percent"`
    MemTotal                 int32   `json:"slave/mem_total"`
    MemUsed                  int32   `json:"slave/mem_used"`
    RecoveryErrors           int32   `json:"slave/recovery_errors"`
    Registered               int32   `json:"slave/registered"`
    TasksFailed              int32   `json:"slave/tasks_failed"`
    TasksFinished            int32   `json:"slave/tasks_finished"`
    TasksKilled              int32   `json:"slave/tasks_killed"`
    TasksLost                int32   `json:"slave/tasks_lost"`
    TasksRunning             int32   `json:"slave/tasks_running"`
    TasksStaging             int32   `json:"slave/tasks_staging"`
    TasksStarting            int32   `json:"slave/tasks_starting"`
    UptimeSecs               float64 `json:"slave/uptime_secs"`
    ValidFrameworkMessages   int32   `json:"slave/valid_framework_messages"`
    ValidStatusUpdates       int32   `json:"slave/valid_status_updates"`
    StagedTasks              int32   `json:"staged_tasks"`
    StartedTasks             int32   `json:"started_tasks"`
    Load1Min                 float64 `json:"system/load_1min"`
    Load5Min                 float64 `json:"system/load_5min"`
    Load15Min                float64 `json:"system/load_15min"`
    MemFreeBytes             int64   `json:"system/mem_free_bytes"`
    MemTotalBytes            int64   `json:"system/mem_total_bytes"`
    TotalFrameworks          int32   `json:"total_frameworks"`
    Uptime                   float64 `json:"uptime"`
}


func (s *Slave) LoadState(client MesosClient) (error) {
    url := client.slaveStateURL(s.HostName)
    _, _, err := client.doApiRequest(url, s)
    if err != nil {
        s.StateLoadError = true
    } else {
        s.StateLastUpdated = time.Now()
    }
    return err
}

func (s *Slave) LoadStats(client MesosClient) (error) {
    ss := &SlaveStats{}
    url := client.slaveStatsURL(s.HostName)
    if _, _, err := client.doApiRequest(url, ss); err != nil {
        return err
    }
    s.StatsLastUpdated = time.Now()
    s.Stats = ss
    return nil
}


func (s *Slave) GetFramework(framework string) (map[string]*Framework) {
    fs := make(map[string]*Framework)
    for _, f := range s.Frameworks {
        if f.Name == framework {
            fs[f.Id] = f
        }
    }
    return fs
}
