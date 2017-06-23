FROM ubuntu:16.04

ADD chores /opt/chores

EXPOSE 8080

ENTRYPOINT ["/opt/chores"]


