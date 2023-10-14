run:
	go build -o bin && ./bin
sqlc:
	sqlc generate
gooseup:
	cd sql/schema && goose postgres postgres://postgres:amin235711@amin-laptop.local:5432/shadowsocks up
goosedown:
	cd sql/schema && goose postgres postgres://postgres:amin235711@amin-laptop.local:5432/shadowsocks down