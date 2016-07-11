# TWIT

TWIT is a simple CLI **T**ool for **W**r**I**ting **T**emplates. It is really
just a wrapper Go's text and html templates implementation which allows for
command line usage.

## Usage

First create a YAML encoded file to house your parameters:

```yaml
# bands.yml
title: Great Bands
bands:
  - name: The White Stripes
    members:
      - Jack White
      - Meg White
  - name: The Strokes
    mambers:
      - Julian Casablancas
      - Nick Valensi
      - Albert Hammond, Jr.
      - Nikolai Fraiture
      - Fabrizio Moretti
  - name: Yeah Yeah Yeahs
    members:
      - Karen O
      - Nick Zinner
      - Brian Chase
```

The create a Go formatted template called *bands.tpl*:

```go
<!doctype html>
<html>
	<head>
		<title>{{ .title }}</title>
	</head>
	<body>
		<h1>{{ .title }}</h1>
		{{ range $band := .bands }}
			<h2>{{$band.name}}</h2>
			<ul>
				{{ range $member := $band.members }}
					<li>{{ $member }}</li>
				{{ end }}
			</ul>
		{{ end }}
	</body>
</html>

```

Then run TWIT:

```
$ twit bands.tpl bands.html --params-file=bands.yml --html
```

Or you could pass in JSON instead:

```
$ twit bands.tpl bands.html --params='{"title", "Best Bands Ever"}'
```
