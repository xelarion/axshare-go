package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http/httputil"
)

// 恢复发生 panic 异常的程序
func RecoverProgram(errMsg ...interface{}) {
	err := recover()
	if err != nil {
		logrus.Error("error: " + fmt.Sprint(err) + " ; " + fmt.Sprint(errMsg...))
	}
}

func RecoveryLogToLogrus() gin.HandlerFunc {
	logger := logrus.StandardLogger()

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				logger.WithError(err).WithField("httpRequest", string(httpRequest)).Error("http error!")
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
