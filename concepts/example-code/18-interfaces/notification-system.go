package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel int
	isUrgent bool
}

type groupMessage struct {
	groupName string
	messageContent string
	priorityLevel int
}

type systemAlert struct {
	alertCode string
	messageContent string
}

func (dm directMessage) importance() int {

	if dm.isUrgent {
		return 50
	}
	return dm.priorityLevel
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string,int){

	switch notifyType := n.(type) {
		case directMessage : 
			return notifyType.senderUsername,notifyType.importance()
		case groupMessage :
			return notifyType.groupName,notifyType.importance()
		case systemAlert : 
			return notifyType.alertCode,notifyType.importance()
		default :
			return "",0	
	}
	
}