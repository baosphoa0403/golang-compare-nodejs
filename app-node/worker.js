import { parentPort } from "worker_threads";
import bcrypt from "bcrypt";

parentPort.on("message", (taskData) => {
  const password = taskData.password;
  bcrypt.hash(password, 10, (err, hashedPassword) => {
    if (err) {
      parentPort.postMessage({ error: err });
    } else {
      parentPort.postMessage({ hashedPassword });
    }
  });
});
