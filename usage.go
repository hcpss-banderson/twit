package main

import (
	"fmt"
	"os"
)

const (
	// BANNER lets people know what this is.
	BANNER = `  _________          _______ _______
 |__   __\ \        / /_   _|__   __|
    | |   \ \  /\  / /  | |    | |
    | |    \ \/  \/ /   | |    | |
    | |     \  /\  /   _| |_   | |
    |_|      \/  \/   |_____|  |_|

A tool for simple templating.

Version %s
Copyright (C) Brendan Anderson
Distributed under the terms of the MIT license

`
	// USAGE show people how to use the program.
	USAGE = `
Usage:
  twit <template> <destination> [options]

Arguments:
  template                The path to a Golang formatted template.
  destination             The path to template output. Existing files will be
                          overwritten.

Options:
  -p, --params=PARAMS     Params can either be the path to a YAML formatted file
                          with template parameters or a JSON encoded string of
                          template parameters.
  -n, --no-escape         Disable automatic output escaping.
  -h, --help              Display this help information.
  -v, --version           Display version information.

Examples:
  twit settings.php.tpl ../default/setting.php --params='{"dbname": "drupal"}'
  twit default.conf.tpl /etc/apache2/sites-available/default.conf -p apache.yml
  twit template.tpl page.html -p param1.yml -p param2.yml -p '{"key": "value"}'
`
)

func twitUsage() {
	fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, VERSION))
	fmt.Fprint(os.Stderr, fmt.Sprint(USAGE))
}
