# TWIT

TWIT is a simple CLI **T**ool for **W**r**I**ting **T**emplates. It is really
just a wrapper Go's text and html templates implementation which allows for
command line usage.

## Usage

### Templates

Templates are written in the standard
[Go template](https://golang.org/pkg/text/template) format.

### Filters

Filter support comes from the [leekchan/gtf](https://github.com/leekchan/gtf)
package.

### Parameters

TWIT accepts parameters for the template as either a YAML file or as a JSON
string.

### Example

For our example, let's image that we needed to write an Apache vhosts
configuration file.

First we create a template. We'll call it *vhosts.conf.tpl*:

```go
Listen {{ .port }}

ServerName {{ .server_name }}
DocumentRoot "{{ .document_root }}"

{{ range $host := .hosts }}
    <VirtualHost {{ $host.ip }}>
        DocumentRoot "{{ $host.document_root }}"
        ServerName {{ $host.server_name }}
    </VirtualHost>
{{ end }}
```

Then we will create a YAML encoded file to hold the parameters. We'll call it
*vhosts.yml*:

```yaml
port:          80
server_name:   server.company.com
document_root: /srv/main
hosts:
  - ip:            172.20.30.1
    document_root: /srv/website1
    server_name:   site1.company.com
  - ip:            172.20.30.2
    document_root: /srv/website2
    server_name:   site2.company.com
```

Generate the Apache configuration with TWIT:

```
$ twit vhosts.conf.tpl /etc/apache2/sites-available/vhosts.conf -p vhosts.yml
```
