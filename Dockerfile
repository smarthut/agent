FROM scratch
COPY agent /
EXPOSE 8080
ENTRYPOINT ["/agent"]
