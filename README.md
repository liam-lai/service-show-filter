# service-show-file
Some description

## Setup
```
cd $GOPATH
mkdir -p src/github.com/liam-lai/
cd src/github.com/liam-lai/
git clone https://github.com/liam-lai/service-show-filter.git
```
```
make dependency
make test
```

## Env Variable
```
export PORT=8080
export LOG_LEVEL=info
```

## Run
```
make dev
```

## Deploy
```
make release
make deploy
```

## TODO
* DEPLOY
* ~~GO MODULE~~
* README