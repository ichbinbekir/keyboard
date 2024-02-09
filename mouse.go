package keyboard

func MouseEvent(event uintptr) error {
	if _, _, err := _mouse_event.Call(event); err != errOperationComleted {
		return err
	}

	return nil
}
