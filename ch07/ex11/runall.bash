go run main.go &
echo 'http://localhost:8000/list'
../fetch http://localhost:8000/list
echo 'http://localhost:8000/update?item=socks&price=6'
../fetch http://localhost:8000/update?item=socks&price=6
echo 'http://localhost:8000/list'
../fetch http://localhost:8000/list
echo 'http://localhost:8000/update?item=shirt&price=30'
../fetch http://localhost:8000/update?item=shirt&price=30
echo 'http://localhost:8000/update?item=socks&price=a'
../fetch http://localhost:8000/update?item=socks&price=a
echo 'http://localhost:8000/list'
../fetch http://localhost:8000/list
