build:
	go build main.go
	docker build -t app .
	docker tag -f app snowpunk/app:$(v)
	docker push snowpunk/app:$(v)

redeploy:
	fleetctl stop app@$(node).service
	fleetctl destroy app@$(node).service
	fleetctl submit app@$(node).service
	fleetctl start app@$(node).service
