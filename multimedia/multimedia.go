package multimedia

import "fmt"

type ContenidoWeb struct{
	MultimediaS []Multimedia
}

func (cw *ContenidoWeb) Mostrar(){
	cw.Mostrar()
}

type Multimedia interface{
	Mostrar()
}

type Imagen struct{
  Titulo string
  Formato string
  Canales string
}

func (i *Imagen) Mostrar(){
  fmt.Println(i.Titulo)
  fmt.Println(i.Formato)
  fmt.Println(i.Canales)
}

type Audio struct{
  Titulo string
  Formato string
  Duracion string
}

func (a *Audio) Mostrar(){
  fmt.Println(a.Titulo)
  fmt.Println(a.Formato)
  fmt.Println(a.Duracion)
}

type Video struct{
  Titulo string
  Formato string
  Frames int64
}

func (v *Video) Mostrar(){
  fmt.Println(v.Titulo)
  fmt.Println(v.Formato)
  fmt.Println(v.Frames)
}