FROM golang:alpine as scratch

WORKDIR /project

COPY . .
RUN go build -o project-clickhouse .

FROM scratch as prod
COPY --from=scratch /project/project-clickhouse /usr/local/bin/project-clickhouse

ARG command
ENV COMMAND ${command}
CMD project-clickhouse ${COMMAND}