import React from "react";
import { useNavigate } from "react-router-dom";
import { checkCredentials } from "@/lib/auth";
import Footer from "@/components/Footer";
import { AuthLoading } from "@/components/AuthLoading";
import { useUser, useRuns, setRuns } from "@/lib/store";
import Header from "./components/Header";
import { BACKEND_URL } from "@/consts/config";
import axios from "axios";
import { notify } from "@/lib/notify";
import { getAxiosErrorMessage } from "@/lib/axios-error-handler";
import { Run } from "@/models/run";
import { ScrollArea } from "@/components/ui/scroll-area";
import { RunCard } from "@/components/RunCard";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Clock } from "lucide-react";
import { NoActiveTestsCard } from "@/components/NoActiveTestsCard";

function App() {
  const navigate = useNavigate();
  const runs = useRuns();
  const currentUser = useUser();

  React.useEffect(() => {
    getRuns();
    const interval = setInterval(getRuns, 2000);
    return () => clearInterval(interval);
  }, []);

  const getRuns = async () => {
    try {
      const response = await axios.get(`${BACKEND_URL}/runs`, {});
      setRuns(response.data);
    } catch (error) {
      notify.error(getAxiosErrorMessage(error));
    }
  };

  const inProgressRuns = runs.filter(
    (run: Run) => run.status !== "passed" && run.status !== "failed",
  );

  return (
    <>
      {runs.length == 0 ? (
        <AuthLoading />
      ) : (
        <div className="flex h-screen flex-col justify-between">
          <Header />
          <div className="flex flex-col justify-start p-4 lg:p-32 lg:pt-16">
            <div className="flex flex-row items-center justify-between">
              <h2>In Progress</h2>
            </div>
            <div className="mt-4 grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
              {inProgressRuns.length > 0 ? (
                inProgressRuns.map((run: Run) => (
                  <RunCard key={run.id} run={run} />
                ))
              ) : (
                <NoActiveTestsCard />
              )}
            </div>
            <div className="flex flex-row items-center justify-between pt-8">
              <h2>Completed</h2>
            </div>
            <div className="mt-4 grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
              {runs
                .filter(
                  (run: Run) =>
                    run.status === "passed" || run.status === "failed",
                )
                .map((run: Run) => (
                  <RunCard key={run.id} run={run} />
                ))}
            </div>
          </div>
          <Footer />
        </div>
      )}
    </>
  );
}

export default App;
