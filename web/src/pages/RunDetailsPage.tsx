import React, { useState } from "react";
import Footer from "@/components/Footer";
import { AuthLoading } from "@/components/AuthLoading";
import { BACKEND_URL } from "@/consts/config";
import axios from "axios";
import { notify } from "@/lib/notify";
import { getAxiosErrorMessage } from "@/lib/axios-error-handler";
import { initRun, Run, RunTest } from "@/models/run";
import { useParams, useNavigate } from "react-router-dom";
import Header from "@/components/Header";
import { Button } from "@/components/ui/button";
import {
  ArrowLeft,
  Clock,
  AlertCircle,
  AlertTriangle,
  CheckCircle2,
  XCircle,
  ExternalLink,
  Hammer,
  ServerCog,
} from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { RunTestCard } from "@/components/RunTestCard";
import { Progress } from "@/components/ui/progress";

function RunDetailsPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [run, setRun] = useState<Run>(initRun);
  const [isCompleted, setIsCompleted] = useState(false);

  React.useEffect(() => {
    getRun();
    const interval = setInterval(() => {
      if (!isCompleted) {
        getRun();
      }
    }, 1000);
    return () => clearInterval(interval);
  }, [id, isCompleted]);

  const getRun = async () => {
    try {
      const response = await axios.get(`${BACKEND_URL}/runs/${id}`);
      setRun(response.data);
      if (
        response.data.status.toLowerCase() === "passed" ||
        response.data.status.toLowerCase() === "failed"
      ) {
        setIsCompleted(true);
      }
    } catch (error) {
      notify.error(getAxiosErrorMessage(error));
    }
  };

  const getRunStatusCardClass = (status: string) => {
    switch (status.toLowerCase()) {
      case "passed":
        return "bg-emerald-500 text-white";
      case "failed":
        return "bg-rose-500 text-white";
      case "in_progress":
        return "bg-gradient-to-br from-gr-pink to-gr-purple text-white";
      case "building":
        return "bg-blue-500 text-white";
      case "initializing":
        return "bg-yellow-500 text-white";
      default:
        return "bg-neutral-800 text-white";
    }
  };

  const getRunStatusText = (status: string) => {
    switch (status.toLowerCase()) {
      case "passed":
        return "Passed";
      case "failed":
        return "Failed";
      case "in_progress":
        return "In Progress";
      case "building":
        return "Building";
      case "initializing":
        return "Initializing";
      case "queued":
        return "Queued";
      default:
        return "Unknown";
    }
  };

  const groupTestsByStatus = (tests: RunTest[]) => {
    return {
      partial: tests.filter((test) => test.status.toLowerCase() === "partial"),
      failed: tests.filter((test) => test.status.toLowerCase() === "failed"),
      passed: tests.filter((test) => test.status.toLowerCase() === "passed"),
      inProgress: tests.filter(
        (test) => test.status.toLowerCase() === "in_progress",
      ),
    };
  };

  const formatDate = (date: Date) => {
    return new Date(date).toLocaleString("en-US", {
      month: "numeric",
      day: "numeric",
      year: "numeric",
      hour: "numeric",
      minute: "2-digit",
      hour12: true,
    });
  };

  const isRunCompleted = () =>
    run.status === "passed" || run.status === "failed";

  const completedTests = run.run_tests.filter(
    (test) => test.status.toLowerCase() !== "in_progress",
  ).length;
  const totalTests = run.run_tests.length;
  const progress = totalTests > 0 ? (completedTests / totalTests) * 100 : 0;

  const getElapsedTime = () => {
    const startTime = new Date(run.created_at);
    const currentTime = new Date();
    const durationMs = currentTime.getTime() - startTime.getTime();
    if (durationMs < 1000) {
      return `${durationMs}ms`;
    } else if (durationMs < 60000) {
      return `${(durationMs / 1000).toFixed(0)}s`;
    } else {
      const minutes = Math.floor(durationMs / 60000);
      const seconds = ((durationMs % 60000) / 1000).toFixed(0);
      return `${minutes}m ${seconds}s`;
    }
  };

  const getCompletedTime = () => {
    const startTime = new Date(run.created_at);
    const timestamps = run.run_tests.flatMap((test) => [
      new Date(test.created_at).getTime(),
      ...test.run_test_results.map((result) =>
        new Date(result.created_at).getTime(),
      ),
    ]);
    const endTime = Math.max(...timestamps);
    if (timestamps.length === 0) {
      return "N/A";
    }
    const durationMs = endTime - startTime.getTime();
    if (durationMs < 1000) {
      return `${durationMs}ms`;
    } else if (durationMs < 60000) {
      return `${(durationMs / 1000).toFixed(2)}s`;
    } else {
      const minutes = Math.floor(durationMs / 60000);
      const seconds = ((durationMs % 60000) / 1000).toFixed(2);
      return `${minutes}m ${seconds}s`;
    }
  };

  const RunQueuedCard = () => {
    return (
      <Card className="col-span-full">
        <CardContent className="flex flex-col items-center justify-center space-y-4 pt-6">
          <Clock className="h-12 w-12 text-muted-foreground" />
          <CardTitle className="text-center">Run Queued</CardTitle>
          <p className="text-center text-muted-foreground">
            The run is queued. It will start automatically when there is an
            available runner.
          </p>
        </CardContent>
      </Card>
    );
  };

  const RunBuildingCard = () => {
    return (
      <Card className="col-span-full">
        <CardContent className="flex flex-col items-center justify-center space-y-4 pt-6">
          <Hammer className="h-12 w-12 text-blue-500" />
          <CardTitle className="text-center">Building Service</CardTitle>
          <p className="text-center text-muted-foreground">
            The requetested service is being built. It will be pushed to the
            Docker registry when complete.
          </p>
        </CardContent>
      </Card>
    );
  };

  const RunInitializingCard = () => {
    return (
      <Card className="col-span-full">
        <CardContent className="flex flex-col items-center justify-center space-y-4 pt-6">
          <ServerCog className="h-12 w-12 text-yellow-500" />
          <CardTitle className="text-center">Initializing Containers</CardTitle>
          <p className="text-center text-muted-foreground">
            The neccesary test containers are being initialized. This may take a
            few minutes.
          </p>
        </CardContent>
      </Card>
    );
  };

  const NoTestsCard = () => {
    return (
      <Card className="col-span-full">
        <CardContent className="flex flex-col items-center justify-center space-y-4 pt-6">
          <AlertCircle className="h-12 w-12 text-muted-foreground" />
          <CardTitle className="text-center">No Results</CardTitle>
          <p className="text-center text-muted-foreground">
            There are no results available for this run. There might have been a
            problem trying to build the request service, or during
            initialization of the test containers. See the production logs on{" "}
            <a
              href="https://portainer.gauchoracing.com"
              target="_blank"
              rel="noopener noreferrer"
              className="text-gr-pink hover:underline"
            >
              Portainer
            </a>{" "}
            for more information.
          </p>
        </CardContent>
      </Card>
    );
  };

  return (
    <>
      {run.id == "" ? (
        <AuthLoading />
      ) : (
        <div className="flex h-screen flex-col justify-start">
          <Header />
          <div className="flex flex-col justify-start p-4 lg:p-32 lg:pt-8">
            <div className="mb-4">
              <Button
                variant={"ghost"}
                onClick={() => navigate("/")}
                className="flex items-center"
              >
                <ArrowLeft className="mr-2 h-4 w-4 text-gray-400" />
                Back to home
              </Button>
            </div>
            <Card className="mb-6">
              <CardHeader>
                <div className="flex items-center justify-between">
                  <CardTitle className="text-2xl">{run.name}</CardTitle>
                  <Card
                    className={`ml-2 rounded-md border-none px-2 py-1 text-xs font-medium ${getRunStatusCardClass(run.status)}`}
                  >
                    {getRunStatusText(run.status)}
                  </Card>
                </div>
              </CardHeader>
              <CardContent>
                <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
                  <div>
                    <p className="text-sm text-muted-foreground">Service</p>
                    <p className="text-lg">{run.service}</p>
                  </div>
                  <div>
                    <p className="text-sm text-muted-foreground">Commit</p>
                    <div className="flex flex-wrap items-center gap-2 overflow-clip">
                      <p className="font-mono text-lg">{run.commit}</p>
                      <a
                        href={`https://github.com/gaucho-racing/mapache/commit/${run.commit}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-gr-pink hover:text-gr-pink/80"
                      >
                        <ExternalLink className="h-4 w-4" />
                      </a>
                    </div>
                  </div>
                  <div>
                    <p className="text-sm text-muted-foreground">Created At</p>
                    <p className="text-lg">{formatDate(run.created_at)}</p>
                  </div>
                  <div>
                    <p className="text-sm text-muted-foreground">
                      GitHub Check Run ID
                    </p>
                    <div className="flex flex-wrap items-center gap-2 overflow-clip">
                      <p className="font-mono text-lg">
                        {run.github_check_run_id}
                      </p>
                      <a
                        href={`https://github.com/gaucho-racing/mapache/runs/${run.github_check_run_id}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-gr-pink hover:text-gr-pink/80"
                      >
                        <ExternalLink className="h-4 w-4" />
                      </a>
                    </div>
                  </div>
                  {isRunCompleted() && (
                    <div>
                      <p className="text-sm text-muted-foreground">
                        Elapsed Time
                      </p>
                      <p className="text-lg">{getCompletedTime()}</p>
                    </div>
                  )}
                </div>
                {run.status.toLowerCase() === "in_progress" && (
                  <div className="mt-4">
                    <div className="mb-2 flex items-end justify-between text-sm">
                      <div>
                        <p className="text-sm text-muted-foreground">
                          Test Progress
                        </p>
                        <p className="text-lg">
                          {completedTests}/{totalTests} completed
                        </p>
                      </div>
                      <span className="text-muted-foreground">
                        Elapsed time: {getElapsedTime()}
                      </span>
                    </div>
                    <Progress
                      value={progress}
                      className="h-2 [&>div]:bg-gradient-to-r [&>div]:from-gr-pink [&>div]:to-gr-purple"
                    />
                  </div>
                )}
              </CardContent>
            </Card>
            {groupTestsByStatus(run.run_tests).inProgress.length > 0 && (
              <>
                <div className="mb-4 flex items-center gap-2">
                  <Clock className="h-8 w-8 text-gr-pink" />
                  <h2 className="text-2xl font-semibold">In Progress Tests</h2>
                </div>
                <div className="mb-4">
                  {run.run_tests
                    .filter(
                      (result) => result.status.toLowerCase() === "in_progress",
                    )
                    .map((result) => (
                      <div key={result.id} className="mb-4">
                        <RunTestCard test={result} />
                      </div>
                    ))}
                </div>
              </>
            )}
            {groupTestsByStatus(run.run_tests).partial.length > 0 && (
              <>
                <div className="mb-4 flex items-center gap-2">
                  <AlertTriangle className="h-8 w-8 text-amber-500" />
                  <h2 className="text-2xl font-semibold">
                    Partially Failed Tests
                  </h2>
                </div>
                <div className="mb-4">
                  {run.run_tests
                    .filter(
                      (result) => result.status.toLowerCase() === "partial",
                    )
                    .map((result) => (
                      <div key={result.id} className="mb-4">
                        <RunTestCard test={result} />
                      </div>
                    ))}
                </div>
              </>
            )}
            {groupTestsByStatus(run.run_tests).failed.length > 0 && (
              <>
                <div className="mb-4 flex items-center gap-2">
                  <XCircle className="h-8 w-8 text-red-500" />
                  <h2 className="text-2xl font-semibold">Failed Tests</h2>
                </div>
                <div className="mb-4">
                  {run.run_tests
                    .filter(
                      (result) => result.status.toLowerCase() === "failed",
                    )
                    .map((result) => (
                      <div key={result.id} className="mb-4">
                        <RunTestCard test={result} />
                      </div>
                    ))}
                </div>
              </>
            )}
            {groupTestsByStatus(run.run_tests).passed.length > 0 && (
              <>
                <div className="mb-4 flex items-center gap-2">
                  <CheckCircle2 className="h-8 w-8 text-green-500" />
                  <h2 className="text-2xl font-semibold">Passed Tests</h2>
                </div>
                <div className="mb-4">
                  {run.run_tests
                    .filter(
                      (result) => result.status.toLowerCase() === "passed",
                    )
                    .map((result) => (
                      <div key={result.id} className="mb-4">
                        <RunTestCard test={result} />
                      </div>
                    ))}
                </div>
              </>
            )}
            {isRunCompleted() && run.run_tests.length === 0 && <NoTestsCard />}
            {run.status == "queued" && <RunQueuedCard />}
            {run.status == "building" && <RunBuildingCard />}
            {run.status == "initializing" && <RunInitializingCard />}
          </div>
          <Footer />
        </div>
      )}
    </>
  );
}

export default RunDetailsPage;
