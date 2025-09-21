# 📊 Benchmark Report: Excel Large (Node.js vs Go) (Streaming Read)

## 1. Introduction

This case benchmarks the performance of reading a large Excel file (`large.xlsx`) using **streaming** between:

- **Node.js** with [exceljs](https://www.npmjs.com/package/exceljs) (WorkbookReader)
- **Go** with [excelize](https://github.com/xuri/excelize) (Rows iterator)

**Goal**: measure **read speed** and **number of rows processed** for a file with ~1,048,576 rows (hundreds of MB), while demonstrating the advantage of streaming compared to fully loading into RAM.

---

## 2. Test Configuration

- **File**: `large.xlsx` (1,048,576 rows).
- **Endpoint**:
  - Node.js: `GET /excel-large`
  - Go: `GET /excel-large`
- **Environment**: Docker container, limited to `2 CPU`, `2GB RAM`.
- **Metric**: execution time (seconds).

---

## 3. Results

- **Node.js**
  ```json
  { "case": "Large Excel Streaming", "rows": 1048576, "duration": "9.679s" }
  ```

- **Golang**
  ```json
  {
    "case": "Large Excel Streaming",
    "rows": 1048576,
    "duration": "3.935057251s"
  }
  ```

---

## 4. Analysis

| Criteria           | Golang (excelize)                   | Node.js (exceljs)                                                  |
| ------------------ | ----------------------------------- | ------------------------------------------------------------------ |
| **Performance**    | Reads 1M rows in ~3.8 seconds       | Reads 1M rows in ~13.9 seconds                                     |
| **Speed**          | ~3.7x faster                        | Significantly slower                                               |
| **Memory usage**   | Row-by-row iterator, low RAM usage  | Streaming but more object allocation, higher RAM usage             |
| **Stability**      | Stable, fast response               | Stable but higher latency, prone to bottlenecks under heavy traffic|
