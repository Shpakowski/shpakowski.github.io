package handlers

// HandleRestart restarts the blockchain node by calling stop and start handlers.
func HandleRestart(_ interface{}) {
	HandleStop(nil)
	HandleStart(nil)
}
