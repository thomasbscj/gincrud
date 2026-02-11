package main

import (
	"errors"
	"fmt"
	"os"
)

func main(){
	argc := len(os.Args)
	if argc < 3 {
		printDefaultMessage()
		return
	}
	argv := os.Args

	command := argv[1]
	resourceName := argv[2]
	outdir := argv[3]
	var err error
	switch(command){
		case "help": printHelpMessage()
		case "create" : err = createResource(resourceName, outdir)
		case "clean" : {
			reTy := argv[4]
			err = createClean(reTy, resourceName, outdir)
			}
		default : printDefaultMessage()
	}
	if err != nil {
		fmt.Print(err)
	}
}

func printHelpMessage(){
	fmt.Printf(
`To use gincrud, write the command like this:

gincrud [command] [args]

these are the commands:
create - creates a crud with the resource you named : [resource name] [output directory]
clean - [resource name] [output diretory] [[model, controller, resource, service]] - creates a resource based on clean design
`)
}

func createClean(resourceType string, resourceName string, outdir string)  error{
	switch(resourceType){
		case "controller" : return createControllerFile(resourceName, outdir)
		case "service" : return createServiceFile(resourceName, outdir)
		case "repository" : return createRepositoryFile(resourceName, outdir)
		case "resource" : return createRepositoryFile(resourceName, outdir)
		case "model" : return createModelFile(resourceName, outdir)
		default: return errors.New("Not known resource")
	}
}

func createResource(resourceName string, outdir string)error{
	err := os.Mkdir(outdir + resourceName, 0755)
	if err != nil {
		panic(1)
	}
	outdir = outdir + resourceName + "/"
	err = createControllerFile(resourceName, outdir)
	if err != nil{
		return err
	}

	err = createModelFile(resourceName, outdir)
	if err != nil{
		return err
	}
	err = createRepositoryFile(resourceName, outdir)
	if err != nil{
		return err
	}
	err = createServiceFile(resourceName, outdir)

	return err
}

func printDefaultMessage(){
	fmt.Printf("Not valid command.\n")
	printHelpMessage()
}

func createRepositoryFile(resourceName string, outdir string) error{
	repoFile, err := os.Create(outdir + resourceName + ".repository.go")
	if err != nil {
		return err
	}
	_, err = repoFile.Write([]byte(
		getRepoFileContent(
			returnLow(resourceName),
			returnCaptalized(resourceName))))
	return err
}

func createServiceFile(resourceName string, outdir string) error{
	serviceFile, err := os.Create(outdir + resourceName + ".service.go")
	if err != nil {
		return err
	}
	_, err = serviceFile.Write([]byte(
		getServiceFileContent(
			returnLow(resourceName),
			returnCaptalized(resourceName))))
	return err
}

func createModelFile(resourceName string, outdir string) error{
	modelFile, err := os.Create(outdir + resourceName + ".model.go")
	if err != nil {
		return err
	}
	_, err = modelFile.Write([]byte(
		getModelFileContent(
			returnLow(resourceName),
			returnCaptalized(resourceName))))
	return err
}

func createControllerFile(resourceName string, outdir string) error{
	controllerFile, err := os.Create(outdir + resourceName + ".controller.go")
	if err != nil {
		return err
	}
	_, err = controllerFile.Write([]byte(
		getControllerFileContent(
			returnLow(resourceName),
			returnCaptalized(resourceName))))
	return err
}

func getModelFileContent(resourceNameLow string, resourceNameHigh string) string{
	return fmt.Sprintf(
`package %s

type %sModel struct {

}	
`,resourceNameLow, resourceNameHigh)
}

func getRepoFileContent(resourceNameLow string, resourceNameHigh string) string{
	return fmt.Sprintf(
`package %s


type RepoType struct {
	database string
}

type i%sRepo interface{
	getAll%s(string)([]%sModel, error)
	get%sById(string)(%sModel, error)
	create%s(%sModel) error
	update%s(string) error
	delete%s(string) error
}

var std%sRepo RepoType

func GetStd%sRepo() RepoType {
	return std%sRepo
}

func (r RepoType)getAll%s(string)([]%sModel, error){

	return []%sModel{}, nil
}

func (r RepoType)get%sById(string)(%sModel, error){

	return %sModel{}, nil
}

func (r RepoType)create%s(%sModel) error{

	return nil
}

func (r RepoType)update%s(string) error{

	return nil
}

func (r RepoType)delete%s(string) error{

	return nil
}	
`,resourceNameLow,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh,
	resourceNameHigh)
}

func getServiceFileContent(resourceNameLow string, resourceNameHigh string) string{
	return fmt.Sprintf(
`package %s

var %sService %sServiceType

func Init(){
	%sService = get%sService(std%sRepo)
}

type %sServiceType struct {
	repo i%sRepo
}

type i%sService interface {
	get%sById(string) (%sModel, error)
	getAll%ss() (%sModel, error)
	create%ss(%sModel) error
	update%ss() error
	delete%ss() error
}

func get%sService(repo i%sRepo) %sServiceType{
	return %sServiceType{repo: repo}
}


func (s %sServiceType)getAll%s(string)([]%sModel, error){

	return []%sModel{}, nil
}

func (s %sServiceType)get%sById(string)(%sModel, error){

	return %sModel{}, nil
}

func (s %sServiceType)create%s(%sModel) error{

	return nil
}

func (s %sServiceType)update%s(string) error{

	return nil
}

func (s %sServiceType)delete%s(string) error{

	return nil
}	
`,
resourceNameLow,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh)
}

func getControllerFileContent(resourceNameLow string, resourceNameHigh string) string{
	return fmt.Sprintf(
`package %s

import (
	"github.com/gin-gonic/gin"
)

func register%sRoute(r *gin.RouterGroup){

	v := r.Group("%ss")

	v.GET("", getAll%ss)

	v.GET("/:id", get%sById)

	v.POST("", create%s)

	v.PUT("/:id", update%s)

	v.DELETE("/:id", delete%s)

}

func getAll%ss(ctx *gin.Context){

}

func get%sById(ctx *gin.Context){

}

func create%s(ctx *gin.Context){

}

func update%s(ctx *gin.Context){

}

func delete%s(ctx *gin.Context){

}`,
resourceNameLow,
resourceNameHigh,
resourceNameLow,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
resourceNameHigh,
)
}

func returnCaptalized(word string) string{
	bword := []byte(word)
	if bword[0] >= byte(97){
		bword[0] = bword[0] - 32
	}
	return string(bword)
}

func returnLow(word string) string{
	bword := []byte(word)
	if bword[0] < byte(97){
		bword[0] = bword[0] + 32
	}
	return string(bword)
}


