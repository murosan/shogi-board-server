FROM envoyproxy/envoy:latest

ENV TZ=Asia/Tokyo

RUN apt-get install -y logrotate
COPY ./envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml
EXPOSE 8080
