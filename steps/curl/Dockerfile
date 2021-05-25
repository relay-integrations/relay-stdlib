FROM relaysh/core
COPY "./step.sh" "/nebula/step-api.sh"
ENTRYPOINT []
CMD ["/bin/bash", "/nebula/step-api.sh"]

LABEL "org.opencontainers.image.title"="apicall"
LABEL "org.opencontainers.image.description"="Makes an API call and returns the status code + response body"
LABEL "com.puppet.nebula.sdk.version"="v1"