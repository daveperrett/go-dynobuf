all:
	protoc --go_out=. --java_out=. *.proto
	go get
