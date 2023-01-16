package generateCode

import (
	"fmt"
	"github.com/google/uuid"
)

// GenerateCode 生成唯一code
func GenerateCode(category int) (resp string) {
	guid := uuid.New()
	guidStr := guid.String()
	switch category {
	case 1:
		resp = fmt.Sprintf("dept_code_%s", guidStr)
	}
	return resp
}
