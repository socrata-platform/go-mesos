package mesos

import (
    "io"
    "io/ioutil"
    "regexp"
)

type Config struct {
    DiscoveryURL string
    MasterURL string
    MasterPort int
    SlavePort int
    LogOutput io.Writer
    RequestTimeout int
}

func NewDefaultConfig() Config {
    return Config{
        MasterPort: 5050,
        SlavePort: 5051,
        LogOutput:      ioutil.Discard,
        RequestTimeout: 5,
    }
}


func (c *Config) getScheme() (string) {
    re, _ := regexp.Compile(`http(s)?:`)
    return re.FindString(c.DiscoveryURL)
}
