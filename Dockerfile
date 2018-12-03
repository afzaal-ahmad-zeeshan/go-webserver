FROM golang:1.10-alpine3.7
COPY . . 

# I know, compilation is the best way to go, but still I am skipping it for now.
CMD [ "go", "run", "./src/program.go" ]