import { RunTest } from "@/models/run";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { Clock, ChevronDown, CheckCircle, XCircle } from "lucide-react";
import { useState } from "react";
import { cn } from "@/lib/utils";

interface RunTestCardProps {
  test: RunTest;
  defaultExpanded?: boolean;
}

export function RunTestCard({
  test,
  defaultExpanded = false,
}: RunTestCardProps) {
  const [isExpanded, setIsExpanded] = useState(defaultExpanded);

  const getTestStatusCardClass = (status: string) => {
    switch (status.toLowerCase()) {
      case "passed":
        return "bg-emerald-500 text-white";
      case "failed":
        return "bg-rose-500 text-white";
      case "in_progress":
        return "bg-gradient-to-br from-gr-pink to-gr-purple text-white";
      case "partial":
        return "bg-amber-500 text-white";
      default:
        return "bg-neutral-800 text-white";
    }
  };

  const getTestStatusText = (status: string) => {
    switch (status.toLowerCase()) {
      case "passed":
        return "Passed";
      case "failed":
        return "Failed";
      case "in_progress":
        return "In Progress";
      case "partial":
        return "Partial";
      default:
        return "Unknown";
    }
  };

  const getTestResultIcon = (status: string) => {
    switch (status.toLowerCase()) {
      case "passed":
        return <CheckCircle className="h-5 w-5 text-emerald-500" />;
      case "failed":
        return <XCircle className="h-5 w-5 text-rose-500" />;
      case "partial":
        return <XCircle className="h-5 w-5 text-amber-500" />;
      default:
        return <XCircle className="h-5 w-5 text-neutral-500" />;
    }
  };

  const calculateTestDuration = () => {
    if (test.run_test_results.length === 0) return "0s";

    const timestamps = test.run_test_results.map((result) =>
      new Date(result.created_at).getTime(),
    );

    const startTime = new Date(test.created_at).getTime();
    const endTime = Math.max(...timestamps);
    const durationMs = endTime - startTime;

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

  return (
    <Card key={test.id}>
      <CardHeader
        className="cursor-pointer rounded-md hover:bg-neutral-900"
        onClick={() => setIsExpanded(!isExpanded)}
      >
        <div className="flex items-center justify-between">
          <div>
            <div className="flex items-center gap-2">
              <CardTitle className="text-lg">{test.name}</CardTitle>
              <ChevronDown
                className={cn(
                  "h-4 w-4 transition-transform duration-200",
                  isExpanded && "rotate-180",
                )}
              />
            </div>
            <div className="mt-1 flex items-center text-sm text-muted-foreground">
              <Clock className="mr-1 h-3 w-3" />
              Completed in {calculateTestDuration()}
            </div>
          </div>
          <Card
            className={`ml-2 rounded-md border-none px-2 py-1 text-xs font-medium ${getTestStatusCardClass(test.status)}`}
          >
            {getTestStatusText(test.status)}
          </Card>
        </div>
      </CardHeader>
      <div
        className={cn(
          "overflow-hidden transition-all duration-200",
          isExpanded ? "" : "max-h-0",
        )}
      >
        <CardContent>
          <div className="space-y-4">
            {test.run_test_results
              .sort((a, b) => {
                if (
                  a.status.toLowerCase() === "failed" &&
                  b.status.toLowerCase() !== "failed"
                )
                  return -1;
                if (
                  a.status.toLowerCase() !== "failed" &&
                  b.status.toLowerCase() === "failed"
                )
                  return 1;
                return 0;
              })
              .map((result) => (
                <div key={result.id}>
                  <div className="flex items-center gap-2">
                    {getTestResultIcon(result.status)}
                    <p className="font-medium">{result.signal_name}</p>
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
                  {test.run_test_results.indexOf(result) !==
                    test.run_test_results.length - 1 && (
                    <Separator className="my-4" />
                  )}
                </div>
              ))}
          </div>
        </CardContent>
      </div>
    </Card>
  );
}
