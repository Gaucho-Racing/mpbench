package model

type Run struct {
	ID               string `gorm:"primaryKey" json:"id"`
	Name             string `json:"name"`
	Service          string `json:"service"`
	Commit           string `json:"commit"`
	Status           string `json:"status"`
	GithubCheckRunID int    `json:"github_check_run_id"`
}

func (Run) TableName() string {
	return "run"
}

type RunTest struct {
	ID     string `gorm:"primaryKey" json:"id"`
	RunID  string `json:"run_id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (RunTest) TableName() string {
	return "run_test"
}

type RunTestResult struct {
	ID         string  `gorm:"primaryKey" json:"id"`
	RunTestID  string  `json:"run_test_id"`
	SignalName string  `json:"signal_name"`
	Data       string  `json:"data"`
	Value      float64 `json:"value"`
	Expected   float64 `json:"expected"`
}

func (RunTestResult) TableName() string {
	return "run_test_result"
}
