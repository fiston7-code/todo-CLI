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

// update task

// delete task
func (s *Store) DeleteTask(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil

		}

	}
	return fmt.Errorf("impossible de supprimer : le tache avec l'ID %d n'existe pas", id)

}

// MÉTHODE UpdateTask(id int, newTitle string, newDesc string) error
//   → boucle sur les tasks avec index (for i := range)
//   → SI tasks[i].ID == id
//       → modifie tasks[i].Title et tasks[i].Desc
//       → retourne nil
//   → SI rien trouvé → retourne erreur

func (s *Store) UpdateTask(id int, newTitle string, newDesc string) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Title = newTitle
			s.tasks[i].Desc = newDesc
			return nil

		}
	}
	return fmt.Errorf("tâche avec l'ID %d introuvable", id)
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

// MÉTHODE UpdateTask(id int, newTitle string, newDesc string) error
//   → boucle sur les tasks avec index (for i := range)
//   → SI tasks[i].ID == id
//       → modifie tasks[i].Title et tasks[i].Desc
//       → retourne nil
//   → SI rien trouvé → retourne erreur
// MÉTHODES
//   → Create()
//   → Show()
//   → Update()
//   → Delete()
//   → Load()   ← lit tasks.json
//   → Save()   ← écrit tasks.json

// ERREURS
//   → tache introuvable
