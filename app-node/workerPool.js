import { Worker } from "worker_threads";

class WorkerPool {
  constructor(workerFile, size) {
    this.size = size; // Số worker tối đa trong pool
    this.pool = []; // Mảng chứa các worker
    this.tasks = []; // Mảng các tác vụ chờ xử lý
    this.workerFile = workerFile;
  }

  // Khởi tạo worker pool
  initialize() {
    for (let i = 0; i < this.size; i++) {
      const worker = new Worker(this.workerFile);
      worker.on("message", (message) => {
        const task = this.tasks.shift();
        if (task) task.resolve(message); // Giải quyết tác vụ
        this._addWorkerToPool(worker); // Trả lại worker cho pool
      });
      worker.on("error", (err) => {
        console.error("Worker error:", err);
      });
      this.pool.push(worker);
    }
  }

  // Lấy worker từ pool
  _getWorker() {
    return new Promise((resolve) => {
      if (this.pool.length > 0) {
        resolve(this.pool.pop()); // Trả lại worker từ pool
      } else {
        // Nếu không còn worker, đợi một worker có sẵn
        this.tasks.push({ resolve });
      }
    });
  }

  // Trả lại worker vào pool
  _addWorkerToPool(worker) {
    this.pool.push(worker);
  }

  // Gửi công việc cho worker và nhận kết quả
  runTask(taskData) {
    return this._getWorker().then((worker) => {
      return new Promise((resolve, reject) => {
        worker.postMessage(taskData); // Gửi dữ liệu công việc cho worker
        this.tasks.push({ resolve, reject });
      });
    });
  }

  // Hủy tất cả worker trong pool
  close() {
    this.pool.forEach((worker) => worker.terminate());
  }
}

export default WorkerPool;
