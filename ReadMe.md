# Tạo một sample database bằng Golang và go-pg

### Trước tiên, dùng Docker để chạy một PostgreSQL server, expose ra cổng 5432

```bash
docker run --name db -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres:11.1-alpine
```

### Có thể cài thêm pgAdmin để kiểm tra dữ liệu sau khi tạo

```shell
docker run -p 80:80 \
-e "PGADMIN_DEFAULT_EMAIL=email_cua_ban@gmail.com" \
-e "PGADMIN_DEFAULT_PASSWORD=password_cua_ban" \
-d dpage/pgadmin4
```

### Clone project về

```shell
git clone https://github.com/handuy/sample_db.git
```

### Chạy lệnh sau

```go
go run *.go
```