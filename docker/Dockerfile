FROM golang:1.18

ARG docker_id
ARG docker_user

RUN useradd -G sudo,root -u ${docker_id} -d /home/${docker_user} ${docker_user}
RUN mkdir /home/${docker_user}
RUN chown ${docker_user}:${docker_user} /home/${docker_user}

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN apt-get update
RUN apt-get install net-tools lsof

WORKDIR /usr/src/app

CMD ["air"]