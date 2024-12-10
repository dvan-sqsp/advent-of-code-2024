package day09

type File struct {
	ID, StartIdx, EndIdx, Length int
}

func NewFile(id, startIdx, endIdx, length int) File {
	return File{id, startIdx, endIdx, length}
}
