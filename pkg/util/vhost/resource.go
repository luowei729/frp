// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/version"
)

var NotFoundPagePath = ""

const (
	NotFound = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>页面未找到 - 404</title>
    <style>
        /* 基础样式 */
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: #fff;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            overflow: hidden;
        }

        /* 容器样式 */
        .container {
            text-align: center;
            max-width: 600px;
            padding: 20px;
            background: rgba(255, 255, 255, 0.1);
            border-radius: 15px;
            backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
        }

        /* 标题样式 */
        h1 {
            font-size: 120px;
            margin: 0;
            font-weight: 700;
            color: #fff;
            text-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
        }

        /* 副标题样式 */
        h2 {
            font-size: 32px;
            margin: 10px 0;
            font-weight: 600;
        }

        /* 文字样式 */
        p {
            font-size: 18px;
            margin: 20px 0;
            line-height: 1.6;
        }

        /* 按钮样式 */
        .btn {
            display: inline-block;
            padding: 12px 30px;
            font-size: 16px;
            font-weight: 600;
            color: #667eea;
            background: #fff;
            border-radius: 50px;
            text-decoration: none;
            transition: all 0.3s ease;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
        }

        .btn:hover {
            transform: translateY(-3px);
            box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
        }

        /* 动画效果 */
        @keyframes float {
            0%, 100% {
                transform: translateY(0);
            }
            50% {
                transform: translateY(-20px);
            }
        }

        .floating {
            animation: float 4s ease-in-out infinite;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1 class="floating">404</h1>
        <h2>页面未找到</h2>
        <p>抱歉，您访问的页面似乎不存在。请检查网址是否正确，或返回首页。</p>
        <a href="https://www.lkz.pub" class="btn">返回首页</a>
    </div>
</body>
</html>
`
)

func getNotFoundPageContent() []byte {
	var (
		buf []byte
		err error
	)
	if NotFoundPagePath != "" {
		buf, err = os.ReadFile(NotFoundPagePath)
		if err != nil {
			log.Warnf("read custom 404 page error: %v", err)
			buf = []byte(NotFound)
		}
	} else {
		buf = []byte(NotFound)
	}
	return buf
}

func NotFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")

	content := getNotFoundPageContent()
	res := &http.Response{
		Status:        "Not Found",
		StatusCode:    404,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        header,
		Body:          io.NopCloser(bytes.NewReader(content)),
		ContentLength: int64(len(content)),
	}
	return res
}
