# gen_swagger

gen_swagger is a tool to covert swagger definition with code-message format response.

Assume that, we have a response object:

```json
{
    "name": "tom",
    "age": 28
}
```

then, gen_swagger could convert it into:

```json
{
    "code": 0,
    "msg":"ok",
    "data": {
        "name": "tom",
        "age": 28   
    }
}
```


## easy to use

```bash
gen_swagger -f example.swagger.json -o gen_example.swagger.json
```