package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)


type Page struct{
	Title string 
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename , p.Body , 0600)
}

func loadPage(title string) (*Page , error){
	filename := title + ".txt"
	body , err := os.ReadFile(filename)

	if err != nil{
		return nil , err
	}

	return &Page{ Title: title,Body: body, } , nil
}

func viewHandler( w http.ResponseWriter , r *http.Request ){

	title := r.URL.Path[len("/view/"):]

	p , _ := loadPage(title)


	fmt.Fprintf( w, " <h1>%s</h1>  <h5>%s</h5>" , p.Title , p.Body)
}


func main() {

	// p1 := &Page{ Title: "Avengers" , Body: []byte("Avengers assemble") }
	// p1.save()

	// p2 , err:= loadPage(p1.Title)

	// if err != nil{
	// 	log.Fatal(err)
	// }

	// fmt.Println( string(p2.Body) )


	http.HandleFunc("/view/" , viewHandler)
	log.Fatal(http.ListenAndServe(":8080" , nil))

	
}