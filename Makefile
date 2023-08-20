format:
	go fmt ./... 

mocks:
	mockgen -source .\internal\app\repositories\product_repository.go  -destination .\internal\app\repositories\mock_product_repository.go -package repositories
	mockgen -source .\internal\app\repositories\campaign_repository.go  -destination .\internal\app\repositories\mock_campaign_repository.go -package repositories
	mockgen -source .\internal\app\repositories\order_repository.go  -destination .\internal\app\repositories\mock_order_repository.go -package repositories

	mockgen -source .\internal\faketime\time-simulator.go  -destination .\internal\faketime\mock_time-simulator.go -package faketime

	mockgen -source .\internal\app\services\product_service.go  -destination .\internal\app\services\mock_product_service.go -package services
	mockgen -source .\internal\app\services\campaign_service.go  -destination .\internal\app\services\mock_campaign_service.go -package services
	mockgen -source .\internal\app\services\order_service.go  -destination .\internal\app\services\mock_order_service.go -package services

unit-tests:
	go test -v ./...

unit-test-coverage:
	go test -v ./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html

mysql:
	docker build -t tempmysql:trial -f Dockerfile.mysql .
	docker run -d -p 3306:3306 --name mysqlCont -e MYSQL_ROOT_PASSWORD=12345 tempmysql:trial

	docker build -t hb-campaign:multistage -f Dockerfile.multistage .
	docker run -d --name hbCont hb-campaign:multistage

# To purge your computer from this project's related docker
purge:
	docker stop mysqlCont
	docker rm mysqlCont
	docker image rm tempmysql:trial
	docker stop hbCont
	docker rm hbCont
	docker image rm hb-campaign:trial

run-docker:
	docker compose up

