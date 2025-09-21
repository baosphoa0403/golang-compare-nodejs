curl http://localhost:3000/hash-password

curl http://localhost:3000/hash-password-sync

curl http://localhost:8888/hash-password

# need generate mock xlsx
cd ./data && go run man.go 

docker-compose up --build -d    

# compare nodejs and golang
curl http://localhost:3000/excel-small  <-> curl http://localhost:8888/excel-small

curl http://localhost:3000/excel-medium <-> curl http://localhost:8888/excel-medium

curl http://localhost:3000/excel-large <-> curl http://localhost:8888/excel-large






