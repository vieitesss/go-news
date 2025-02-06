set quiet

alias b := build
alias r := run

executable := "./go-news" 

_default:
  just -l

build:
  go build .

clean:
  rm {{executable}}

run:
  if [ {{ path_exists("{{executable}}") }} = "false" ]; then just build; fi
  "{{executable}}"
  just clean
