# Benchmark Import Excel: Node.js vs Go

Dự án này dùng để so sánh hiệu năng giữa **Node.js (exceljs)** và **Go (excelize)** trong việc import file Excel với nhiều kích thước khác nhau.  

## 🎯 Mục tiêu
- Đo **tốc độ xử lý** (parse / insert DB).  
- Ghi nhận **sử dụng bộ nhớ (RAM)** và **CPU**.  
- Kiểm tra khả năng **xử lý đồng thời (concurrency)** khi vừa import vừa phục vụ API.  
- Đưa ra cái nhìn thực tế khi chọn công nghệ cho hệ thống cần import dữ liệu lớn.

---

## 🚀 Các kịch bản test

### 1. File nhỏ (baseline)
- **Kích thước**: ~10 MB (~50k dòng).  
- **Hành động**: load toàn bộ file Excel vào RAM và parse.  
- **Mục tiêu**: đo tốc độ cơ bản, tính đơn giản code.  
- **Kỳ vọng**: Node.js và Go khá tương đồng, Go dùng ít RAM hơn.

---

### 2. File vừa (stress bộ nhớ)
- **Kích thước**: 100 MB – 500 MB (~1 triệu dòng).  
- **Hành động**: load toàn bộ sheet vào RAM.  
- **Mục tiêu**: đo peak memory, xem GC ảnh hưởng thế nào.  
- **Kỳ vọng**: Node.js dễ bị spike RAM, Go ổn định hơn.

---

### 3. File lớn (streaming)
- **Kích thước**: 1 GB – 5 GB (hàng triệu dòng).  
- **Hành động**: đọc streaming row-by-row thay vì load toàn bộ.  
- **Mục tiêu**: test khả năng giữ RAM ổn định với file cực lớn.  
- **Kỳ vọng**: cả hai xử lý được nhờ streaming, nhưng Go tận dụng được nhiều core nên nhanh hơn.

---

### 4. Validate + Insert DB
- **File**: ~100 MB.  
- **Hành động**: stream row → validate (VD: check email hợp lệ) → batch insert PostgreSQL.  
- **Mục tiêu**: đo throughput (rows/sec), p95 latency khi insert.  
- **Kỳ vọng**: Go nhanh hơn và dùng CPU hiệu quả hơn nhờ goroutine.

---

### 5. Song song Import + API traffic
- **Setup**: chạy HTTP API server đồng thời import file lớn.  
- **Hành động**: gửi 100 request/giây trong lúc import.  
- **Mục tiêu**: đo API latency khi có import task nặng chạy song song.  
- **Kỳ vọng**: Node.js event loop bị block → latency tăng; Go vẫn ổn định nhờ goroutines tách biệt.

---

## ⚙️ Cấu hình môi trường

- **Docker** + **Docker Compose**  
- Node.js 20 (trong Docker)  
- Go 1.23+ (trong Docker)  
- PostgreSQL 15 (cho test insert DB)  

### Giới hạn tài nguyên
- **Baseline**: `--cpus=1`, `--memory=2g` → so sánh công bằng Node vs Go (1 core).  
- **Stress test**: `--cpus=4`, `--memory=8g` → kiểm tra concurrency khi Go tận dụng đa nhân.

---

## 📊 Chỉ số cần theo dõi
- ⏱️ **Thời gian chạy** (ms / giây).  
- 💾 **Peak memory** (MB).  
- ⚡ **CPU usage** (%).  
- 📈 **Throughput** (rows/giây).  
- 🌐 **API latency** (p95/p99 khi có traffic song song).  

---

## ✅ Kết quả mong đợi
- **File nhỏ**: Node.js ~ Go, nhưng Go ít RAM hơn.  
- **File vừa/lớn**: Node.js dễ choke vì GC/memory, Go ổn định.  
- **Streaming + DB**: Go nhanh & tận dụng CPU tốt hơn.  
- **Concurrent traffic**: Node.js dễ bị block loop, Go vẫn phục vụ request đều.  
