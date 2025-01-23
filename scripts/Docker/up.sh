start_db_container() {
    docker run -d \
        --name book_management_system_db \
        -e MYSQL_ROOT_PASSWORD=password \
        --network book_management_system \
        --memory=512m \
        --memory-swap=512m \
        -p 3306:3306 \
        -v db_data:/var/lib/mysql \
        mysql:8.0
}

start_backend_container() {
    docker run -d \
        --name book_management_system_backend \
        -e  DB="root:password@tcp(book_management_system_db:3306)/book_management_system?charset=utf8mb4&parseTime=True&loc=Local" \
        -e PORT="4000" \
        --network book_management_system \
        -p 4000:4000 \
        sayan4444/book-management-system:latest
}


if ! docker network ls | grep -q book_management_system; then
    docker network create book_management_system
fi

start_db_container

sleep 10

docker exec -i book_management_system_db mysql -u root -ppassword -e "CREATE DATABASE IF NOT EXISTS book_management_system;"

start_backend_container