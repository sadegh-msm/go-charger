# Charge redeemer 

> Code generator and redeemer

## introduction
this project includes 2 services `wallet service` and `offering service`.

## wallet service
this service has all the information about account and phone number and their balance. I also used SQL database for `ACID` principles and also used ORM (GORM) for interacting with database. <br>

#### end points:
there are 4 `end points` that you can use to work with this service :

```go
func NewRouter(e *echo.Echo) *echo.Echo {
    e.POST("/newacc", handlers.NewAccount)
    e.POST("/charge", handlers.Charge)
    e.POST("/use", handlers.Use)
    e.GET("/balance", handlers.Balance)

    return e
}
````

## offering service
this service will handle all offer codes and the users that uses offer codes. I also used `redis` database for performance. <br>

#### end points:
there are 3 `end points` that you can use to work with this service :

```go
func NewRouter(e *echo.Echo) *echo.Echo {
    e.POST("/newcode", handlers.NewCode)
	// this API needs to call charge API from wallet service
    e.POST("/redeem", handlers.Redeem)
    e.GET("/codeusers", handlers.CodeUsers)

    return e
}
````

