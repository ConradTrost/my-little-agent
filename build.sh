env GOOS=linux GOARCH=arm go build -v -ldflags "-s -w" -o my-little-agent/main cmd/my-little-agent/main.go


tar -czvf my_little_agent.tar.gz ./my-little-agent