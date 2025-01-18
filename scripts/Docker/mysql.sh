docker run -d \
    --name db \
    -e MYSQL_ROOT_PASSWORD=password \
    --network book_management_system \
    --memory=512m \
    --memory-swap=512m \
    -p 3306:3306 \
    -v db_data:/var/lib/mysql \
    mysql:8.0