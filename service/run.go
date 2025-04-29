package service

import (
	"fmt"
	"mpbench/database"
	"mpbench/model"
	"mpbench/utils"
)

func GetAllRuns() []model.Run {
	var runs []model.Run
	database.DB.Order("created_at DESC").Find(&runs)
	for i, run := range runs {
		run.RunTests = GetRunTestsByRunID(run.ID)
		runs[i] = run
	}
	return runs
}

func GetRunsByCommit(commit string) []model.Run {
	var runs []model.Run
	database.DB.Where("commit = ?", commit).Order("created_at DESC").Find(&runs)
	for i, run := range runs {
		run.RunTests = GetRunTestsByRunID(run.ID)
		runs[i] = run
	}
	return runs
}

func GetRunsByService(service string) []model.Run {
	var runs []model.Run
	database.DB.Where("service = ?", service).Order("created_at DESC").Find(&runs)
	for i, run := range runs {
		run.RunTests = GetRunTestsByRunID(run.ID)
		runs[i] = run
	}
	return runs
}

func GetRunsByStatus(status string) []model.Run {
	var runs []model.Run
	database.DB.Where("status = ?", status).Order("created_at DESC").Find(&runs)
	for i, run := range runs {
		run.RunTests = GetRunTestsByRunID(run.ID)
		runs[i] = run
	}
	return runs
}

func GetRunByID(id string) model.Run {
	var run model.Run
	database.DB.First(&run, "id = ?", id)
	run.RunTests = GetRunTestsByRunID(run.ID)
	return run
}

func GetRunByGithubCheckRunID(githubCheckRunID int) model.Run {
	var run model.Run
	database.DB.Where("github_check_run_id = ?", githubCheckRunID).First(&run)
	run.RunTests = GetRunTestsByRunID(run.ID)
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

func GetAllRunTests() []model.RunTest {
	var runTests []model.RunTest
	database.DB.Order("created_at DESC").Find(&runTests)
	for i, runTest := range runTests {
		runTest.RunTestResults = GetRunTestResultsByRunTestID(runTest.ID)
		runTests[i] = runTest
	}
	return runTests
}

func GetRunTestsByRunID(runID string) []model.RunTest {
	var runTests []model.RunTest
	database.DB.Where("run_id = ?", runID).Order("created_at ASC").Find(&runTests)
	for i, runTest := range runTests {
		runTest.RunTestResults = GetRunTestResultsByRunTestID(runTest.ID)
		runTests[i] = runTest
	}
	return runTests
}

func GetRunTestByID(id string) model.RunTest {
	var runTest model.RunTest
	database.DB.First(&runTest, "id = ?", id)
	runTest.RunTestResults = GetRunTestResultsByRunTestID(runTest.ID)
	return runTest
}

func CreateRunTest(runTest model.RunTest) error {
	if runTest.ID == "" {
		return fmt.Errorf("run test id cannot be empty")
	} else if runTest.RunID == "" {
		return fmt.Errorf("run id cannot be empty")
	} else if runTest.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if database.DB.Where("id = ?", runTest.ID).Updates(&runTest).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New run test created with id: " + runTest.ID)
		if result := database.DB.Create(&runTest); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("Run test with id: " + runTest.ID + " has been updated!")
	}
	return nil
}

func GetAllRunTestResults() []model.RunTestResult {
	var runTestResults []model.RunTestResult
	database.DB.Order("created_at DESC").Find(&runTestResults)
	return runTestResults
}

func GetRunTestResultsByRunTestID(runTestID string) []model.RunTestResult {
	var runTestResults []model.RunTestResult
	database.DB.Where("run_test_id = ?", runTestID).Order("created_at ASC").Find(&runTestResults)
	return runTestResults
}

func GetRunTestResultByID(id string) model.RunTestResult {
	var runTestResult model.RunTestResult
	database.DB.First(&runTestResult, "id = ?", id)
	return runTestResult
}

func CreateRunTestResult(runTestResult model.RunTestResult) error {
	if runTestResult.ID == "" {
		return fmt.Errorf("run test result id cannot be empty")
	} else if runTestResult.RunTestID == "" {
		return fmt.Errorf("run test id cannot be empty")
	}
	if database.DB.Where("id = ?", runTestResult.ID).Updates(&runTestResult).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New run test result created with id: " + runTestResult.ID)
		if result := database.DB.Create(&runTestResult); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("Run test result with id: " + runTestResult.ID + " has been updated!")
	}
	return nil
}
