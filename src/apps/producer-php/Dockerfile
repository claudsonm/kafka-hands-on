FROM alpine:3.14

ARG user=producer
ARG uid=1000
ENV user ${user}

RUN apk --no-cache add \
    php8 \
    php8-ctype \
    php8-curl \
    php8-dom \
    php8-fileinfo \
    php8-ftp \
    php8-iconv \
    php8-json \
    php8-mbstring \
    php8-mysqlnd \
    php8-openssl \
    php8-pdo \
    php8-pdo_sqlite \
    php8-pear \
    php8-phar \
    php8-posix \
    php8-session \
    php8-simplexml \
    php8-sqlite3 \
    php8-tokenizer \
    php8-xml \
    php8-xmlreader \
    php8-xmlwriter \
    php8-zlib \
    php8-gd \
    php8-pcntl \
    php8-pecl-swoole --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing/ \
    php8-pecl-rdkafka

COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

RUN ln -sf /usr/bin/php8 /usr/bin/php && \
    # Create system user to run Composer and Artisan Commands
    addgroup -g $uid -S $user && \
    adduser -G $user -u $uid -h /home/$user -S $user && \
    mkdir -p /home/$user/.composer && \
    chown -R $user:$user /home/$user

COPY .docker/php/entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

COPY . /var/www/html
RUN chown -R $user:$user /var/www/html
WORKDIR /var/www/html

USER $user
ENTRYPOINT ["entrypoint.sh"]
CMD ["php", "/var/www/html/artisan", "octane:start", "--host=0.0.0.0"]
