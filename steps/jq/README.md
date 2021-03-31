# stdlib-step-jq

This step applies a jq filter to some arbitrary input data.

## Specification

This step expects the following fields in the `spec` section of a workflow step definition that uses it:

| Setting  | Data type | Description                                                                     | Default | Required |
|----------|-----------|---------------------------------------------------------------------------------|---------|----------|
| `input`  | Any       | The input data to provide to jq                                                 | None    | Yes      |
| `raw`    | Boolean   | Whether to output the result as a raw string (-r flag)                          | false   | No       |
| `args`   | Object    | Arguments to pass to the jq filter as predefined variables, eg `something: bar` | None    | No       |
| `filter` | String    | The jq filter to run, eg `'map(select(.foo != $something ))'`                   | None    | Yes      |

## Usage

```yaml
step:
  name: grab-something
  image: relaysh/stdlib-step-jq
  spec:
    input:
    - foo: "bar"
    - foo: "baz"
    - foo: "quux"
    args:
      something: "bar"
    filter: "map(select(.foo != $something))"
```
