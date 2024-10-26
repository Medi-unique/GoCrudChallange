package main

import (
	"github.com/gin-gonic/gin"
	"example.com/person_crud/router"
)

func main() {
	r := gin.Default()
	r = router.Router(r)
	r.Run()
}
