package mail

import (
	"bytes"
	"html/template"
	"path/filepath"
)

//Is the template that is going To be parsed on the html file
type Message struct {
	//It could be the name of the person
	Name string
	//One link To redirect if interested
	URL string
	//The content of the message
	Body string
}

//Receive the string name of the html file inside resources and the message To parse
func NewTemplateHandler(filename string, message *Message) (*bytes.Buffer,error) {
	templ := template.Must((template.ParseFiles(filepath.Join("resources", filename))))
	buf := new(bytes.Buffer)
	//join the html file with the message To be sended
	if err:= templ.Execute(buf,message); err!=nil{
		return nil,err
	}

	return buf,nil
}



