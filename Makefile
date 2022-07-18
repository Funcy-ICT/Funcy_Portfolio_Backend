DOCKER=docker-compose

run:
	nohup ./file-server/file &
	$(DOCKER) up --build
down:
	$(DOCKER) down
	rm nohup.out