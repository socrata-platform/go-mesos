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
	ContainerDestroyErrors		float64	`json:"containerizer/mesos/container_destroy_errors"`
	ContainerLaunchErrors		float64	`json:"slave/container_launch_errors"`
	CpusPercent			float64	`json:"slave/cpus_percent"`
	CpusRevocablePercent		float64	`json:"slave/cpus_revocable_percent"`
	CpusRevocableTotal		float64	`json:"slave/cpus_revocable_total"`
	CpusRevocableUsed		float64	`json:"slave/cpus_revocable_used"`
	CpusTotal			float64	`json:"slave/cpus_total"`
	CpusUsed			float64	`json:"slave/cpus_used"`
	DiskPercent			float64	`json:"slave/disk_percent"`
	DiskRevocablePercent		float64	`json:"slave/disk_revocable_percent"`
	DiskRevocableTotal		float64	`json:"slave/disk_revocable_total"`
	DiskRevocableUsed		float64	`json:"slave/disk_revocable_used"`
	DiskTotal			float64	`json:"slave/disk_total"`
	DiskUsed			float64	`json:"slave/disk_used"`
	ExecDirMaxAge			float64	`json:"slave/executor_directory_max_allowed_age_secs"`
	ExecutorsPreempted		float64	`json:"slave/executors_preempted"`
	ExecutorsRegistering		float64	`json:"slave/executors_registering"`
	ExecutorsRunning		float64	`json:"slave/executors_running"`
	ExecutorsTerminated		float64	`json:"slave/executors_terminated"`
	ExecutorsTerminating		float64	`json:"slave/executors_terminating"`
	FrameworksActive		float64	`json:"slave/frameworks_active"`
	InvalidFrameworkMessages	float64	`json:"slave/invalid_framework_messages"`
	InvalidStatusUpdates		float64	`json:"slave/invalid_status_updates"`
	MemPercent			float64	`json:"slave/mem_percent"`
	MemRevocablePercent		float64	`json:"slave/mem_revocable_percent"`
	MemRevocableTotal		float64	`json:"slave/mem_revocable_total"`
	MemRevocableUsed		float64	`json:"slave/mem_revocable_used"`
	MemTotal			float64	`json:"slave/mem_total"`
	MemUsed				float64	`json:"slave/mem_used"`
	RecoveryErrors			float64	`json:"slave/recovery_errors"`
	Registered			float64	`json:"slave/registered"`
	TasksFailed			float64	`json:"slave/tasks_failed"`
	TasksFinished			float64	`json:"slave/tasks_finished"`
	TasksKilled			float64	`json:"slave/tasks_killed"`
	TasksLost			float64	`json:"slave/tasks_lost"`
	TasksRunning			float64	`json:"slave/tasks_running"`
	TasksStaging			float64	`json:"slave/tasks_staging"`
	TasksStarting			float64	`json:"slave/tasks_starting"`
	UptimeSecs			float64	`json:"slave/uptime_secs"`
	ValidFrameworkMessages		float64	`json:"slave/valid_framework_messages"`
	ValidStatusUpdates		float64	`json:"slave/valid_status_updates"`
	Load1Min			float64	`json:"system/load_1min"`
	Load5Min			float64	`json:"system/load_5min"`
	Load15Min			float64	`json:"system/load_15min"`
	MemFreeBytes			float64	`json:"system/mem_free_bytes"`
	MemTotalBytes			float64	`json:"system/mem_total_bytes"`
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
    ss := new(SlaveStats)
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
