package book

type Book struct {
	title  string
	author string
}

type BookManager struct {
	books []Book
}

func (bm *BookManager) AddBookByBook(book Book) {
	bm.books = append(bm.books, book)
}

func (bm *BookManager) AddBookByTitleAndAuthor(_title, _author string) {
	book := Book{
		title:  _title,
		author: _author,
	}
	bm.books = append(bm.books, book)
}

func (bm *BookManager) ListBooks() []Book {
	return bm.books
}

func (bm *BookManager) FindBookByAuthor(_author string) *Book {
	for i := range bm.books {
		if bm.books[i].author == _author {
			return &bm.books[i]
		}
	}
	return nil
}

func (bm *BookManager) FindBookByTitle(_title string) *Book {
	for i := range bm.books {
		if bm.books[i].title == _title {
			return &bm.books[i]
		}
	}
	return nil
}

func (bm *BookManager) CheckOutBook(index int) *BookManager {
	if index < 0 || index >= len(bm.books) {
		return nil
	}
	book := bm.books[index]
	bm.books = append(bm.books[:index], bm.books[index+1:]...)
	return &BookManager{books: []Book{book}}
}
