all:
	protoc --go_out=. --ruby_out=. *.proto
	go get
