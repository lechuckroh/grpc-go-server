#!/bin/sh

http_port=${API_HTTP_PORT:-9080}
curl --fail http://localhost:"${http_port}"/healthcheck || exit 1
