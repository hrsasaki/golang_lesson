# 1st command-line arg for success
# 2nd command-line arg for success without suffix "http://"
# 3rd command-line arg for failure with exit status 1
go run main.go http://gopl.io www.google.co.jp http://bad.gopl.io
