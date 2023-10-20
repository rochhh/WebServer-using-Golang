package main

import (
	// "fmt"
	"log"
	"os"
	"net/http"
	"html/template"
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
	// fmt.Println(len("/view/"))  -> 6
	
	p , err := loadPage(title)

	if err != nil {
		http.Redirect(w , r , "/edit/"+title , http.StatusNotFound )
	}

	t , _ := template.ParseFiles("view.html")

	t.Execute(w,p)


}


func editHandler( w http.ResponseWriter , r *http.Request ){

	title := r.URL.Path[ len("/view/") : ]

	p , err := loadPage(title)

	if err != nil{
		p = &Page{ Title: title  }
	}

	t , _ := template.ParseFiles("edit.html")
	t.Execute(w , p)

}


func saveHandler( w http.ResponseWriter , r *http.Request ){

	title := r.URL.Path[len("/save/") : ]

	body := r.FormValue("body")

	p := &Page{ Title: title , Body: []byte(body) }
	p.save()

	http.Redirect(w,r , "/view/"+title , http.StatusNotFound )

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
	http.HandleFunc("/edit/" , editHandler)
	http.HandleFunc("/save/" , saveHandler)
	log.Fatal(http.ListenAndServe(":8080" , nil))

	
}