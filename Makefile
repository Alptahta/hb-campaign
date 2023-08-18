mysql:
	docker build -t tempmysql:trial -f Dockerfile.mysql .
	docker run -d -p 3306:3306 --name mysqlCont -e MYSQL_ROOT_PASSWORD=12345 tempmysql:trial

# To purge your computer from this project's related docker
purge:
	docker stop mysqlCont
	docker rm mysqlCont
	docker image rm tempmysql:trial