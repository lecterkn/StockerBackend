package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type JancodeUsecase struct{}

func NewJancodeUsecase() JancodeUsecase {
	return JancodeUsecase{}
}

const (
	JANCODE_LOOKUP_API_URL        = "https://api.jancodelookup.com/"
	JANCODE_LOOKUP_API_QUERY_TYPE = "code"
)

func (s JancodeUsecase) GetProductByCode(janCode string) (*JancodeUsecaseOutput, error) {
	// AppId取得
	appId := s.getAppId()

	// クエリパラメータを設定
	params := url.Values{}
	params.Add("appId", appId)
	params.Add("query", janCode)
	params.Add("type", JANCODE_LOOKUP_API_QUERY_TYPE)

	// url組み立て
	url := fmt.Sprintf("%s?%s", JANCODE_LOOKUP_API_URL, params.Encode())

	// リクエスト送信
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// ボディ読み取り
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// ボディjson化
	var jsonBody jancodeApiResponseJson
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		return nil, err
	}

	if jsonBody.Products == nil || len(jsonBody.Products) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return &JancodeUsecaseOutput{
		Name:      jsonBody.Products[0].ItemName,
		BrandName: jsonBody.Products[0].BrandName,
		MakerName: jsonBody.Products[0].MakerName,
	}, nil
}

func (JancodeUsecase) getAppId() string {
	appId, exists := os.LookupEnv("JANCODELOOKUP_APP_ID")
	if !exists {
		panic("\"JANCODELOOKUP_APP_ID\" is not set")
	}
	return appId
}

type jancodeApiResponseJsonProduct struct {
	ItemName  string `json:"itemName"`
	BrandName string `json:"brandName"`
	MakerName string `json:"makerName"`
}

type jancodeApiResponseJson struct {
	Products []jancodeApiResponseJsonProduct `json:"product"`
}

type JancodeUsecaseOutput struct {
	Name      string
	BrandName string
	MakerName string
}
