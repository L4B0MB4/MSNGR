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
		ctx.AbortWithStatusJSON(404, resp)
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

func AbortWithCustomError(ctx *gin.Context, err error) {

	resp := api.Response{
		Errors: []api.Error{{
			Id:     GetTraceId(ctx),
			Detail: "Unkown error occured",
		},
		},
	}
	ctx.AbortWithStatusJSON(500, resp)
}

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
