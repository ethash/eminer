package client

// Client interface
type Client interface {
	GetWork() ([]string, error)
	SubmitHashrate(params interface{}) (bool, error)
	SubmitWork(params interface{}) (bool, error)
}
