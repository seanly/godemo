FROM rockylinux:8

WORKDIR /data/app

COPY ./godemo godemo

ENTRYPOINT ["./godemo"]

