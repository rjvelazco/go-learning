package tasktracer

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

func CreateTask(id int, title string, description string) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      "new",
	}
}

func (t *Task) UpdateStatus(status string) {
	t.Status = status
}

func (t *Task) UpdateTitle(title string) {
	t.Title = title
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
}
