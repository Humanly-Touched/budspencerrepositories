package models

type Worker struct {
	Uuid  string
	Host  string
	Ip    string
	Port  int
	State int
}
