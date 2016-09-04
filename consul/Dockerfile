FROM progrium/consul:latest
MAINTAINER Democracy Works, Inc. <dev@democracy.works>

ENV GOMAXPROCS 10
ENV DOCKER_VERSION 1.9.1

ADD https://get.docker.io/builds/Linux/x86_64/docker-${DOCKER_VERSION} /bin/docker
RUN chmod +x /bin/docker

COPY ./etcd-bootstrap /bin/etcd-bootstrap

ENTRYPOINT ["/bin/etcd-bootstrap"]
CMD []
