#!/bin/sh
migrator="soda"
config="$(pwd)/db/database.yml"
migrations="$(pwd)/db/migrations"

if ! type "$migrator" > /dev/null; then
  echo "installing $migrator ...";
  go get github.com/gobuffalo/pop/...;
  go install github.com/gobuffalo/pop/soda;
fi

if [ "$1" == "u" ] || [ "$1" == "up" ]; then
  soda m -c "$config" -p "$migrations";
  exit;
fi

if [ "$1" == "d" ] || [ "$1" == "down" ]; then
  soda m down -c "$config" -p "$migrations";
  exit;
fi

if [ "$1" == "r" ] || [ "$1" == "reset" ]; then
  soda reset -c "$config" -p "$migrations";
  exit;
fi

if [ "$1" == "g" ] || [ "$1" == "gen" ] || [ "$1" == "generate" ]; then
  if [ "$2" == "" ]; then
    echo "fail: missing migration name";
    exit 1;
  fi
  soda g sql -c "$config" -p "$migrations" "$2";
  exit;
fi

soda m -c "$config" -p "$migrations"
