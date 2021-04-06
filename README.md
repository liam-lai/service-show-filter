# service-show-filter
API service for filtering series data

## API List
```
/        // filter has drm and has at least one episode 
/status  // return ok
```

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

## Heroku Setup
```
brew install heroku/brew/heroku
heroku login
heroku create
git push heroku master
```

## TODO
* ~~DEPLOY~~
* ~~GO MODULE~~
* ~~README~~