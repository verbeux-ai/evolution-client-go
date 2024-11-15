package evolution

const (
	createInstanceEndpoint             = "instance/create"             // POST
	restartInstanceEndpoint            = "instance/restart/%s"         // POST
	logoutInstanceEndpoint             = "instance/logout/%s"          // DELETE
	deleteInstanceEndpoint             = "instance/delete/%s"          // DELETE
	getConnectInstanceEndpoint         = "instance/connect/%s"         // GET
	fetchInstancesEndpoint             = "instance/fetchInstances"     // GET
	getConnectionStateInstanceEndpoint = "instance/connectionState/%s" // GET

	findChatsEndpoint    = "chat/findChats/%s"         // GET
	readMessagesEndpoint = "chat/markMessageAsRead/%s" // POST

	sendMessageTextEndpoint  = "message/sendText/%s"  // POST
	sendMessageMediaEndpoint = "message/sendMedia/%s" // POST

	findLabelsEndpoint  = "label/findLabels/%s"  //GET
	handleLabelEndpoint = "label/handleLabel/%s" // POST
)
