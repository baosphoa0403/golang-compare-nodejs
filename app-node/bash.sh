wrk -t4 -c100 -d30s http://localhost:3000/hash-password

wrk -t4 -c100 -d30s http://localhost:3000/hash-password-sync

wrk -t4 -c100 -d30s http://localhost:8888/hash-password

docker-compose up --build -d    