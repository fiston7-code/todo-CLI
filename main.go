package main

import (
	"fmt"

	"github.com/fiston7-code/todo-golang/models"
)

func main() {
	store := models.NewStore()
	store.CreateTask("Tester ", "fonctionne")
	store.CreateTask("Tester la date", "Vérifier si le formatage fonctionne")
	store.CreateTask("pause", "manger bien")
	store.CreateTask("travail sur ma amchine", "coder en golang")
	store.Show()
	var taskToDelete int

	fmt.Println("\nquel numero de tache voulais-vous supprimer ?")
	fmt.Scanln(&taskToDelete)

	errSupprimer := store.DeleteTask(taskToDelete)

	if errSupprimer != nil {
		fmt.Printf("desole %v ne peut etre supprimer\n", errSupprimer)

	} else {
		fmt.Println("✅ Contact supprimé avec succès.")
		store.Show()

	}

}
