# gobooru [![Build Status](https://travis-ci.org/ayame-git/gobooru.svg?branch=master)](https://travis-ci.org/ayame-git/gobooru)
API wrapper for boorus written in go

## Examples
### Query
```go
posts, _ := api.GetPosts(nil)
```
You can also pass optional parameters as `url.Value`
```go
v := url.Value{}
v.Set("limit", "10")
posts, _ := api.GetPosts(v)
```

## License
gobooru is free software licenced under the Apache-2.0 license. Details provided in the LICENSE file
