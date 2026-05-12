package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fiston7-code/todo-golang/models"
)

func main() {
	store := models.NewStore()
	reader := bufio.NewReader(os.Stdin)
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
	var taskToUpdateById int
	fmt.Print("\nquel numero de tache voulais-vous modifier ? ")
	fmt.Scanln(&taskToUpdateById)

	fmt.Print("quel titre ? ")
	taskToUpdateByTitle, _ := reader.ReadString('\n')
	taskToUpdateByTitle = strings.TrimSpace(taskToUpdateByTitle)

	fmt.Print("quel description ? ")
	taskToUpdateByDesc, _ := reader.ReadString('\n')
	taskToUpdateByDesc = strings.TrimSpace(taskToUpdateByDesc)

	errUpdate := store.UpdateTask(taskToUpdateById, taskToUpdateByTitle, taskToUpdateByDesc)
	if errUpdate != nil {
		fmt.Printf("VOUS NE POUVEZ PAS MODIFIER %v\n", errUpdate)
	} else {
		fmt.Println("✅ Tâche mise à jour avec succès.")
	}
	store.Show()
}
