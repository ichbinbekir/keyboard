package keyboard

import "testing"

func TestConfigMerge(t *testing.T) {
	// Original test for ChannelSize override
	cfg := mergeConfigs(Config{ChannelSize: 5})
	if cfg.ChannelSize != 5 {
		t.Fatalf("expected ChannelSize to be 5, got %d", cfg.ChannelSize)
	}
	if !cfg.HandleKeyboard { // Default should be true
		t.Fatal("expected HandleKeyboard to be true by default")
	}

	// Original problematic test (modified to check if HandleKeyboard remains true, as expected with !IsZero() logic)
	cfg = mergeConfigs(Config{HandleKeyboard: false}, Config{HandleMouseWheel: true})
	if cfg.HandleKeyboard != DefaultConfig().HandleKeyboard { // HandleKeyboard from first config should be ignored, so it should remain true
		t.Fatalf("expected HandleKeyboard to remain true, got %t", cfg.HandleKeyboard)
	}
	if !cfg.HandleMouseWheel { // HandleMouseWheel from second config should be merged
		t.Fatalf("expected HandleMouseWheel to be true, got %t", cfg.HandleMouseWheel)
	}

	// New test explicitly demonstrating the failure to merge 'false' for boolean fields.
	t.Run("SendRepeatedKeyDowns_False_NotMerged", func(t *testing.T) {
		cfg := mergeConfigs(Config{SendRepeatedKeyDowns: false})
		if cfg.SendRepeatedKeyDowns == DefaultConfig().SendRepeatedKeyDowns { // If it's still the default value (meaning 'false' was ignored)
			t.Fatalf("BUG: mergeConfigs failed to merge 'false' for SendRepeatedKeyDowns. It remained %t (default).", cfg.SendRepeatedKeyDowns)
		}
	})

	t.Run("HandleKeyboard_False_NotMerged", func(t *testing.T) {
		cfg := mergeConfigs(Config{HandleKeyboard: false})
		if cfg.HandleKeyboard == DefaultConfig().HandleKeyboard { // If it's still the default value (meaning 'false' was ignored)
			t.Fatalf("BUG: mergeConfigs failed to merge 'false' for HandleKeyboard. It remained %t (default).", cfg.HandleKeyboard)
		}
	})
}