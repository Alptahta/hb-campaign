# hb-campaign

Simple module written in GO which reads given file and then manipultes products and campaigns in DB

To use DB with docker:
`make mysql`

To remove DB related images and containers:
`make purge`

Useful commands:
`docker exec -it mysqlCont bash`
`mysql -uroot -p`