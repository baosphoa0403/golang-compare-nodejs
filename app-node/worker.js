import { parentPort } from "worker_threads";
import bcrypt from "bcrypt";

parentPort.on("message", async (msg) => {
  if (msg === "start") {
    const password = "thisIsASecurePassword123"; // Password để hash
    // Simulate CPU-bound task: Hash password
    const hashedPassword = await bcrypt.hash(password, 10);

    parentPort.postMessage(hashedPassword);
  }
});
