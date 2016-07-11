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
  -f, --params-file=FILE  The path to a YAML formatted file containing the
                          template parameters.
  -p, --params=PARAMS     A JSON formatted string containing the template
                          parameters.
      --html              Template in HTML mode. This does some output escaping.
  -h, --help              Display this help information.
  -v, --version           Display version information.

Examples:
  twit settings.php.tpl ../default/setting.php --params='{"dbname": "drupal"}'
  twit default.conf.tpl /etc/apache2/sites-available/default.conf -f=apache.yml
`
)

func twitUsage() {
	fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, VERSION))
	fmt.Fprint(os.Stderr, fmt.Sprint(USAGE))
}
