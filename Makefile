all:
	go build -o restapisample ./cmd/restapisample/...

clean:
	rm restapisample
