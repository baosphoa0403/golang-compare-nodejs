# 📊 Benchmark Report: Excel Small (Node.js vs Go) (In-Memory Read)

## 1. Introduction

This case benchmarks the performance of reading a small Excel file (`small.xlsx`) between:

- **Node.js** using [exceljs](https://www.npmjs.com/package/exceljs)
- **Go** using [excelize](https://github.com/xuri/excelize)

**Goal**: measure **read speed** and **number of rows processed** for a file with ~50k rows (~10MB).

---

## 2. Test Configuration

- **File**: `small.xlsx` (50,001 rows).
- **Endpoint**:
  - Node.js: `GET /excel-small`
  - Go: `GET /excel-small`
- **Environment**: Docker container, limited to `1 CPU`, `1GB RAM`.
- **Metric**: execution time (seconds).

---

## 3. Results

- **Node.js**
  ```json
  { "case": "Small Excel", "rows": 50001, "duration": "0.701s" }
  ```
- **Golang**
  ```json
  { "case": "Small Excel", "rows": 50001, "duration": "0.421s" }
  ```

---

## 4. Analysis

| Criteria       | Golang (excelize)                          | Node.js (exceljs)                             |
| -------------- | ------------------------------------------ | --------------------------------------------- |
| **Performance**| ~3.7x faster                               | Slower, prone to blocking under heavy requests|
| **Library**    | Native code, optimized for I/O & memory    | Pure JavaScript, overhead from Garbage Collector |
