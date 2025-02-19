package runner

import (
	"mpbench/model"

	"github.com/google/uuid"
)

func CreateGR25Run(commit string) (model.Run, error) {
	run := model.Run{
		ID:      uuid.New().String(),
		Commit:  commit,
		Status:  "queued",
		Name:    "mpbench / gr25",
		Service: "gr25",
	}
	return run, nil
}
