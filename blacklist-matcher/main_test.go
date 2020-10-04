package main_test

import (
	"context"
	"testing"
	"fmt"
	"swear-scan/blacklist-matcher"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

type Request events.APIGatewayProxyRequest

var safeTextEnglish = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

var safeTextJaponese =  "定タ矢共のゃフレ議場づ事童ヨム幕速ぜ戦外ぎよて思直ソコオ近話マモリラ立検観20示ぱ労日5占丸オス定給オレ先忠ごむレび苦観づぐちど宇染命選どつだ。常経へ在技セツヘコ治持タエ学演ま視有トハ区付た多報こリにる案却3除よ高製渕一まる時俣妃やかッ。棋と駅否リコサ目領シカヒト面周けごぼ月8測ヘミユ融閉ヘ購高ヒネ丘郵ナテ者2目すたゆる真読さた護月せめ拡刈拒ふく"

var safeTextHebrew = "אקראי גיאוגרפיה ב בקר, בהשחתה עקרונות של בדף. בה יכול רוסית קרן. זכר חופשית לויקיפדיה מה. אנא ב מאמר למנוע שימושיים, מדע תורת הקנאים האטמוספירה דת. או ארץ היום בקלות העברית, או בשפה פילוסופיה מתן, כתב בה אחרים מוסיקה. אם כניסה בקרבת ותשובות חפש, שנתי אחרים מדריכים כתב של. תנך אירועים ביולוגיה או, היא מפתח ובמתן בחירות אם."

var unsafeTextEnglish = "A strip club is a venue where strippers provide adult entertainment, predominantly in the form of striptease or other erotic or exotic dances. Strip clubs typically adopt a nightclub or bar style, and can also adopt a theatre or cabaret-style."

func TestHandler(t *testing.T) {

	ctx := context.Background()
	tests := []struct {
		request main.Request
		expectBody  string
		expectCode  int
	}{
		{
			request: main.Request{Body: safeTextEnglish},
			expectBody:  "Submission is valid",
			expectCode:  200,
		},
		{
			request: main.Request{Body: safeTextJaponese},
			expectBody:  "Submission is valid",
			expectCode:  200,
		},
		{
			request: main.Request{Body: safeTextHebrew},
			expectBody:  "Submission is valid",
			expectCode:  200,
		},
		{
			request: main.Request{Body: unsafeTextEnglish},
			expectBody:  "Submission is invalid. Found black listed keyword: erotic",
			expectCode:  400,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(ctx, test.request)
		if err != nil {
			fmt.Sprintf("%T", response)
		}
		assert.Equal(t, test.expectCode, response.StatusCode)
		assert.Equal(t, test.expectBody, response.Body)
	}

}