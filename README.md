# 📊 Compare Node.js vs Go for Excel Import

## 1. Introduction
This project benchmarks and compares the performance of **Node.js (exceljs)** and **Go (excelize)** when importing Excel files of different sizes.

The purpose is to evaluate:
- **Execution speed** (time to read/parse files).
- **Memory and CPU usage**.
- **Stability** under large datasets.
- **Concurrency** when handling requests during heavy import tasks.

---

## 2. Project Structure
- **app-node/** → Node.js implementation (Express + exceljs + worker threads).  
- **app-golang/** → Go implementation (net/http + excelize).  
- **data/** → Test Excel files (`small.xlsx`, `medium.xlsx`, `large.xlsx`, `import.xlsx`).  
- **docker-compose.yml** → Unified environment for running Node.js, Go, and PostgreSQL.  
- **Benchmarks**:
  - `memory-excel-small-README-EN.md`
  - `memory-excel-medium-README-EN.md`
  - `stream-excel-large-README-EN.md`

---

## 3. Test Cases
1. **Small Excel (In-Memory)**  
   - ~50k rows (~10MB).  
   - Both Node.js and Go read into RAM.  
   - Go is ~1.7x faster.  

2. **Medium Excel (In-Memory)**  
   - ~1M rows (~200–300MB).  
   - Node.js crashed (OOM).  
   - Go succeeded (~9.2s).  

3. **Large Excel (Streaming)**  
   - ~1M rows, hundreds of MB.  
   - Both used streaming mode.  
   - Go: ~3.8s vs Node.js: ~13.9s.  

---

## 4. Environment
- **Dockerized setup**:
  - Node.js 20
  - Go 1.23+
  - PostgreSQL 15 (for validation + insert tests)
- Resource limits:
  - Baseline: `--cpus=1`, `--memory=2g`
  - Stress test: `--cpus=2`, `--memory=2g`

---

## 5. Results Summary

| Case                  | Node.js (exceljs)         | Go (excelize)             | Result |
|------------------------|---------------------------|---------------------------|--------|
| Small (50k rows)       | ~0.70s                   | ~0.42s                   | Go ~1.7x faster |
| Medium (1M rows)       | ❌ Crash (OOM)            | ~9.25s                   | Go stable |
| Large (1M rows stream) | ~13.94s                  | ~3.80s                   | Go ~3.7x faster |

---

## 6. Conclusion
- For **small imports**, both Node.js and Go are viable, though Go is faster.  
- For **medium to large imports**, **Go clearly outperforms Node.js** in terms of speed and memory stability.  
- **Streaming** helps both, but Go remains more efficient.  
- For production systems requiring **large data ingestion + concurrency**, Go is the recommended choice.  
