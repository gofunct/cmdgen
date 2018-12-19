# Build Stage
FROM {{cmdgen.docker_build_image}}:{{cmdgen.docker_build_image_version}} AS build-stage

LABEL app="build-{{cmdgen.app_name}}"
LABEL REPO="https://github.com/{{cmdgen.github_username}}/{{cmdgen.app_name}}"

ENV PROJPATH=/go/src/github.com/{{cmdgen.github_username}}/{{cmdgen.app_name}}

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/{{cmdgen.github_username}}/{{cmdgen.app_name}}
WORKDIR /go/src/github.com/{{cmdgen.github_username}}/{{cmdgen.app_name}}

RUN make build-alpine

# Final Stage
FROM {{cmdgen.docker_image}}

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/{{cmdgen.github_username}}/{{cmdgen.app_name}}"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/{{cmdgen.app_name}}/bin

WORKDIR /opt/{{cmdgen.app_name}}/bin

COPY --from=build-stage /go/src/github.com/{{cmdgen.github_username}}/{{cmdgen.app_name}}/bin/{{cmdgen.app_name}} /opt/{{cmdgen.app_name}}/bin/
RUN chmod +x /opt/{{cmdgen.app_name}}/bin/{{cmdgen.app_name}}

# Create appuser
RUN adduser -D -g '' {{cmdgen.app_name}}
USER {{cmdgen.app_name}}

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/{{cmdgen.app_name}}/bin/{{cmdgen.app_name}}"]
