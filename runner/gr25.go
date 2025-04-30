package runner

import (
	"mpbench/model"
	"mpbench/service"
	"mpbench/utils"

	"github.com/google/uuid"
)

func CreateGR25Runs(commit string) ([]model.Run, error) {
	runs := []model.Run{}
	// Create gr25 unit test run
	id, err := service.CreateCheckRun(commit, "gr25 / unit")
	if err != nil {
		utils.SugarLogger.Errorf("Error creating check run: %v", err)
	} else {
		utils.SugarLogger.Infof("Check run created with ID: %d", id)

		run := model.Run{
			ID:               uuid.New().String(),
			Commit:           commit,
			Status:           "queued",
			Name:             "gr25 / unit",
			Service:          "gr25",
			GithubCheckRunID: id,
		}
		err = service.CreateRun(run)
		if err != nil {
			utils.SugarLogger.Errorf("Error creating run: %v", err)
		} else {
			Queue.Add(run)
			runs = append(runs, run)
		}
	}

	// Create gr25 benchmark run
	id, err = service.CreateCheckRun(commit, "gr25 / benchmark")
	if err != nil {
		utils.SugarLogger.Errorf("Error creating check run: %v", err)
	} else {
		utils.SugarLogger.Infof("Check run created with ID: %d", id)

		run := model.Run{
			ID:               uuid.New().String(),
			Commit:           commit,
			Status:           "queued",
			Name:             "gr25 / benchmark",
			Service:          "gr25",
			GithubCheckRunID: id,
		}
		err = service.CreateRun(run)
		if err != nil {
			utils.SugarLogger.Errorf("Error creating run: %v", err)
		} else {
			Queue.Add(run)
			runs = append(runs, run)
		}
	}
	return runs, nil
}
