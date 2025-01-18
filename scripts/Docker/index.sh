if ! docker network ls | grep -q book_management_system; then
    docker network create book_management_system
fi

scripts/Docker/mysql.sh

sleep 20

docker exec -i db mysql -u root -ppassword -e "CREATE DATABASE IF NOT EXISTS book_management_system;"

scripts/Docker/backend.sh