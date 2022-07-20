package builder

import "errors"

type INotification interface {
	SetTitle(title NotiTitle) *Notification
	SetBody(body NotiBody) *Notification
	Send() (err error)
}

type Notification struct {
	Title NotiTitle
	Body  NotiBody
}

type NotiTitle string

const (
	InfoTitle  NotiTitle = "info"
	WarmTitle  NotiTitle = "warm"
	ErrorTitle NotiTitle = "error"
)

type NotiBody string

const (
	InfoBody  NotiBody = "Hi, this is info"
	WarmBody  NotiBody = "Here have some warming, you should take a look"
	ErrorBody NotiBody = "Here got an error, please fixed it quickly"
)

func NewNotification() INotification {
	return &Notification{}
}

func (noti *Notification) SetTitle(title NotiTitle) *Notification {
	noti.Title = title
	return noti
}

func (noti *Notification) SetBody(body NotiBody) *Notification {
	return noti
}

func (noti *Notification) Send() (err error) {
	if len(noti.Title) == 0 {
		return errors.New("title is empty")
	}
	if len(noti.Body) == 0 {
		return errors.New("body is empty")
	}
	noti.saveToDB()
	return nil
}

func (noti *Notification) saveToDB() {

}
