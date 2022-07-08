FROM golang:1.17.8-alpine3.15


RUN mkdir /app
ADD  . /app
WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

# Install Compile Daemon for go. We'll use it to watch changes in go files
RUN go get -u github.com/githubnemo/CompileDaemon

COPY *.go ./




COPY . .
COPY ./entrypoint.sh /entrypoint.sh

# wait-for-it requires bash, which alpine doesn't ship with by default. Use wait-for instead
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh
EXPOSE 8282

ENTRYPOINT [ "sh", "/entrypoint.sh" ]