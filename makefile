main: main.go
	go build main.go

.PHONY: clean
clean:
	rm *.o main