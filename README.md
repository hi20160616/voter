# voter
A simple vote web app.

# Easy start
```
sudo go run ./cmd/server/server.go
```

# API
Look forward: [Read me](api/README.md) located at *api/README.md*

# Docker and MySQL
Look forward: [Read me](internal/data/db/mysql/README.md) located at *internal/data/db/mysql/README.md*

# Language
Set language in configs.json at configs folder, can change the view of language.  
`"Language": "zh_CN"` in configs.json is default setting, you can change it yourself.  
For example, `en_US` while use template folder `templates/default/en_US`, similarly, change `Language` option to `zh_CN`, the web will use `templates/default/zh_CN`.

# Release
Set goreleaser config file and run `goreleaser release`  
If token not found error occur, view the page: https://goreleaser.com/errors/multiple-tokens/  
Also, we need generate token, at github: `Settings/Developer settings` -> `Personal access tokens (classic)`, generate a key and copy to the file:`~/.config/goreleaser/github_token`  
