import { Run } from "@/models/run";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";

interface RunCardProps {
  run: Run;
}

const getStatusBadgeClass = (status: string) => {
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

const getStatusBadgeText = (status: string) => {
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

export function RunCard({ run }: RunCardProps) {
  const completedTests = run.run_tests.filter(
    (test) => test.status.toLowerCase() !== "in_progress",
  ).length;
  const totalTests = run.run_tests.length;
  const progress = totalTests > 0 ? (completedTests / totalTests) * 100 : 0;

  return (
    <Card className="h-full transition-all duration-200 hover:bg-neutral-900">
      <CardHeader className="pb-2">
        <div className="flex items-start justify-between">
          <CardTitle className="text-lg">{run.name}</CardTitle>
          <Card
            className={`ml-2 rounded-md border-none px-2 py-1 text-xs font-medium ${getStatusBadgeClass(run.status)}`}
          >
            {getStatusBadgeText(run.status)}
          </Card>
        </div>
      </CardHeader>
      <CardContent>
        <p className="text-md mb-2 text-muted-foreground">
          Commit: <code>{run.commit.substring(0, 7)}</code>
        </p>
        {run.status.toLowerCase() === "passed" ||
        run.status.toLowerCase() === "failed" ? (
          <div className="flex items-center text-sm text-muted-foreground">
            <span>Tests: {run.run_tests.length}</span>
            <span className="mx-2">•</span>
            <span>
              Finished:{" "}
              {new Date(
                run.run_tests.length > 0
                  ? Math.max(
                      ...run.run_tests.map((test) =>
                        new Date(test.created_at).getTime(),
                      ),
                    )
                  : new Date(run.created_at).getTime(),
              ).toLocaleString("en-US", {
                month: "numeric",
                day: "numeric",
                year: "numeric",
                hour: "numeric",
                minute: "2-digit",
                hour12: true,
              })}
            </span>
          </div>
        ) : (
          <div className="flex items-center text-sm text-muted-foreground">
            <span>
              Started:{" "}
              {new Date(run.created_at).toLocaleString("en-US", {
                month: "numeric",
                day: "numeric",
                year: "numeric",
                hour: "numeric",
                minute: "2-digit",
                hour12: true,
              })}
            </span>
          </div>
        )}
        {run.status.toLowerCase() === "in_progress" && (
          <div className="mt-4">
            <div className="mb-1 flex justify-between text-sm">
              <span className="text-muted-foreground">Test Progress</span>
              <span className="text-muted-foreground">
                {completedTests}/{totalTests} completed
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
  );
}
