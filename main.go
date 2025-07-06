package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/amiamiyamamoto/gong/internal/cmd"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:embed assets/gong.mp3
var gongSound []byte // 音源データがこの変数にバイト配列として読み込まれる

func playCommand(args []string) error {
	fmt.Println("🔔 Playing gong sound...")

	// MP3データの存在確認
	if len(gongSound) == 0 {
		return fmt.Errorf("gong sound data is empty")
	}

	fmt.Printf("Loaded %d bytes of gong sound data\n", len(gongSound))

	// MP3データをデコード
	streamer, format, err := mp3.Decode(io.NopCloser(bytes.NewReader(gongSound)))
	if err != nil {
		return fmt.Errorf("failed to decode MP3: %v", err)
	}
	defer streamer.Close()

	fmt.Printf("Sample Rate: %d, Channels: %d\n", format.SampleRate, format.NumChannels)

	// スピーカーを初期化
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return fmt.Errorf("failed to initialize speaker: %v", err)
	}

	fmt.Println("Playing...")

	// 再生完了を待つためのチャンネル
	done := make(chan bool)

	// 音声を再生
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// 再生完了を待つ
	<-done

	fmt.Println("✅ Gong sound finished playing!")
	return nil
}
func main() {
	// Create command registry
	registry := cmd.NewRegistry()

	// Register commands
	registry.Register(&cmd.Command{
		Name:        "help",
		Description: "Show help information",
		Handler:     cmd.HelpCommand,
	})

	registry.Register(&cmd.Command{
		Name:        "version",
		Description: "Show version information",
		Handler:     cmd.VersionCommand,
	})

	registry.Register(&cmd.Command{
		Name:        "play",
		Description: "Play gong sound",
		Handler:     playCommand,
	})

	// Parse command line arguments
	// If no arguments provided, default to playing gong sound
	if len(os.Args) < 2 {
		// Default action: play gong sound
		if err := playCommand([]string{}); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	commandName := os.Args[1]
	args := os.Args[2:]

	// Execute command
	if err := registry.Execute(commandName, args); err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Run 'gong help' for usage information")
		os.Exit(1)
	}
}
