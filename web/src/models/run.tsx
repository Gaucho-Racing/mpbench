export interface Run {
  id: string;
  name: string;
  service: string;
  commit: string;
  status: string;
  github_check_run_id: number;
  run_tests: RunTest[];
  created_at: Date;
}

export interface RunTest {
  id: string;
  run_id: string;
  name: string;
  status: string;
  data: string;
  run_test_results: RunTestResult[];
  created_at: Date;
}

export interface RunTestResult {
  id: string;
  run_test_id: string;
  signal_name: string;
  status: string;
  value: string;
  expected: string;
  created_at: Date;
}

export const initRun: Run = {
  id: "",
  name: "",
  service: "",
  commit: "",
  status: "",
  github_check_run_id: 0,
  run_tests: [],
  created_at: new Date(),
};

export const initRunTest: RunTest = {
  id: "",
  run_id: "",
  name: "",
  status: "",
  data: "",
  run_test_results: [],
  created_at: new Date(),
};

export const initRunTestResult: RunTestResult = {
  id: "",
  run_test_id: "",
  signal_name: "",
  status: "",
  value: "",
  expected: "",
  created_at: new Date(),
};
