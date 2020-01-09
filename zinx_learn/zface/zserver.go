package zface

type Zserver interface {
	Start()
	Stop()
	Serve()
}
