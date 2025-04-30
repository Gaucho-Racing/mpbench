import React, { useState } from "react";
import Footer from "@/components/Footer";
import { AuthLoading } from "@/components/AuthLoading";
import { useRuns, setRuns, getRuns } from "@/lib/store";
import { BACKEND_URL } from "@/consts/config";
import axios from "axios";
import { notify } from "@/lib/notify";
import { getAxiosErrorMessage } from "@/lib/axios-error-handler";
import { initRun, Run } from "@/models/run";
import { RunCard } from "@/components/RunCard";
import { NoActiveTestsCard } from "@/components/NoActiveTestsCard";
import { useParams, useNavigate } from "react-router-dom";
import Header from "@/components/Header";
import { Button } from "@/components/ui/button";
import { ArrowLeft, Check, X, Clock, AlertCircle } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { Badge } from "@/components/ui/badge";

function RunDetailsPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [run, setRun] = useState<Run>(initRun);

  React.useEffect(() => {
    getRun();
    const interval = setInterval(getRun, 1000);
    return () => clearInterval(interval);
  }, [id]);

  const getRun = async () => {
    try {
      const response = await axios.get(`${BACKEND_URL}/runs/${id}`);
      setRun(response.data);
    } catch (error) {
      notify.error(getAxiosErrorMessage(error));
    }
  };

  const getStatusCardClass = (status: string) => {
    switch (status.toLowerCase()) {
      case "passed":
        return "bg-emerald-500 text-white";
      case "failed":
        return "bg-rose-500 text-white";
      case "in_progress":
        return "bg-gradient-to-br from-gr-pink to-gr-purple text-white";
      case "building":
        return "bg-slate-500 text-white";
      case "initializing":
        return "bg-yellow-500 text-white";
      default:
        return "bg-neutral-800 text-white";
    }
  };

  const getStatusText = (status: string) => {
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
      default:
        return "Unknown";
    }
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

  const isRunInProgress = () => run.status == "in_progress";

  return (
    <>
      {run.id == "" ? (
        <AuthLoading />
      ) : (
        <div className="flex h-screen flex-col justify-between">
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
                    className={`ml-2 rounded-md border-none px-2 py-1 text-xs font-medium ${getStatusCardClass(run.status)}`}
                  >
                    {getStatusText(run.status)}
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
                    <p className="font-mono text-lg">{run.commit}</p>
                  </div>
                  <div>
                    <p className="text-sm text-muted-foreground">Created At</p>
                    <p className="text-lg">{formatDate(run.created_at)}</p>
                  </div>
                  <div>
                    <p className="text-sm text-muted-foreground">
                      GitHub Check Run ID
                    </p>
                    <p className="text-lg">{run.github_check_run_id}</p>
                  </div>
                </div>
              </CardContent>
            </Card>

            <div className="mb-4">
              <h2 className="text-2xl font-semibold">Tests</h2>
            </div>
            <div className="space-y-4">
              {run.run_tests.map((test) => (
                <Card key={test.id}>
                  <CardHeader>
                    <div className="flex items-center justify-between">
                      <CardTitle className="text-lg">{test.name}</CardTitle>
                      <Card
                        className={`ml-2 rounded-md border-none px-2 py-1 text-xs font-medium ${getStatusCardClass(test.status)}`}
                      >
                        {getStatusText(test.status)}
                      </Card>
                    </div>
                  </CardHeader>
                  <CardContent>
                    <div className="space-y-4">
                      {test.run_test_results.map((result) => (
                        <div key={result.id}>
                          <div className="flex items-center justify-between">
                            <p className="font-medium">{result.signal_name}</p>
                            <Card
                              className={`ml-2 rounded-md border-none px-2 py-1 text-xs font-medium ${getStatusCardClass(result.status)}`}
                            >
                              {getStatusText(result.status)}
                            </Card>
                          </div>
                          <div className="mt-2 grid grid-cols-2 gap-4 text-sm">
                            <div>
                              <p className="text-muted-foreground">Value</p>
                              <p className="font-mono">{result.value}</p>
                            </div>
                            <div>
                              <p className="text-muted-foreground">Expected</p>
                              <p className="font-mono">{result.expected}</p>
                            </div>
                          </div>
                          <Separator className="my-4" />
                        </div>
                      ))}
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          </div>
          <Footer />
        </div>
      )}
    </>
  );
}

export default RunDetailsPage;
