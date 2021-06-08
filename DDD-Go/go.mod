module DDD-Go

go 1.16

require (
	github.com/go-kit/kit v0.10.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/gorilla/mux v1.8.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace DDD-Go/application => ../application

replace DDD-Go/infrastructure => ../infrastructure

replace DDD-Go/domain/model => ../domain/model

replace DDD-Go/domain/service => ../domain/service
