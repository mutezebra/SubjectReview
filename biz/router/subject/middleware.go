// Code generated by hertz generator.

package subject

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/mutezebra/subject-review/pkg/middleware"
)

func rootMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Cors()}
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _subjectMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.JWT()}
}

func _addforgetedsubjectMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _addsuccesssubjectMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getsubjectsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getanswersubjectrecordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getneededreviewsubjectsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _addnewsubjectMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _addforgetsubjectMw() []app.HandlerFunc {
	// your code...
	return nil
}
