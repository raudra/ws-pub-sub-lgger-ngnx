FROM nginx:latest
COPY default /etc/nginx/conf.d/default.conf
COPY index.html /usr/share/nginx/html/index.html

EXPOSE 80/tcp

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
