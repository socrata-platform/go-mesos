package mesos

import (
    "regexp"
    "fmt"
)


func DiscoverCluster(client MesosClient) (*Cluster, error) {
    uri := "master/state.json"
    url := client.buildDiscoveryURL(uri)

    cluster := &Cluster{}
    if _, _, err := client.doApiRequest(url, cluster); err != nil {
        return cluster, ErrClusterDiscoveryError
    }
    client.setMasterURL(cluster.getLeader())

    if len(cluster.Frameworks) == 0 {
        url := fmt.Sprintf("%s/%s", client.masterURL, uri)
        client.doApiRequest(url, cluster)
    }

    return cluster, nil
}


type Cluster struct {
    LeaderPID  string       `json:"leader"`
    Version    string       `json:"version"`
    Frameworks []*Framework `json:"frameworks,omitempty"`
    Slaves     []*Slave     `json:"slaves"`
}


func (c *Cluster) getLeader() (string) {
    // String to extract IP from: master@10.110.37.146:5050
    re, _ := regexp.Compile(`[\d]+\.[\d]+\.[\d]+\.[\d]+`)
    return re.FindString(c.LeaderPID)
}


func (c *Cluster) LoadSlaveStates(client MesosClient) (error) {
    var erred bool

    for i := range c.Slaves {
        if err := c.Slaves[i].LoadState(client); err != nil {
            erred = true
        }
    }
    if erred { return ErrSlaveStateLoadError }
    return nil
}


func (c *Cluster) LoadSlaveStats(client MesosClient) (error) {
    var erred bool

    for i := range c.Slaves {
        if err := c.Slaves[i].LoadStats(client); err != nil {
            erred = true
        }
    }
    if erred { return ErrSlaveStateLoadError }
    return nil
}


func (c *Cluster) GetFramework(framework string) (map[string]*Framework) {
    fs := make(map[string]*Framework)
    for _, f := range c.Frameworks {
        if f.Name == framework {
            fs[f.Id] = f
        }
    }
    return fs
}


func (c *Cluster) GetSlaveById(slaveId string) (*Slave) {
    var slave *Slave
    for _, s := range c.Slaves {
        if s.Id == slaveId {
            slave = s
        }
    }
    return slave
}


func (c *Cluster) GetSlaveByHostName(hostName string) (*Slave) {
    var slave *Slave
    for _, s := range c.Slaves {
        if s.HostName == hostName {
            slave = s
        }
    }
    return slave
}
