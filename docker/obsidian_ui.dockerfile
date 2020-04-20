FROM nginx
RUN apt-get update     \
 && apt-get install -y \
  apt-utils            \
  procps               \
  curl                 \
  lsof                 \
 && true
RUN mkdir -p /www
COPY ./static /www
RUN mkdir -p /www/js/ui
COPY ./ui /www/js/ui
RUN mkdir -p /etc/nginx.bak && mv /etc/nginx/* /etc/nginx.bak/
COPY ./nginx /etc/nginx/
RUN date > /build-timestamp
RUN ln -sf /dev/stdout /var/log/nginx/access.log && ln -sf /dev/stderr /var/log/nginx/error.log
