# 📊 Benchmark Report: Excel Medium (Node.js vs Go) (In-Memory Read)

## 1. Introduction

This case benchmarks the performance of reading a medium Excel file (`medium.xlsx`) between:

- **Node.js** using [exceljs](https://www.npmjs.com/package/exceljs)
- **Go** using [excelize](https://github.com/xuri/excelize)

**Goal**: measure **read speed** and **number of rows processed** for a file with ~1 million rows (size ~200–300MB).

---

## 2. Test Configuration

- **File**: `medium.xlsx` (1,000,001 rows).
- **Endpoint**:
  - Node.js: `GET /excel-medium`
  - Go: `GET /excel-medium`
- **Environment**: Docker container, limited to `2 CPU`, `2GB RAM`.
- **Metric**: execution time (seconds).

---

## 3. Results

- **Node.js**
  ⛔ App crashed, no response returned (OOM when loading workbook).

- **Golang**
  ```json
  { "case": "Medium Excel", "rows": 1000001, "duration": "9.257096546s" }
  ```
