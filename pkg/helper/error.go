package helper

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/L4B0MB4/MSNGR/pkg/models/api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// helper for serializing validation errors (and similar) properly
// everything but validation and json-unmarshall errors are considered unkown
func AbortWithBadRequest(ctx *gin.Context, ginBindingError error) {

	errs := validator.ValidationErrors{}
	ok := errors.As(ginBindingError, &errs)

	if ok {

		responseErrors := make([]api.Error, len(errs))
		for i := 0; i < len(errs); i++ {
			responseErrors[i] = api.Error{
				Id:     GetTraceId(ctx),
				Detail: fmt.Sprintf("Validation failed for field '%s'", errs[i].Field()),
			}
		}
		resp := api.Response{
			Errors: responseErrors,
		}
		ctx.AbortWithStatusJSON(400, resp)
		return

	}
	unmarshalErr := &json.UnmarshalTypeError{}
	ok = errors.As(ginBindingError, &unmarshalErr)
	if ok {

		resp := api.Response{
			Errors: []api.Error{{
				Id:     GetTraceId(ctx),
				Detail: fmt.Sprintf("Validation failed for field '%s'", unmarshalErr.Field),
			},
			},
		}
		ctx.AbortWithStatusJSON(400, resp)
		return
	}

	resp := api.Response{
		Errors: []api.Error{{
			Id:     GetTraceId(ctx),
			Detail: "Unkown error occured",
		},
		},
	}
	ctx.AbortWithStatusJSON(500, resp)

}

// aborts all but the current handler and returns statuscode 200
func AbortWithOk(ctx *gin.Context, err error) {

	resp := api.Response{
		Data: api.Detail{
			Description: err.Error(),
		},
		Errors: []api.Error{},
	}
	ctx.AbortWithStatusJSON(200, resp)
}

// aborts all but the current handler and returns statuscode 500 with a specific error message from the custom error
func AbortWithCustomError(ctx *gin.Context, err error) {

	resp := api.Response{
		Errors: []api.Error{{
			Id:     GetTraceId(ctx),
			Detail: err.Error(),
		},
		},
	}
	ctx.AbortWithStatusJSON(500, resp)
}

// aborts all but the current handler and returns statuscode 500 with "unkown error"-message
func AbortWithUnkownError(ctx *gin.Context, err error) {

	log.Error().Ctx(ctx).Err(err).Msg("Unkown error occured. Aborting request")
	resp := api.Response{
		Errors: []api.Error{{
			Id:     GetTraceId(ctx),
			Detail: "Unkown error occured",
		},
		},
	}
	ctx.AbortWithStatusJSON(500, resp)
}
