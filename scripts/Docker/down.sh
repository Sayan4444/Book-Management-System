docker kill book_management_system_db
docker kill book_management_system_backend

sleep 4

docker rm book_management_system_db
docker rm book_management_system_backend