# Go Sandbox
This is my sandbox to learn Go language

## Notes
1. `go mod init %MODULE_PATH%` module path will be the path where go will be able to fetch your module. Can be github link to your repo
2. Need to do `go mod tidy` to update module requirements
3. use `go mod edit -replace %ONLINE_MODULE_PATH=%LOCAL_MODULE_PATH%` to point the module to local module