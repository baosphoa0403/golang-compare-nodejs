# 📊 Benchmark Report: Excel Large (Node.js vs Go) (Streaming Read)

## 1. Giới thiệu

Case này kiểm tra hiệu năng đọc file Excel lớn (`large.xlsx`) bằng cơ chế **streaming** giữa:

- **Node.js** với thư viện [exceljs](https://www.npmjs.com/package/exceljs) (WorkbookReader)
- **Go** với thư viện [excelize](https://github.com/xuri/excelize) (Rows iterator)

Mục tiêu: đo **tốc độ đọc** và **số dòng xử lý** cho file có ~1,048,576 rows (dung lượng hàng trăm MB), đồng thời chứng minh ưu thế của streaming so với cách đọc toàn bộ vào RAM.

---

## 2. Cấu hình test

- **File**: `large.xlsx` (1,048,576 dòng).
- **Endpoint gọi**:
  - Node.js: `GET /excel-large`
  - Go: `GET /excel-large`
- **Môi trường**: Docker container, giới hạn `2 CPU`, `2GB RAM`.
- **Đo lường**: thời gian thực thi (ms).

---

## 3. Kết quả

- **Node.js**
  ```json
  { "case": "Large Excel Streaming", "rows": 1048576, "duration": "9679ms" }
  ```
- **Golang**

  ```json
  {
    "case": "Large Excel Streaming",
    "rows": 1048576,
    "duration": "3.935057251s"
  }
  ```

## 4. Phân tích

| Tiêu chí           | Golang (excelize)                 | Node.js (exceljs)                                                    |
| ------------------ | --------------------------------- | -------------------------------------------------------------------- |
| **Hiệu năng**      | Đọc xong 1M rows trong \~3.8 giây | Đọc xong 1M rows trong \~13.9 giây                                   |
| **Tốc độ**         | Nhanh hơn \~3.7 lần               | Chậm hơn đáng kể                                                     |
| **Quản lý bộ nhớ** | Iterator row-by-row, RAM thấp     | Streaming nhưng object allocation nhiều hơn, RAM cao                 |
| **Tính ổn định**   | Ổn định, phản hồi nhanh           | Ổn định nhưng latency cao, dễ bottleneck khi nhiều request song song |
