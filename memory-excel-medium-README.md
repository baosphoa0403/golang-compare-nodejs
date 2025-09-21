# 📊 Benchmark Report: Excel Medium (Node.js vs Go) (In-Memory Read)

## 1. Giới thiệu

Case này kiểm tra hiệu năng đọc file Excel vừa (`medium.xlsx`) giữa:

- **Node.js** với thư viện [exceljs](https://www.npmjs.com/package/exceljs)
- **Go** với thư viện [excelize](https://github.com/xuri/excelize)

Mục tiêu: đo **tốc độ đọc** và **số dòng xử lý** cho file có ~1 triệu rows (dung lượng ~200–300MB).

---

## 2. Cấu hình test

- **File**: `medium.xlsx` (1,000,001 dòng).
- **Endpoint gọi**:
  - Node.js: `GET /excel-medium`
  - Go: `GET /excel-medium`
- **Môi trường**: Docker container, giới hạn `2 CPU`, `2GB RAM`.
- **Đo lường**: thời gian thực thi (ms).

---

## 3. Kết quả

- **Node.js**
  ⛔ App crash, không trả response (OOM khi load workbook).

- **Golang**

```json
{ "case": "Medium Excel", "rows": 1000001, "duration": "9.257096546s" }
```
