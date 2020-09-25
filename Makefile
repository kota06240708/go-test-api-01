# コンテナを起動
.PHONY: start
start:
	docker-compose up -d 

# コンテナを起動 (コンソールにデバックを表示させる)
.PHONY: start-build
start-build:
	docker-compose up -d --build

# ログを表示
.PHONY: logs
logs:
	docker-compose logs

# goのログを表示
.PHONY: logs-go
logs-go:
	docker logs -f go-test-api-01_golang_1

# 開発終了
.PHONY: kill
kill:
	docker-compose kill

# DBデータをdump
.PHONY: dump
dump:
	docker exec -it mysql57 sh -c 'mysqldump wordpress -u wordpress -pwordpress 2> /dev/null' > db/mysql.dump.sql

# volume毎削除
.PHONY: down
down:
	docker-compose down --volumes

# コンテナの状態を表示
.PHONY: ps
ps:
	docker-compose ps

# 全てのコンテナの状態を表示
.PHONY: ps-all
ps-all:
	docker ps -a

# mysqlのコンテナの中に入る
.PHONY: on-db
on-db:
	docker exec -it mysql-container bin/bash

# nodeのコンテナの中に入る
.PHONY: on-go
on-go:
	docker exec -i -t go-test-api-01_golang_1 sh

# コンテナ、イメージを削除
.PHONY: clean
clean:
	@if [ "$(image)" != "" ] ; then \
			docker rmi -f $(image); \
	fi
	@if [ "$(contener)" != "" ] ; then \
			docker rm -f $(contener); \
	fi

# コンテナをリスタート
.PHONY: restart
restart: kill start
