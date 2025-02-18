package service

import (
	"fmt"
	"mpbench/database"
	"mpbench/model"
	"mpbench/utils"
)

func GetAllRuns() []model.Run {
	var runs []model.Run
	database.DB.Find(&runs)
	return runs
}

func GetRunsByCommit(commit string) []model.Run {
	var runs []model.Run
	database.DB.Where("commit = ?", commit).Find(&runs)
	return runs
}

func GetRunsByService(service string) []model.Run {
	var runs []model.Run
	database.DB.Where("service = ?", service).Find(&runs)
	return runs
}

func GetRunsByStatus(status string) []model.Run {
	var runs []model.Run
	database.DB.Where("status = ?", status).Find(&runs)
	return runs
}

func GetRunByID(id string) model.Run {
	var run model.Run
	database.DB.First(&run, "id = ?", id)
	return run
}

func GetRunByGithubCheckRunID(githubCheckRunID int) model.Run {
	var run model.Run
	database.DB.Where("github_check_run_id = ?", githubCheckRunID).First(&run)
	return run
}

func CreateRun(run model.Run) error {
	if run.ID == "" {
		return fmt.Errorf("run id cannot be empty")
	}
	if database.DB.Where("id = ?", run.ID).Updates(&run).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New run created with id: " + run.ID)
		if result := database.DB.Create(&run); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("Run with id: " + run.ID + " has been updated!")
	}
	return nil
}
