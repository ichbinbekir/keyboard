package keyboard

type Event interface{}

type ButtonEvent struct {
	KeyCode  uint32
	KeyState bool
}

// WheelEvent, fare tekerleği olayını temsil eder.
type WheelEvent struct {
	// Delta, tekerleğin ne kadar döndürüldüğünü belirtir.
	// Pozitif bir değer tekerleğin ileri (kullanıcıdan uzağa) döndürüldüğünü gösterir.
	// Negatif bir değer, tekerleğin geriye (kullanıcıya doğru) döndürüldüğünü gösterir.
	Delta int16
}

// MouseMoveEvent, fare hareket olayını temsil eder.
type MouseMoveEvent struct {
	X, Y int32
}
