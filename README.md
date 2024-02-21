### все команды
- локальная проверка
go run . 

- проверка откомпилированного кода
go install
~/go/bin/myfiber ./views
curl -X GET  http://172.20.0.50:3000/json2
curl -X GET  http://localhost:3000/nm/Vitaly

- проверка в docker container
docker compose build
docker compose up

d push localhost:5001/myfiber

- проверка в kubernates
kubectl apply -f myfiber.yaml
kubectl apply -f myfiber-lb.yaml
kubectl get pods,svc

curl -X GET  http://172.20.0.50:7777/json2
curl -X GET  http://localhost:7777/nm/Vitaly




