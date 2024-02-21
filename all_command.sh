go install .
go run .
docker compose build
docker compose push
kubectl apply myfiber-lb.yaml
kubectl apply myfiber.yaml

curl localhost:7777/nm/Петрович
curl localhost:7777/json2

curl localhost:7777/quit


