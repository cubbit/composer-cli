package status

import "strings"

// terminalWriter parses Bubble Tea's terminal escape sequence output to
// capture the actual rendered content of each frame.
type terminalWriter struct {
	line   strings.Builder
	frames []string
}

func (w *terminalWriter) Write(p []byte) (int, error) {
	n := len(p)
	i := 0

	for i < len(p) {
		b := p[i]

		switch {
		case b == '\033' && i+1 < len(p) && p[i+1] == '[':
			i += 2
			seqStart := i
			for i < len(p) && !(p[i] >= 0x40 && p[i] <= 0x7E) {
				i++
			}
			seq := string(p[seqStart:i])
			if i < len(p) {
				i++
			}

			if seq == "2K" || seq == "K" {
				if w.line.Len() > 0 {
					w.frames = append(w.frames, w.line.String())
				}
				w.line.Reset()
			}

		case b == '\r':
			if w.line.Len() > 0 {
				w.frames = append(w.frames, w.line.String())
			}
			w.line.Reset()
			i++

		case b == '\n':
			i++

		case b < 0x20:
			i++

		default:
			w.line.WriteByte(b)
			i++
		}
	}
	return n, nil
}

func (w *terminalWriter) renderedFrames() []string {
	var result []string
	for _, f := range w.frames {
		if trimmed := strings.TrimSpace(f); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
