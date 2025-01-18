docker run -d \
    --name backend \
    -e  DB="root:password@tcp(db:3306)/book_management_system?charset=utf8mb4&parseTime=True&loc=Local" \
    -e PORT="4000" \
    --network book_management_system \
    -p 4000:4000 \
    sayan4444/book-management-system:latest

    