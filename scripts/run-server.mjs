/**
 * Dev server launcher (same idea as GoCMS scripts/run-server.mjs).
 *
 * - Runs `go run ./cmd/server` from the repository root so APP_STATIC_DIR / web/static resolve.
 * - stdio: inherit so Ctrl+C in the terminal reaches the Go toolchain and server.
 * - Forwards SIGINT/SIGTERM to the child so the listener is released (port freed).
 *
 * Closing a browser tab does not stop an HTTP server; stop this terminal job (Ctrl+C)
 * or close the whole terminal panel so the process tree exits.
 */
import { spawn } from "node:child_process";
import { fileURLToPath } from "node:url";
import path from "node:path";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const root = path.resolve(__dirname, "..");

const child = spawn("go", ["run", "./cmd/server"], {
  cwd: root,
  stdio: "inherit",
  env: { ...process.env },
  windowsHide: true,
});

function forwardStop() {
  if (child.exitCode !== null || child.signalCode) {
    return;
  }
  if (process.platform === "win32") {
    try {
      child.kill();
    } catch {
      /* ignore */
    }
    return;
  }
  child.kill("SIGTERM");
}

for (const sig of ["SIGINT", "SIGTERM", "SIGHUP"]) {
  process.on(sig, forwardStop);
}

child.on("exit", (code, signal) => {
  if (signal) {
    process.exit(1);
  }
  process.exit(code ?? 0);
});
