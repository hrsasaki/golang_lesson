PCGS=`go list -json $1 | go run main.go`
go list -json $PCGS
