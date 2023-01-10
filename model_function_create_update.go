/*
 * FunLess Platfom API
 *
 * The API for the FunLess Platform
 *
 * API version: 0.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"os"
)

type FunctionCreateUpdate struct {
	// Name of the function
	Name string `json:"name,omitempty"`
	// File with the code of the function
	Code **os.File `json:"code,omitempty"`
	// Events that can trigger the function
	Events []V1fnmoduleNameEvents `json:"events,omitempty"`
}
