# 📊 Benchmark Report: Excel Small (Node.js vs Go)

## 1. Giới thiệu

Case này kiểm tra hiệu năng đọc file Excel nhỏ (`small.xlsx`) giữa:

- **Node.js** với thư viện [exceljs](https://www.npmjs.com/package/exceljs)
- **Go** với thư viện [excelize](https://github.com/xuri/excelize)

Mục tiêu: đo **tốc độ đọc** và **số dòng xử lý** cho file có ~50k rows (~10MB).

---

## 2. Cấu hình test

- **File**: `small.xlsx` (50,001 dòng).
- **Endpoint gọi**:
  - Node.js: `GET /excel-small`
  - Go: `GET /excel-small`
- **Môi trường**: Docker container, giới hạn `1 CPU`, `1GB RAM`.
- **Đo lường**: thời gian thực thi (ms).

---

## 3. Kết quả

- **Node.js**
  ```json
  { "case": "Small Excel", "rows": 50001, "duration": "1890ms" }
  ```
- **Golang**
  ```json
  { "case": "Small Excel", "rows": 50001, "duration": "514.636292ms" }
  ```
## 4. Phân tích

| Tiêu chí              | Golang (excelize)                        | Node.js (exceljs)                                  |
|-----------------------|-------------------------------------------|---------------------------------------------------|
| **Hiệu năng**         | Nhanh hơn ~3.7 lần                       | Chậm hơn, dễ nghẽn khi nhiều request              |
| **Thư viện**          | Native code, tối ưu I/O & bộ nhớ          | JavaScript thuần, overhead từ Garbage Collector   |
