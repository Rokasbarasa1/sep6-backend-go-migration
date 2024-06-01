package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Single function to validate integer parameters
func ValidateInt(param, rules string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, _ := strconv.Atoi(c.Param(param))
		fmt.Println(fmt.Sprintf("Parameter: '%s' rule '%s' value %d", param, rules, value))

		if err := validate.Var(value, rules); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Validation failed for %s: %s", param, err.Error())})
			c.Abort()
			return
		}
		c.Set(param, value) // Optionally set the converted value in the context
		c.Next()
	}
}

func ValidateString(param string, rule string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := validate.Var(c.Param(param), rule)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": param + " " + err.Error()})
			c.Abort()
			return
		}
	}
}

// Validation middleware function
func ValidateParams(rules map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for param, rule := range rules {
			value := c.Param(param)
			fmt.Println(fmt.Sprintf("Parameter: %s rule %s value %d", param, rule, value))
			if value == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": param + " is required"})
				c.Abort()
				return
			}

			// Apply validation using the validator package
			err := validate.Var(value, rule)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": param + " " + err.Error()})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
