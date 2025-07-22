package solid

// It says that higher level dependencies should depend on interface of the lower level depdenco
// i.e shoukd not depend on the actual implementation 

type EmailService struct {

}

func (e EmailService) Send() error {
	return nil
}


// type NotificationService struct {
//    emailService EmailService
// }

// now in this case lets say in future we want to have push notification to device then we might to change 
// the class variable which will make it tougher to add code if else woould be needed on extra variabke 
// we can do something like this take this  out to interface 

type NotificationService struct {
	senderService SenderService
}

type SenderService interface {
	Send() error 
}


type PushInboxService struct {

}

func (p PushInboxService) Send() error {
	return nil
}