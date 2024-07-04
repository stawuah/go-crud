DB_URL="postgresql://stephenawuah@localhost:5432/vim?sslmode=disable"

.PHONY: up down

up:
	migrate -database "postgresql://stephenawuah@localhost:5432/vim?sslmode=disable" -path db/ up 

down:
	migrate -database "postgresql://stephenawuah@localhost:5432/vim?sslmode=disable" -path db/ down

status:
	$(MIGRATE_CMD) version

create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir db -seq $$name

run:
	go run cmd/api/main.go

fixforce:
	migrate -database "postgresql://stephenawuah@localhost:5432/vim?sslmode=disable" -path db/ force 3