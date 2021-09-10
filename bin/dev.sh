#!/bin/sh

if ! type "fiber" > /dev/null; then
  echo "installing fiber ...";
  go get -u github.com/gofiber/cli/fiber;
fi

# if [ "$1" == "u" ] || [ "$1" == "up" ]; then
#   soda m -c "$config" -p "$migrations";
#   exit;
# fi

fiber dev
