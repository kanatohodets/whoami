FROM alpine
ADD whoami whoami-service
ENTRYPOINT [ "/whoami-service" ]
