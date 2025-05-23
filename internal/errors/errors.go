package errors

type UrlNotFoundError struct{}

func (err UrlNotFoundError) Error() string {
	return "Url Not Found"
}

type TgChatNotFoundError struct{}

func (err TgChatNotFoundError) Error() string {
	return "TgChat Not Found"
}

type BadRequestError struct{}

func (err BadRequestError) Error() string {
	return "Bad Request"
}

type UnknownCommand struct{}

func (err UnknownCommand) Error() string {
	return "I don't know that command. Enter /help"
}
