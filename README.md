# One Now

## Build Proto
```
mkdir -p frontend/src/gen backend/gen

buf lint proto
buf format proto -w
buf generate
```

## Build Backend
```
cd backend
rsync -av ../note gen/

go run main.go
go build
```

## Build Frontend
```
cd frontend

yarn install
yarn test
yarn start
yarn build
```