package constants

// SubjectType
const (
	Go语言 = iota
	计算机网络
	算法
	前端
	Mysql
)

var TypeSlice = []int16{Go语言, 计算机网络, 算法, 前端, Mysql}

// Mysql
const (
	SubjectTable         = "subject"
	SubjectWithUserTable = "user_with_subject"
	RemindTable          = "remind"
)

// Redis
const (
	ManagerListKey = "ml"
	SubjectIDKey   = "sid"
)

const (
	SubjectNameLength    = 300
	RemindInterval       = 60 * 60  // 每到整点就提醒
	DefaultRemindEndTime = 3600 * 6 // 往后六小时内有等待记忆的都会提醒
	SendEmailRoutines    = 5
)

const (
	HTMLTemplate = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f4f4f9;
					color: #333;
					margin: 0;
					padding: 0;
				}
				.container {
					width: 80%;
					margin: 20px auto;
					background-color: #fff;
					border-radius: 8px;
					box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
					padding: 20px;
				}
				h1 {
					text-align: center;
					color: #444;
				}
				table {
					width: 100%;
					border-collapse: collapse;
					margin-top: 20px;
				}
				th, td {
					padding: 12px;
					text-align: left;
					border-bottom: 1px solid #ddd;
				}
				th {
					background-color: #f8f8f8;
					color: #555;
				}
				tr:hover {
					background-color: #f1f1f1;
				}
				.footer {
					text-align: center;
					margin-top: 20px;
					color: #888;
				}
			</style>
		</head>
		<body>
		
			<div class="container">
				<h1>Subject Records</h1>
				
				<table>
					<thead>
						<tr>
							<th>ID</th>
							<th>Name</th>
							<th>Answer</th>
							<th>Subject Type</th>
							<th>Phase</th>
							<th>Learn Times</th>
							<th>Last Review</th>
							<th>Next Review</th>
						</tr>
					</thead>
					<tbody>
						{{range .Records}}
						<tr>
							<td>{{.ID}}</td>
							<td>{{.Name}}</td>
							<td>{{.Answer}}</td>
							<td>{{.SubjectType}}</td>
							<td>{{.Phase}}</td>
							<td>{{.LearnTimes}}</td>
							<td>{{.LastReviewAt}}</td>
							<td>{{.NextReviewAt}}</td>
						</tr>
						{{end}}
					</tbody>
				</table>
			</div>
		
			<div class="footer">
				<p>This is an automatically generated email. Please do not reply.</p>
			</div>
		
		</body>
		</html>
`
)
