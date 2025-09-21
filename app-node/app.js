import express from "express";
import bcrypt from "bcrypt";
import WorkerPool from "./workerPool.js";
import ExcelJS from "exceljs";
// import pkg from "pg";

// const { Pool } = pkg;
const app = express();
const port = 3000;

const pool = new WorkerPool("./worker.js", 4);
pool.initialize();

// ========== CASE 0: Hello ==========
app.get("/me", (req, res) => {
  setTimeout(() => {
    console.log("set timeout 0s");
  }, 0);
  res.send({ result: "hello gia bao 123" });
});

// ========== CASE 1: Hash password sync ==========
app.get("/hash-password-sync", async (req, res) => {
  const start = Date.now();
  const password = "thisIsASecurePassword123";
  const hashedPassword = await bcrypt.hash(password, 10);
  const end = Date.now();

  res.send({
    result: "Hash completed (sync)",
    hashedPassword,
    duration: `${end - start}ms`,
  });
});

// ========== CASE 2: Hash password with worker pool ==========
app.get("/hash-password", (req, res) => {
  const start = Date.now();
  const password = "thisIsASecurePassword123";

  pool
    .runTask({ password })
    .then((result) => {
      const end = Date.now();
      res.send({
        result: "Hash completed (worker pool)",
        duration: `${end - start}ms`,
      });
    })
    .catch((err) => res.status(500).send({ error: err.message }));
});

// ========== CASE 3: Read small Excel ==========
app.get("/excel-small", async (req, res) => {
  const start = Date.now();
  const workbook = new ExcelJS.Workbook();
  await workbook.xlsx.readFile("./data/small.xlsx");
  const worksheet = workbook.worksheets[0];
  res.send({
    case: "Small Excel",
    rows: worksheet.rowCount,
    duration: `${Date.now() - start}ms`,
  });
});

// ========== CASE 4: Read medium Excel ==========
app.get("/excel-medium", async (req, res) => {
  const start = Date.now();
  const workbook = new ExcelJS.Workbook();
  await workbook.xlsx.readFile("./data/medium.xlsx");
  const worksheet = workbook.worksheets[0];
  res.send({
    case: "Medium Excel",
    rows: worksheet.rowCount,
    duration: `${Date.now() - start}ms`,
  });
});

// ========== CASE 5: Stream large Excel ==========
app.get("/excel-large", async (req, res) => {
  const start = Date.now();
  let count = 0;

  const workbookReader = new ExcelJS.stream.xlsx.WorkbookReader(
    "./data/large.xlsx"
  );
  for await (const worksheetReader of workbookReader) {
    for await (const row of worksheetReader) {
      count++;
    }
  }

  res.send({
    case: "Large Excel Streaming",
    rows: count,
    duration: `${Date.now() - start}ms`,
  });
});

// ========== CASE 6: Validate + Insert DB ==========
// const pgPool = new Pool({
//   connectionString: "postgres://user:pass@localhost:5432/benchmark",
// });

// app.get("/excel-validate-insert", async (req, res) => {
//   const start = Date.now();
//   const workbook = new ExcelJS.Workbook();
//   await workbook.xlsx.readFile("./data/import.xlsx");
//   const worksheet = workbook.worksheets[0];
//   const client = await pgPool.connect();

//   let inserted = 0;
//   for (let i = 1; i <= worksheet.rowCount; i++) {
//     const row = worksheet.getRow(i);
//     const name = row.getCell(1).text;
//     const email = row.getCell(2).text;

//     if (!email.includes("@")) continue;
//     await client.query("INSERT INTO users (name, email) VALUES ($1, $2)", [
//       name,
//       email,
//     ]);
//     inserted++;
//   }

//   client.release();
//   res.send({
//     case: "Validate + Insert",
//     inserted,
//     duration: `${Date.now() - start}ms`,
//   });
// });

// ========== Start server ==========
app.listen(port, () => {
  console.log(`Server is running at http://localhost:${port}`);
});
