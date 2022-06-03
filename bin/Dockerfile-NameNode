FROM centos:centos8

WORKDIR /filestorage/

# RUN yum install golang

COPY ./namenode /filestorage/namenode

# CMD [  ]

ENTRYPOINT  ["./namenode"]

VOLUME ["/filestorage/.datas", "/filestorage/.conf"]
