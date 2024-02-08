package main

type InstanceInfo struct {
	name     string
	id       string
	publicIp string
}

func (i InstanceInfo) List() string {
	return i.name + i.id + i.publicIp
}
