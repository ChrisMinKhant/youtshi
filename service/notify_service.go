package service

import "log"

type NotifyService struct {
}

func (notifyService *NotifyService) DropMessageToKafka(request []any) {
	log.Printf("This is requested payload >>> %v,%v", request...)
}
