package constants

import (
	"fmt"
	"os"
)

var BaseUrl = fmt.Sprintf("https://%s", host)

var host = os.Getenv("HOST")

var UrlNotFoundError = map[string]string{"message": "invalid url."}
var InternalServerError = map[string]string{"message": "server error. Please try again later."}
var InvalidRequestBodyError = map[string]string{"message": "invalid request body. Expected argument 'original_url' missing or invalid."}
var InvalidUrlFormatError = map[string]string{"message": "invalid url format. Expected absolute url."}

var ConnectionString = ""

const Base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
