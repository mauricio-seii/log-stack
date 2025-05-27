package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ProcessLogFile(logPath, udpAddr, parsedDir string) error {
	log.Printf("[parser] Opening file: %s", logPath)

	f, err := os.Open(logPath)
	if err != nil {
		return fmt.Errorf("[parser] failed to open file: %w", err)
	}
	defer f.Close()

	conn, err := net.Dial("udp", udpAddr)
	if err != nil {
		return fmt.Errorf("[parser] failed to connect to UDP address %s: %w", udpAddr, err)
	}
	defer conn.Close()

	env := parseEnvFromFilename(filepath.Base(logPath))
	log.Printf("[parser] Parsed environment from filename: %s", env)

	serviceName := parseServiceName(filepath.Base(logPath))
	log.Printf("[parser] Parsed service name from filename: %s", serviceName)

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 1024*1024) // 1MB initial buffer
	scanner.Buffer(buf, 1024*1024)    // max token size 1MB
	for scanner.Scan() {
		line := scanner.Text()
		logLine := fmt.Sprintf("{\"service\":\"%s\",\"environment\":\"%s\",\"message\":\"%s\"}\n", serviceName, env, escapeJSON(line))
		_, err := conn.Write([]byte(logLine))
		if err != nil {
			return fmt.Errorf("[parser] failed to send UDP data: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("[parser] scanner error: %w", err)
	}

	destPath := filepath.Join(parsedDir, filepath.Base(logPath))
	log.Printf("[parser] Moving file to parsed directory: %s", destPath)
	err = copyFile(logPath, destPath)
	if err != nil {
		return fmt.Errorf("[parser] failed to copy file to parsed directory: %w", err)
	}
	return os.Remove(logPath)
}

func parseEnvFromFilename(name string) string {
	re := regexp.MustCompile(`(?i)\.(dev|staging|prod)\.`)
	match := re.FindStringSubmatch(name)
	if len(match) > 1 {
		return match[1]
	}
	return "unknown"
}

func parseServiceName(name string) string {
	serviceName := strings.Split(name, ".")
	return serviceName[0]
}

func escapeJSON(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Sync()
}
