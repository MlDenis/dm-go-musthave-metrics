package templates

import "html/template"

var DataTemplate = template.Must(template.New("data").Parse(`
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
		</head>
		<body>
			<div>
				{{.Items}}
			</div>
		</body>
	</html>`))
