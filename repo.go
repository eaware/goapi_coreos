package main

var currentID int

var todos todosT

// Give us some seed data
func init() {
	repoCreateTodo(todoT{Name: "Write presentation"})
	repoCreateTodo(todoT{Name: "Host meetup"})
}

func repoFindTodo(id int) todoT {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	// return empty Todo if not found
	return todoT{}
}

// this is bad, I don't think it passes race condtions
func repoCreateTodo(t todoT) todoT {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

// func repoDestroyTodo(id int) error {
// 	for i, t := range todos {
// 		if t.ID == id {
// 			todos = append(todos[:i], todos[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
// }
