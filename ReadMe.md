# Tạo một sample database bằng Golang và go-pg

## Database này gồm 7 bảng

![Database design](final_db.png?raw=true "Database design")

- club: 20.000 bản ghi
- league: 1000 bản ghi
- club_league: 1.000.000 bản ghi
- nation: 247 bản ghi
- cup: 1000 bản ghi
- nation_cup: 150.000 bản ghi
- player: 1.000.000 bản ghi

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

### Màn hình terminal hiển thị như ảnh dưới thì là OK

![Final result](result.png?raw=true "Final result")