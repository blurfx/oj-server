# fxoj

## 로컬 개발 환경 설정

### Prerequisites

- Docker and docker-compose
- Golang 1.23+
- [goose](https://github.com/pressly/goose)

### Database

```sh
make docker-up
make dbinit
```

#### Database Migration

`make goose env=[ENV] c=[goose command]` 명령어로 사용합니다.

**Example:**

마이그레이션 적용
- `make goose env=local c=up`

마지막 마이그레이션 롤백
- `make goose env=local c=down`

### Server

[`air`](https://github.com/air-verse/air) 또는 `make run`으로 서버를 실행합니다.
