<a name="v0.0.4"></a>
## [v0.0.4] - 2021-07-04
### Chore
- service and repository require context

### Feat
- git chglog
- **auth:** get user data from redis or jwt
- **auth:** save user login token to redis
- **cache:** add support for redis
- **util:** map to struct

### Fix
- **auth:** use main db to get user data when cache is not available

### Refactor
- **repository:** move contract to each repo
- **service:** move contract to each svc


<a name="v0.0.3"></a>
## [v0.0.3] - 2021-07-03
### Chore
- move database connection to platform dir
- **config:** move database config to global config
- **config:** move database connection to platform dir

### Feat
- **ioc:** include http delivery
- **util:** jwt token creation

### Fix
- **util:** defered file closing statement

### Refactor
- change db migrator to Soda CLI
- attach binary tools within project
- **delivery-http:** create run func for delivery struct

### TODO
- implement context from delivery, to services then repositories


<a name="v0.0.2"></a>
## [v0.0.2] - 2021-05-30
### Feat
- **auth:** save access token to DB
- **config:** create app config for app related configurations
- **docs:** display swagger as endpoint documentation
- **ioc:** configure services
- **ioc:** configure repository
- **repo:** add utility to get database from ioc
- **service:** add utility to get repo from ioc

### Fix
- util package name
- **ioc:** configure token and user repo

### Refactor
- organize all app files in internal/app dir
- volumes setting for docker compose template
- fix little warn
- add recover middleware
- move validator to config package
- change base package name
- **config:** single object for all config


<a name="v0.0.1"></a>
## v0.0.1 - 2021-02-10
### Docs
- **readme:** add architecture concept

### Feat
- **auth:** get user data from  JWT
- **cors:** add CORS with default configurations
- **jwt:** add custom JWT claims
- **logger:** add logger with JSON friendly outputs
- **token:** add token to db after jwt generated
- **utils:** add file utilities
- **utils:** add random utilities

### Fix
- change validator package from config to utils

### Refactor
- move validator from config to utils
- **auth:** add auth service interface
- **auth:** add JWT service interface
- **config:** add jwt config
- **database:** add entry point for database config
- **delivery-http:** separate routes to each controller
- **delivery-http:** move http router to http related modules
- **delivery-http:** change handlers name
- **repository:** add entry point for repository
- **repository:** change repositories name
- **service:** add entry point to access all services
- **service:** change services name
- **structure:** directory naming aim to clean arch
- **structure:** move all input type to DTO
- **structure:** ready to implement auth middleware
- **user:** add user service interface
- **user:** add user repository interface
- **util:** change package name from utils to util


[Unreleased]: https://github.com/hadihammurabi/belajar-go-rest-api/compare/v0.0.4...HEAD
[v0.0.4]: https://github.com/hadihammurabi/belajar-go-rest-api/compare/v0.0.3...v0.0.4
[v0.0.3]: https://github.com/hadihammurabi/belajar-go-rest-api/compare/v0.0.2...v0.0.3
[v0.0.2]: https://github.com/hadihammurabi/belajar-go-rest-api/compare/v0.0.1...v0.0.2
