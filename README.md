# One Now

## Build Proto
```
mkdir -p frontend/src/gen backend/gen

buf lint proto
buf format proto -w
buf generate
```

## Note files
Note filenames have `<uuid>_<timestamp>.md` pattern.

A note can have multiple versions. `uuid` identifies a note. `timestamp` identifies a version of a note.


## Build Backend
```
cd backend

go run main.go -dir=../note -email=test@test.com
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