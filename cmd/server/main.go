package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"faraway/internal/config"
	"faraway/internal/lib/logger/sl"
	"faraway/internal/reader"
)

func main() {
	cfg := config.MustLoad()
	log := sl.SetupLogger(cfg.Env)
	log.Info(
		"starting faraway server",
		slog.String("env", cfg.Env),
		slog.String("version", cfg.Version),
	)

	if err := run(log, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(2)
	}
}

func run(log *slog.Logger, cfg *config.Config) error {
	log.Debug("starting server")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	dsn := fmt.Sprintf(":%d", cfg.TCPPort)
	server, err := net.Listen("tcp", dsn)
	if err != nil {
		return err
	}
	defer server.Close()

	log.Info("setting the source of wisdom", "file", cfg.DataFile)

	fileReader, err := reader.NewFileReader(cfg.DataFile)
	if err != nil {
		return err
	}

	log.Info("server is running on port " + dsn)
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Error("Error accepting connection:", err)
				continue
			}

			wg.Add(1)
			go handleClient(log, conn, fileReader, &wg)
		}
	}()

	<-done

	log.Info("Shutting down server...")

	wg.Wait()

	log.Info("Server has been gracefully shut down.")

	return nil
}

func handleClient(log *slog.Logger, conn net.Conn, reader reader.WisdomReader, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()

	log.Info("Client connected:", conn.RemoteAddr())

	// Send request to server.
	message, err := prepareWordOfWisdom(reader)
	if err != nil {
		log.Error("prepare Word Of Wisdom error:", "error", err.Error())
		return
	}

	_, err = conn.Write(message)
	if err != nil {
		log.Error("Sending error:", "error", err.Error())
		return
	}
	log.Debug("Sent data ", "message", string(message))

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Info("Client disconnected:", conn.RemoteAddr())
			return
		}
		log.Debug("Received data from", conn.RemoteAddr(), string(buffer)[:n])

		_, err = conn.Write(buffer)
		if err != nil {
			return
		}
	}
}

func prepareWordOfWisdom(reader reader.WisdomReader) ([]byte, error) {
	result, err := reader.ReadOne()
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}
