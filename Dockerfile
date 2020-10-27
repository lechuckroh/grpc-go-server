FROM lechuckroh/alpine-curl:latest
LABEL MAINTAINER="Lechuck Roh <lechuckroh@gmail.com>"
RUN mkdir -p /app
COPY build/app /app/
COPY config*.yml /app/
COPY tools/wait /wait
WORKDIR /app
HEALTHCHECK CMD /app/healthcheck.sh
EXPOSE 9080 9090

CMD /wait && ./app
