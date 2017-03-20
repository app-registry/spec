## API Endpoint Discovery

### Purpose

API endpoints are often prefixed with long URIs that are not ergonomic for users specifying the URI via the command-line.
As a result, the specification includes an optional definition of `.well-known URI` ([RFC5785](https://tools.ietf.org/html/rfc5785)) in order to discover a full URI for the API endpoints.
When implemented a user can specify a host URI such as `example.com` and rely on the client to resolve that path to `example.com/registry/` rather than providing this full-length URI on the command-line.

### Definition

A `.well-known` endpoint is composed of two pieces: a path and the data located at that path.
The path for determining a full length URI is `/.well-known/app-registry-uri.json`.
The data located at this endpoint is well-formed JSON with the following [JSON Schema](http://json-schema.org/latest/json-schema-core.html):

```json
{
  "title": "App Registry URI Well Known Data",
  "type": "object",
  "properties": {
    "version": { "type": "string" },
    "uri_prefix": { "type": "string" }
	},
  "required": ["version", "uri_prefix"]
}
```

#### Example

```sh
$ curl https://example.com/.well-known/app-registry-uri.json
```

```json
{
  "version": "v1.0.0",
  "prefix": "https://example.com/registry/app-registry/"
}
```
