package internal

type Empty interface {
	Empty() bool
}

func RemoveEmpty(list []Empty) []Empty {
	return nil
}
