package model

import "time"

type Run struct {
	ID               string    `gorm:"primaryKey" json:"id"`
	Name             string    `json:"name"`
	Service          string    `json:"service"`
	Commit           string    `json:"commit"`
	Status           string    `json:"status"`
	GithubCheckRunID int       `json:"github_check_run_id"`
	RunTests         []RunTest `gorm:"-" json:"run_tests"`
	CreatedAt        time.Time `gorm:"autoCreateTime;precision:6" json:"created_at"`
}

func (Run) TableName() string {
	return "run"
}

type RunTest struct {
	ID             string          `gorm:"primaryKey" json:"id"`
	RunID          string          `json:"run_id"`
	Name           string          `json:"name"`
	Status         string          `json:"status"`
	Data           string          `json:"data"`
	RunTestResults []RunTestResult `gorm:"-" json:"run_test_results"`
	CreatedAt      time.Time       `gorm:"autoCreateTime;precision:6" json:"created_at"`
}

func (RunTest) TableName() string {
	return "run_test"
}

type RunTestResult struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	RunTestID  string    `json:"run_test_id"`
	SignalName string    `json:"signal_name"`
	Status     string    `json:"status"`
	Value      string    `json:"value"`
	Expected   string    `json:"expected"`
	CreatedAt  time.Time `gorm:"autoCreateTime;precision:6" json:"created_at"`
}

func (RunTestResult) TableName() string {
	return "run_test_result"
}
