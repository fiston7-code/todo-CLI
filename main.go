package main

import (
	"github.com/fiston7-code/todo-golang/models"
)

func main() {
	store := models.NewStore()
	store.CreateTask("Tester la date", "Vérifier si le formatage fonctionne")
	store.Show() // La date doit apparaître ici
}
