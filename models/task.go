package models

import (
	"fmt"
	"time"
)

type Task struct {
	ID        int
	Title     string
	Desc      string
	CreatedAt time.Time
}

type Store struct {
	tasks  []Task
	nextID int
}

func NewStore() *Store {
	return &Store{
		tasks:  []Task{},
		nextID: 1,
	}

}

func (s *Store) CreateTask(title string, desc string) {
	newTask := Task{
		ID:        s.nextID,
		Title:     title,
		Desc:      desc,
		CreatedAt: time.Now(),
	}

	s.nextID++
	s.tasks = append(s.tasks, newTask)

}

func (s *Store) Show() {
	if len(s.tasks) == 0 {
		fmt.Println("aucune tâche")
		return
	}
	for _, task := range s.tasks {
		fmt.Printf("%-3d | %-15s | %-20s | %s\n",
			task.ID, task.Title, task.Desc,
			task.CreatedAt.Format("02/01/2006 15:04"))
	}

}

// STRUCT Tache
//   → ID, Title, Desc, CreatedAt

// STRUCT Store
//   → taches []Tache
//   → nextID int

// CONSTRUCTEUR NewStore()
//   → appelle Load() au démarrage

// MÉTHODE Show() sur Store
//   → SI len(tasks) == 0 → affiche "aucune tâche"
//   → SINON boucle sur les tasks
//   → Pour chaque task :
//      fmt.Printf avec ID | Title | Desc | CreatedAt.Format(...)
// MÉTHODES
//   → Create()
//   → Show()
//   → Update()
//   → Delete()
//   → Load()   ← lit tasks.json
//   → Save()   ← écrit tasks.json

// ERREURS
//   → tache introuvable
