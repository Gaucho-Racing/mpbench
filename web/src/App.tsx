import React from "react";
import Footer from "@/components/Footer";
import { AuthLoading } from "@/components/AuthLoading";
import { useRuns, setRuns } from "@/lib/store";
import Header from "./components/Header";
import { BACKEND_URL } from "@/consts/config";
import axios from "axios";
import { notify } from "@/lib/notify";
import { getAxiosErrorMessage } from "@/lib/axios-error-handler";
import { Run } from "@/models/run";
import { RunCard } from "@/components/RunCard";
import { NoActiveTestsCard } from "@/components/NoActiveTestsCard";
import { useNavigate } from "react-router-dom";

function App() {
  const navigate = useNavigate();
  const runs = useRuns();

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
                  <div key={run.id} onClick={() => navigate(`/runs/${run.id}`)}>
                    <RunCard run={run} />
                  </div>
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
                  <div key={run.id} onClick={() => navigate(`/runs/${run.id}`)}>
                    <RunCard run={run} />
                  </div>
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
