package mesos

import (
    "regexp"
    "fmt"
)


func DiscoverCluster(client *Client) (*Cluster, error) {
    uri := "master/state.json"
    url := fmt.Sprintf("%s/%s", client.config.DiscoveryURL, uri)

    cluster := &Cluster{}
    if _, _, err := client.doApiRequest(url, cluster); err != nil {
        return cluster, ErrClusterDiscoveryError
    }
    client.config.MasterURL = fmt.Sprintf("%s//%s:%d", client.config.getScheme(), cluster.getLeader(), client.config.MasterPort)

    if len(cluster.Frameworks) == 0 {
        url := fmt.Sprintf("%s/%s", client.config.MasterURL, uri)
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


func (c *Cluster) LoadSlaveStates(client *Client) (error) {
    var erred bool

    for i := range c.Slaves {
        if err := c.Slaves[i].LoadState(client); err != nil {
            erred = true
        }
    }
    if erred { return ErrSlaveStateLoadError }
    return nil
}


func (c *Cluster) LoadSlaveStats(client *Client) (error) {
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
