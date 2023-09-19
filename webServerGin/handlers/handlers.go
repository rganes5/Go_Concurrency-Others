package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rganes5/LCO/webServerGin/models"
)

func GetEndPoint(c *gin.Context) {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	//c.Query("url")
	res, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "failed to get the endpoint",
		})
		return
	}
	defer res.Body.Close()
	//var body map[string]interface{}
	var body models.Body

	//when using a unmarshall, we should use io.Readall and convert it as a byte slice then do the rest.
	/*
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "Failed to read ",
			})
		}
	*/
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		//if err := json.Unmarshal(bodyBytes, &body); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":  "failed to decode",
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": body,
	})

}

/*
json.NewDecoder:

It works with an io.Reader, which is useful for decoding JSON from an incoming stream, such as an HTTP response body.
It decodes JSON incrementally as it's read from the source.
It's often used when you need to decode JSON from a stream without loading the entire content into memory.
json.Unmarshal:

It works with a byte slice or string containing the entire JSON data.
It decodes the entire JSON data at once into a Go data structure.
It's typically used when you have the complete JSON data in memory, like when reading from a file or a pre-fetched HTTP response body.
In the provided code, the change from json.NewDecoder to json.Unmarshal involves reading the entire response body into a byte slice (bodyBytes)
before decoding it using json.Unmarshal. This is suitable when you want to decode the complete JSON response body into the body struct.
*/

func NumHandler(c *gin.Context) {
	var body models.Input
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to bind",
		})
		return
	}

	var out models.Output

	out.Numb = body.Num * body.Num

	c.JSON(http.StatusOK, gin.H{
		"success": out,
	})

}
