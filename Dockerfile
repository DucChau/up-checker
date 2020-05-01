FROM golang:1.14-alpine

# make working directory
RUN mkdir /code
WORKDIR /code

# add the source code
ADD main.go /code

# build it
RUN go build -o up-checker

# port
EXPOSE 8000

CMD ["./up-checker"]
