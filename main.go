package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "os"
  "os/exec"
  "strings"

  "github.com/gomarkdown/markdown"
  "github.com/gomarkdown/markdown/html"
  "github.com/thatisuday/commando"
)

func execute(command string) string {
  out, err := exec.Command("bash", "-c", command).Output()

  if err != nil {
    fmt.Println("%s", err)
  }

  output := string(out[:])
  return output
}

func main(){
  commando.
  SetExecutableName("markpdf").
  SetVersion("v1.0.0").
  SetDescription("A CLI tool to transform markdown files into PDF files")

  commando.
  Register(nil).
  AddArgument("file", "Markdown File", "none").
  AddFlag("theme,t", "Set the theme (colors) for the PDF file", commando.String, "default").
  AddFlag("orientation,o", "Pass the orientation for the PDF file", commando.String, "Portrait").
  AddFlag("size,s", "Change page size of PDF file", commando.String, "a4").
  AddFlag("grayscale,g", "Set grayscale of PDF file", commando.Bool, false).
  AddFlag("path,p", "Set the final PDF file path", commando.String, "default").
  AddFlag("css,c", "CSS append files", commando.String, "default").
  SetAction(markdowntoHTML)

  commando.Parse(nil)
}

func markdowntoHTML(args map[string]commando.ArgValue, flags map[string]commando.FlagValue){

  userHome, err := os.UserHomeDir()
  if err != nil{
    log.Fatal(err)
  }
  packageDirectory := userHome + "/go/src/github.com/edersonferreira/gomarkpdf/"
  defaultThemes := []string{"default", "dark", "abnt", "programmer"}
  cmdFlags := ""
  orientations := []string{"portrait", "landscape"}
  sizes := []string{"a0","a1","a2","a3","a4","a5","a6","a7","a8","a9","letter"}

  markdownFile := args["file"].Value
  themeBrute, _ := flags["theme"].GetString()
  orientation, _ := flags["orientation"].GetString()
  size, _ := flags["size"].GetString()
  grayscale, _ := flags["grayscale"].GetBool()
  path, _ := flags["path"].GetString()
  cssFiles, _ :=   flags["css"].GetString()

  if !Contains(orientations, strings.ToLower(orientation)){
    fmt.Println("The orientation is with an incorrect value (", orientation,") orientation can be only \"Portrait\" or \"Landscape\"")
    os.Exit(0)
  }
  if !Contains(sizes, strings.ToLower(size)){
    fmt.Println("The size is with an incorrect value (" + size + "). Size can be only A0 to A9 or Letter")
    os.Exit(0)
  }

  if grayscale != false{
    cmdFlags += "-g "
  }

  cssLinks := ""

  if (cssFiles != "default"){
    cssFilesSplit := strings.Split(cssFiles, ",")
    for _, css := range cssFilesSplit{
      cssLinks += "<link rel=\"stylesheet\" href=\"" + css + "\">\n"
    }
  }

  cmdFlags += "-O " + orientation + " "
  cmdFlags += "-s " + size + " "

  theme := themeBrute


  if Contains(defaultThemes, themeBrute){
    theme = packageDirectory + "/styles/" + strings.ReplaceAll(themeBrute, ".css", "")

  }

  markdownText, err := ioutil.ReadFile(markdownFile)
  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  flagsHTML := html.CommonFlags | html.CompletePage | html.HrefTargetBlank
  opts := html.RendererOptions{
    Flags: flagsHTML,
  }
  renderer := html.NewRenderer(opts)

  htmlBody := string(markdown.ToHTML(markdownText, nil, renderer))
  // não mexe
  html := "<!DOCTYPE html>\n<html lang=en>\n<meta charset=UTF-8>\n<meta content=\"width=device-width,initial-scale=1\"name=viewport>\n  <link rel=\"stylesheet\" href=\"" + theme + ".css\">\n" + cssLinks + "<body>\n" + htmlBody

  HTMLFile := strings.ReplaceAll(string(markdownFile), "md", "html")

  PDFFile := strings.ReplaceAll(string(markdownFile), "md", "pdf")

  if path != "default"{
    PDFFile = path
  }

  WriteToFile(HTMLFile, html)

  generatePDF(HTMLFile, PDFFile, cmdFlags)
}

func generatePDF(file string, PDFFile string, cmdFlags string ){
  execute("wkhtmltopdf -L 0 -R 0 -B 0 -T 0 " + cmdFlags + "--enable-local-file-access " + file + " " + PDFFile)

  os.Remove(file)
  fmt.Println("Done")
}

func WriteToFile(filename string, data string) error {
  file, err := os.Create(filename)
  if err != nil {
    return err
  }
  defer file.Close()

  _, err = io.WriteString(file, data)
  if err != nil {
    return err
  }
  return file.Sync()
}
func Contains(a []string, x string) bool {
  for _, n := range a {
    if x == n {
      return true
    }
  }
  return false
}
