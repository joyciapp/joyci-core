FROM golang:1.11

ARG APPLICATION_DIR="/app/joyci-core"
ARG APPLICATION_VERSION="v0.0.1"

LABEL version=${APPLICATION_DIR}
LABEL description="This is JoyCI Core's Dockerfile"

ADD . ${APPLICATION_DIR} 
WORKDIR ${APPLICATION_DIR}

RUN go build
RUN go install

CMD ["/bin/bash"]