//go:build linux
// +build linux

package keyboard

import (
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
	"syscall"
)

// platformSpecific is empty on Linux, but defined to satisfy the Keyboard struct.
type platformSpecific struct{}

// inputEvent, Linux'un girdi olaylarını temsil eden yapıdır.
// /usr/include/linux/input.h dosyasındaki 'struct input_event' yapısına karşılık gelir.
type inputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

const (
	// Olay tipleri
	evKey  = 0x01 // Tuş veya fare düğmesi olayı
	evRel  = 0x02 // Göreceli eksen olayı (fare tekerleği gibi)
	evSyn  = 0x00 // Senkronizasyon olayı

	// Göreceli eksen kodları
	relWheel = 0x08 // Fare tekerleği
)

// findInputDevices, /dev/input altındaki olay aygıtlarını bulur.
func findInputDevices(pattern string) []string {
	matches, _ := filepath.Glob(pattern)
	return matches
}

// readEvents, girdi olaylarını okumak için bir veya daha fazla goroutine başlatır.
// DİKKAT: Bu fonksiyonun çalışabilmesi için programın /dev/input/event* dosyalarına
// okuma izni olması gerekir. Bu genellikle programı 'root' olarak çalıştırmayı veya
// kullanıcıyı 'input' grubuna eklemeyi gerektirir.
func (kb *Keyboard) readEvents() {
	handleMouse := kb.config.HandleMouseButtons || kb.config.HandleMouseWheel || kb.config.HandleMouseMove
	if !kb.config.HandleKeyboard && !handleMouse {
		kb.Errors <- errors.New("klavye veya fare olaylarından en az biri izlenmelidir")
		return
	}

	if kb.config.HandleMouseMove {
		// /dev/input arayüzü genellikle göreceli fare hareketlerini bildirir.
		// Windows versiyonu ise mutlak ekran koordinatlarını kullandığı için,
		// tutarlılığı korumak amacıyla bu özellik Linux'ta şimdilik uygulanmamıştır.
		kb.Errors <- errors.New("HandleMouseMove özelliği Linux'ta henüz desteklenmemektedir")
	}

	devices := findInputDevices("/dev/input/event*")
	if len(devices) == 0 {
		kb.Errors <- errors.New("/dev/input altında hiçbir girdi aygıtı bulunamadı")
		return
	}

	// Her aygıt için ayrı bir olay okuma goroutine'i başlat.
	for _, device := range devices {
		go kb.readDevice(device)
	}

	// readEvents fonksiyonunun ana goroutine'i sonlandırmasını engelle.
	select {}
}

// readDevice, belirtilen aygıt dosyasından sürekli olarak girdi olaylarını okur.
func (kb *Keyboard) readDevice(deviceFile string) {
	file, err := os.Open(deviceFile)
	if err != nil {
		// İzin hataları gibi nedenlerle açılamayan aygıtları sessizce yoksay.
		return
	}
	defer file.Close()

	var event inputEvent
	for {
		// Aygıt dosyasından bir olay oku.
		err := binary.Read(file, binary.LittleEndian, &event)
		if err != nil {
			// Aygıt çıkarılmış veya bir hata oluşmuş olabilir. Bu aygıtı dinlemeyi bırak.
			return
		}

		switch event.Type {
		case evKey:
			if !kb.config.HandleKeyboard && !kb.config.HandleMouseButtons {
				continue
			}

			// event.Value: 0 = bırakıldı, 1 = basıldı, 2 = tekrar
			isDown := event.Value != 0

			// Tekrarlanan tuş basımlarını gönderme ayarı kapalıysa, bu olayları atla.
			if event.Value == 2 && !kb.config.SendRepeatedKeyDowns {
				continue
			}

			// DİKKAT: event.Code, Linux'a özgü bir tuş kodudur (<linux/input-event-codes.h>).
			// Bu, Windows'taki Virtual Key Code (VK_*) ile aynı DEĞİLDİR.
			// Fare tuşları da (BTN_LEFT, BTN_RIGHT vb.) bu olay tipini kullanır.
			kb.keyStates[uint32(event.Code)] = isDown
			kb.Events <- ButtonEvent{KeyCode: uint32(event.Code), KeyState: isDown}

		case evRel:
			if kb.config.HandleMouseWheel && event.Code == relWheel {
				// event.Value, tekerlek hareketinin yönünü ve miktarını belirtir.
				// Genellikle ileri için 1, geri için -1 değerini alır.
				kb.Events <- WheelEvent{Delta: int16(event.Value)}
			}
		}
	}
}

// getState, Linux'ta /dev/input arayüzü üzerinden anlık sorgulanamadığı için
// bu platformda desteklenmemektedir.
func getState(key int) bool {
	// TODO: Olası bir implementasyon, X11 veya başka bir display server API'si kullanmayı gerektirebilir.
	return false
}

// sendInput, Linux'ta /dev/uinput üzerinden karmaşık bir kurulum gerektirdiği için
// bu platformda desteklenmemektedir.
func sendInput(key int, state bool) {
	// TODO: Olası bir implementasyon, /dev/uinput veya X11 Test Extension kullanmayı gerektirebilir.
}
