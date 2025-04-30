import { Card, CardContent, CardTitle } from "@/components/ui/card";
import { Clock } from "lucide-react";

export function NoActiveTestsCard() {
  return (
    <Card className="col-span-full">
      <CardContent className="flex flex-col items-center justify-center space-y-4 pt-6">
        <Clock className="h-12 w-12 text-muted-foreground" />
        <CardTitle className="text-center">No Active Tests</CardTitle>
        <p className="text-center text-muted-foreground">
          There are no tests queued or currently running. You can trigger a new
          test run by pushing a commit to{" "}
          <a
            href="https://github.com/gaucho-racing/mapache"
            target="_blank"
            rel="noopener noreferrer"
            className="text-gr-pink hover:underline"
          >
            github.com/gaucho-racing/mapache
          </a>
          .
        </p>
      </CardContent>
    </Card>
  );
}
