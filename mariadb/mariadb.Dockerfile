# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM mariadb:10.3.9

#RUN /usr/bin/mysqld_safe
#RUN service mysql start
#RUN '/usr/bin/mysqladmin' -u root password '${MYSQL_PASSWORD:?err}'

EXPOSE 3306