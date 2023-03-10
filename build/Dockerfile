FROM golang:1.18 as build

# setup user
RUN useradd -ms /bin/bash developer
USER developer
WORKDIR /home/developer/

COPY --chown=developer:developer . .

# download dependencies and prepare fresh installation
RUN go mod download
RUN go build .
RUN chmod 777 /home/developer/k8s-watcher

FROM debian:buster

COPY --from=build /home/developer/k8s-watcher /k8s-watcher
RUN chmod +x /k8s-watcher
CMD ["/k8s-watcher"]