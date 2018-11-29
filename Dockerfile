FROM go:1.11.2-alpine
COPY . . 
ENTRYPOINT [ "go", "./main/program.go" ]