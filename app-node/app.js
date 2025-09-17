import express from "express";
import { Worker } from "worker_threads";
import bcrypt from "bcrypt";

const app = express();
const port = 3000;

// Benchmark hashing password
const hashPassword = async (password) => {
  return await bcrypt.hash(password, 10);
};

app.get("/me", (req, res) => {
  res.send({
    result: "hello gia bao 123",
  });
});

app.get("/hash-password-sync", (req, res) => {
  const start = Date.now();

  const password = "thisIsASecurePassword123"; // Password để hash

  // Tác vụ hashing password sẽ block main thread cho đến khi hoàn thành
  const hashedPassword = hashPassword(password);

  const end = Date.now();
  const duration = end - start;

  res.send({
    result: "Hash completed",
    hashedPassword: hashedPassword, // Return hashed password
    duration: `${duration}ms`, // Thời gian thực hiện
  });
});

// API để benchmark hashing password
app.get("/hash-password", (req, res) => {
  const start = Date.now();
  const worker = new Worker("./worker.js");

  worker.on("message", (message) => {
    const end = Date.now();
    const duration = end - start;
    console.log("generate pwd", message);

    res.send({ result: "Hash completed", duration: `${duration}ms` });
  });

  worker.on("error", (err) => {
    res.status(500).send({ error: err.message });
  });

  worker.postMessage("start");
});

app.listen(port, () => {
  console.log(`Node.js server is running at http://localhost:${port}`);
});
