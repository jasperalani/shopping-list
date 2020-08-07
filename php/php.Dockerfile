FROM php:7.2-fpm

WORKDIR var/www/html

COPY . .

RUN apt-get update && \
    apt-get install -y git zip

RUN curl --silent --show-error https://getcomposer.org/installer | php && \
    mv composer.phar /usr/local/bin/composer

RUN composer require slim/slim:"4.*"

RUN composer require slim/psr7

EXPOSE 9000