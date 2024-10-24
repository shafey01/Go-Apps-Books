module github.com/shafey01/Go-Apps-Books/go-web-programming-Book/coffe-app

go 1.22.2

replace github.com/shafey01/Go-Apps-Books/go-web-programming-Book/coffe-app/handlers => /home/shafey/learning-go/Go-Apps-Books/coffe-app/handlers

replace github.com/shafey01/Go-Apps-Books/go-web-programming-Book/coffe-app/data => /home/shafey/learning-go/Go-Apps-Books/coffe-app/data

require github.com/shafey01/Go-Apps-Books/go-web-programming-Book/coffe-app/handlers v0.0.0-00010101000000-000000000000

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/shafey01/Go-Apps-Books/go-web-programming-Book/coffe-app/data v0.0.0-00010101000000-000000000000 // indirect
)
