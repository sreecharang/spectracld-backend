package ports 

type DbPort interface {
	CloseDBConnection()
	AddToHistory(answer int32, operation string) error 
}

