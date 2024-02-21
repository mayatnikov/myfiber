# rebuild and deploy new version    
include .env
IMAGE=$(REPO)/$(MYNAME):$(VERSION)
APP=$(MYNAME)
REPL=$(REPLICAS)

all: app docker_image  k8s_deploy
    
app: *.go .env views/*
	go test .
	go build -o app .

image: 
	docker build -t $(IMAGE) .
	docker push $(IMAGE)

deploy: 
	sed  -e s%REPLICAS_PLACEHOLDER%$(REPL)% \
	  -e s%APP_PLACEHOLDER%$(APP)% \
	  -e s%IMAGE_PLACEHOLDER%$(IMAGE)% \
	  myfiber2.yaml | kubectl apply -f -
	sed -e s%APP_PLACEHOLDER%$(APP)% myfiber-lb.yaml | kubectl apply -f - 
	echo "call: curl -v localhost:7777/json2"

delete:
	rm -f app
	kubectl delete deployment $(MYNAME)
	kubectl delete -f myfiber-lb.yaml
	docker compose down	
	docker rmi $(IMAGE)
