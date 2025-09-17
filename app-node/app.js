import express from "express";
import bcrypt from "bcrypt";
import WorkerPool from "./workerPool.js";

const app = express();
const port = 3000;

const pool = new WorkerPool("./worker.js", 4); // 4 worker trong pool
pool.initialize();

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

  const password = "thisIsASecurePassword123"; // Password để hash

  // Sử dụng Worker Pool để xử lý công việc
  pool
    .runTask({ password })
    .then((result) => {
      const end = Date.now();
      const duration = end - start;
      // console.log("Generated password", result.hashedPassword);

      res.send({ result: "Hash completed", duration: `${duration}ms` });
    })
    .catch((err) => {
      res.status(500).send({ error: err.message });
    });
});

app.listen(port, () => {
  console.log(`Node.js server is running at http://localhost:${port}`);
});
