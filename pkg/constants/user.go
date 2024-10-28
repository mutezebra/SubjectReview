package constants

import (
	"time"
)

const (
	EmailCacheDDL = 3 * time.Minute
	MaxAvatarSize = 5 * MB
)

const JwtExpireTime = 48 * time.Hour

// OssAvatarFormat oss
const (
	OssAvatarFormat = "%d.%s" // uid.ext
)

// Mysql
const (
	TableNameOFUser = "user"
)

// Redis
const (
	UserIDKey                = "uid"
	VerifyCodePrefix         = "vc"
	PasswordVerifyCodePrefix = "pvc"
	VerifyCodeDDL            = 2 * time.Minute
)

const (
	VerifyCodeTemplate = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f9f9f9;
					color: #333;
					margin: 0;
					padding: 0;
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
				}
				.container {
					background-color: #fff;
					border-radius: 10px;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
					padding: 20px;
					max-width: 400px;
					text-align: center;
				}
				h1 {
					font-size: 28px;
					color: #444;
					margin-bottom: 20px;
				}
				p {
					font-size: 18px;
					color: #666;
					margin-bottom: 10px;
				}
				.code {
					font-size: 24px;
					font-weight: bold;
					color: #007BFF;
					border: 2px dashed #007BFF;
					padding: 10px;
					border-radius: 5px;
					display: inline-block;
					margin-top: 10px;
					background-color: #f0f8ff;
				}
			</style>
		</head>
		<body>
		
			<div class="container">
				<h1>Subject Review</h1>
				<p>This is your verify code</p>
				<div class="code">{{.Code}}</div>
			</div>
		
		</body>
		</html>
`
)
