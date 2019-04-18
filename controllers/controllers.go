package controllers

type Deployment struct {
	Name    string
	Healthy bool
	Running bool
	Staged  bool
	Created bool
}
