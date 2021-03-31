FROM relaysh/core:latest
RUN apk --no-cache add bash ca-certificates jq
COPY step.sh /relay/step.sh
CMD ["/relay/step.sh"]

LABEL "org.opencontainers.image.title"="jq"
LABEL "org.opencontainers.image.description"="Run a jq filter on a given input"
LABEL "sh.relay.sdk.version"="v1"
