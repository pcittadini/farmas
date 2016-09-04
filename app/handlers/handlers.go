package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"net/http"
	"log"
	"os"
)

var consulHost = os.Getenv("COREOS_PRIVATE_IPV4")
// Root endpoint
func Root(c *gin.Context) {
	res := `{"message":"root_handler","status":"healthy"}`

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write([]byte(res))
}

// SetKey endpoint
func SetKey(c *gin.Context) {

	key := c.Param("key")
	value := c.Param("value")

	log.Println("[ inserting ] key: " + key + ", value: " + value)

	addr := consulHost+":8500"
	cfg := &api.Config{Address:addr}
	csl, err := api.NewClient(cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Get a handle to the KV API
	kv := csl.KV()
	// PUT a new KV pair
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err = kv.Put(p, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := `{"message":"setkey_handler", "key":"added"}`
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write([]byte(res))
}

// GetKey endpoint
func GetKey(c *gin.Context){
	//res := `{"message":"getkey_handlers"}`
	key := c.Param("key")
	addr := consulHost+":8500"
	cfg := &api.Config{Address:addr}
	csl, err := api.NewClient(cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Get a handle to the KV API
	kv := csl.KV()

	// Lookup the pair
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("[ get key ] KV: %v\n", pair)
    c.JSON(http.StatusOK, pair)
}
