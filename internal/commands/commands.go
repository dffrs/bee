package internal

type Command func(args []string) (string, error)
