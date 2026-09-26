package internal

type Command struct {
	Name string
	Run  func(args []string) (string, error)
}
