package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"net/http"
	"log"
)

// Root endpoint
func Root(c *gin.Context) {
	res := `{"message":"root_handlers"}`
	c.JSON(http.StatusOK, res)
}

// SetKey endpoint
func SetKey(c *gin.Context) {

	key := c.Param("key")
	value := c.Param("value")

	log.Println("[ inserting ] key: " + key + ", value: " + value)
	cfg := &api.Config{Address:"localhost:8500"}
	csl, err := api.NewClient(cfg)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}

	// Get a handle to the KV API
	kv := csl.KV()
	// PUT a new KV pair
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err = kv.Put(p, nil)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}
	// Lookup the pair
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}

	log.Printf("KV: %v", pair)
	res := `{"message":"setkey_handler", "key":"added"}`

	c.JSON(http.StatusOK, res)
}

// GetKey endpoint
func GetKey(c *gin.Context){
	res := `{"message":"getkey_handlers"}`
    c.JSON(http.StatusOK, res)
}
