wrk -t4 -c100 -d30s http://localhost:3000/hash-password

wrk -t4 -c100 -d30s http://localhost:3000/hash-password-sync

wrk -t4 -c100 -d30s http://localhost:8888/hash-password

docker-compose up --build -d    

wrk -t4 -c100 -d30s http://localhost:3000/excel-small

wrk -t4 -c100 -d30s http://localhost:3000/excel-medium

wrk -t4 -c100 -d30s http://localhost:3000/excel-large


wrk -t4 -c100 -d30s http://localhost:8888/excel-small

wrk -t4 -c100 -d30s http://localhost:8888/excel-medium

wrk -t4 -c100 -d30s http://localhost:8888/excel-large