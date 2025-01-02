package service

import (
	"gin-blog/pkg/e"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetArticle(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	var code int

	a := ArticleService{}
	data := a.GetArticle(1)
	if data.ID == -1 {
		code = e.ERROR_NOT_EXIST_ARTICLE
	} else {
		code = e.SUCCESS
	}
	assert.Equal(t, http.StatusOK, code)
}

func TestMain(m *testing.M) {
	// 设置测试环境
	gin.SetMode(gin.TestMode)

	// 运行测试
	code := m.Run()
	os.Exit(code)
}

func TestFail(t *testing.T) {
	t.Fatal("not implemented")

}

// func TestGetArticles(t *testing.T) {
// 	// 制作表格测试集
// 	tests := []struct {
// 		name     string
// 		query    string
// 		wantCode int
// 		wantMsg  string
// 		wantData map[string]interface{}
// 	}{
// 		{
// 			name:     "Invalid state",
// 			query:    "state=2",
// 			wantCode: http.StatusOK,
// 			wantMsg:  e.GetMsg(e.INVALID_PARAMS),
// 			wantData: map[string]interface{}{},
// 		},
// 		{
// 			name:     "Invalid tag_id",
// 			query:    "tag_id=0",
// 			wantCode: http.StatusOK,
// 			wantMsg:  e.GetMsg(e.INVALID_PARAMS),
// 			wantData: map[string]interface{}{},
// 		},
// 		{
// 			name:     "No query parameters",
// 			query:    "",
// 			wantCode: http.StatusOK,
// 			wantMsg:  e.GetMsg(e.SUCCESS),
// 			wantData: map[string]interface{}{
// 				"list":  []models.Article{},
// 				"total": 0,
// 			},
// 		},
// 	}

// 	// 运行每一个子测试
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(w)
// 			c.Request, _ = http.NewRequest("GET", "/articles?"+tt.query, nil)

// 			a := ArticleService{}
// 			a.GetArticles(c)

// 			assert.Equal(t, tt.wantCode, w.Code)
// 			var response map[string]interface{}
// 			err := json.Unmarshal(w.Body.Bytes(), &response)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.wantMsg, response["msg"])
// 			assert.Equal(t, tt.wantData, response["data"])
// 		})
// 	}
// }
