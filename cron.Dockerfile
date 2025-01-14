FROM alpine:3.19

ARG USER=docker
ARG UID=1000
ARG GID=1000
ARG PW=docker

RUN addgroup -g ${GID} ${USER} \
    && adduser -D -u ${UID} -G ${USER} -h /home/${USER} -s /bin/ash ${USER} \
    && echo "${USER}:${PW}" | chpasswd \
    && mkdir /app \
    && chown -R ${USER}:${USER} /app

COPY cronApp /app

COPY .env .

CMD [ "/app/cronApp"]