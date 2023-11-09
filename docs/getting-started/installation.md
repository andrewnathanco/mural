# Installing Mural

Since mural is entirely server side rendered HTML you should be able to get mural running by configuring the server properly.

## Prerequisites

- Golang >v1.21.x
- Node >=v20.5.0

```bash
git clone https://github.com/andrewnathanco/mural.git
```

### TMDB

Before running any commands you first need to generate a TMDB API Key to pull down the movie information. Eventually I'd love to allow you to use whatever dataset you want, but for now it has to be TMDB and moies.

[Getting Started With TMDB](https://developer.themoviedb.org/docs/getting-started)

### Sessions

To keep track of unique user sessions I am using [Gorilla Sessions](https://github.com/gorilla/sessions) a helpful package to handle the work for me. In order to do that you need to generate a random session string. I used this site [Random Key Gen](https://randomkeygen.com/).

### Plausible

Since I am hosting the site for anyone to play I am tracking analytics. I use [Plausible](https://plausible.io/sites). Highly highly recommend for your projects. You won't need this as you'll primarily be hosting it locally.

### Worker

As the game is meant to refresh daily, I have set up a cron worker to reset user sessions at 11:59 EST. I would go in and change that to something more frequent for your use case. Modify `/code/worker/worker.go` to whatever time you'd want, I'd suggest every 5-10 minutes.

### Web Server

By default the application runs on port `:1323` at this point I have not plans to change that to be set by an env variable, but if enough people want that I would change it. If you want to modify that port change `/code/main.go`.

## Setup

```bash
cd ./code
mv .env.template .env
```

Now you'll want to modify the `.env` to turn off analytics, include your TMDB API Key and to set your Session key.

At this point you should be all set to run.

```bash
go run main.go
```

If you are developing you'll need to setup tailwind and golang [air](https://github.com/cosmtrek/air).

In one terminal run:

```bash
air
```

In another:

```bash
npx tailwindcss -i ./view/input.css -o ./static/styles.css --watch
```
