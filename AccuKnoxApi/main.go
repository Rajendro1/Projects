package main

import (
	pgdatabase "github.com/Rajendro1/Projects/AccuKnoxApi/pgDatabase"
	
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	pgdatabase.Connect(&gin.Context{})
	// routers.HandleRequest()
}
