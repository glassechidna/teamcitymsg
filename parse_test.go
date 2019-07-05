package teamcitymsg

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	input, err := os.Open("testdata/buildlog.log")
	assert.NoError(t, err)

	expected := []ServiceMessage{
		NewMsgInspectionType("gofmt", "gofmt", "gofmt style violation", "Code style"),
		NewMsgInspection("gofmt", "pkg/xhooks/handler.go", " 	*mux.Router\n }\n \n-func NewHandler    (api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {\n-	r    := mux.NewRouter()\n+func NewHandler(api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {\n+	r := mux.NewRouter()\n \tregister(api, r, sqser, snser, lambdaer, transport)\n 	return &Handler{r}\n }\n", 18),
		NewMsgInspection("gofmt", "pkg/xhooks/handler.go", " 	*mux.Router\n }\n \n-func NewHandler    (api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {\n-	r    := mux.NewRouter()\n+func NewHandler(api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {\n+	r := mux.NewRouter()\n \tregister(api, r, sqser, snser, lambdaer, transport)\n 	return &Handler{r}\n }\n", 19),
		NewMsgInspection("gofmt", "pkg/xhooks/sqs.go", " 	queueUrl string\n }\n \n-func     newSqsHandler(sqsApi sqsiface.SQSAPI, queueArn string) *sqsHandler {\n+func newSqsHandler(sqsApi sqsiface.SQSAPI, queueArn string) *sqsHandler {\n 	arnBits := strings.Split(queueArn, \":\")\n 	region := arnBits[3]\n 	accountId := arnBits[4]\n", 18),
	}

	assert.Equal(t, expected, Parse(input))
}
