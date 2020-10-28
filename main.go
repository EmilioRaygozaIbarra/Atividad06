package main

import (
	"./multimedia"
	"fmt"
)

func main() {	
	var tituloI, tituloA, tituloV, formatoI, formatoA, formatoV, duracionA, canalesI string
	var o, framesV int64
	
	for o != 5{
		i := multimedia.Imagen{}
		a := multimedia.Audio{}
		v := multimedia.Video{}
		fmt.Println("Opciones: ")
		fmt.Println("1.- Ingresar Imagen")
		fmt.Println("2.- Ingresar Audio")
		fmt.Println("3.- Ingresar Video")
		fmt.Println("4.- Mostrar Todo")
		fmt.Println("5.- Salir")
		fmt.Scan(&o)
		if o==1{
			fmt.Println("Titulo de la imagen")
			fmt.Scan(&tituloI)
			fmt.Println("Formato de la imagen")
			fmt.Scan(&formatoI)
			fmt.Println("Canales de la imagen")
			fmt.Scan(&canalesI)
			i.Titulo=tituloI
			i.Formato=formatoI
			i.Canales=canalesI
		}else if o==2{
			fmt.Println("Titulo del audio")
			fmt.Scan(&tituloA)
			fmt.Println("Formato del audio")
			fmt.Scan(&formatoA)
			fmt.Println("Duracion del audio")
			fmt.Scan(&duracionA)
			a.Titulo=tituloA
			a.Formato=formatoA
			a.Duracion=duracionA
		}else if o==3{
			fmt.Println("Titulo del video")
			fmt.Scan(&tituloV)
			fmt.Println("Formato del video")
			fmt.Scan(&formatoV)
			fmt.Println("Frames del video")
			fmt.Scan(&framesV)
			v.Titulo=tituloV
			v.Formato=formatoV
			v.Frames=framesV
		}else if o==4{
			cw := multimedia.ContenidoWeb{
				MultimediaS: []multimedia.Multimedia{
					&i,
					&a,
					&v,
				},
			}
			cw.Mostrar()
		}
	}
}