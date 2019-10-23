package lyric

import (
	"bufio"
	"bytes"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

//Sentence lyric sentence
type Sentence struct {
	ScaleId int
	Scale   string
	Text    string
}

type Lyric struct {
	filename  string
	musicName string
	seLen     int
	content   map[int]*Sentence
}

//New create Lyric by musicName name and its filename
func New(musicName string, filename string) *Lyric {
	lyc := &Lyric{
		filename:  filename,
		musicName: musicName,
		content:   make(map[int]*Sentence),
	}
	return lyc
}

//GetContent get lyric content
func (lyc *Lyric) GetContent() map[int]*Sentence {
	return lyc.content
}

func (lyc *Lyric) NextSentence(ScaleId int) *Sentence {
	switch {
	case ScaleId < 0:
		return lyc.content[0]
	default:
		ScaleId++
		if ScaleId > lyc.seLen-1 {
			return lyc.content[0]
		}
		return lyc.content[ScaleId]
	}
}

//Parse lyric
func (lyc *Lyric) Parse() (err error) {
	f, err := os.Open(lyc.filename)
	if err != nil {
		return err
	}
	sc := bufio.NewScanner(f)

	id := 0
	for sc.Scan() {
		var t, st []byte
		row := sc.Bytes()
		left := bytes.IndexByte(row, '[')
		right := bytes.IndexByte(row, ']')
		if left == -1 || right == -1 {
			continue
		}
		// fill lyric scales
		t = row[left+1 : right]
		if len(row[right+1:]) > 0 {
			st = row[right+1:]
		}
		lyc.content[id] = &Sentence{
			ScaleId: id,
			Scale:   string(t),
			Text:    string(st),
		}
		id++
	}
	lyc.seLen = id
	return nil
}

//SubStrTime sub two str time like
// 	t1=13:14:15,t2=01:02:03 or
//  t1=14:15, t2=15:16, equal t1=00:14:15,t2=00:15:16
func SubStrTime(t1, t2 string) time.Duration {
	return time.Duration(math.Abs(mills(t1) - mills(t2)))
}

func mills(t string) float64 {
	var h, m, s string
	st1 := strings.Split(t, ":")
	switch len(st1) {
	case 1: // s
		s = st1[0]
	case 2: // m:s
		m, s = st1[0], st1[1]
	case 3: // h:m:s
		h, m, s = st1[0], st1[1], st1[2]
	default:
		return 0
	}

	var err error
	var fh, fm, fs float64
	if fh, err = strconv.ParseFloat(h, 64); err != nil || fh > 60 || fh < 0 {
		fh = 0
	}
	if fm, err = strconv.ParseFloat(m, 64); err != nil || fm > 60 || fm < 0 {
		fm = 0
	}
	if fs, err = strconv.ParseFloat(s, 64); err != nil || fs > 60 || fs < 0 {
		fs = 0
	}
	return (fh*3600 + fm*60 + fs) * 1e9
}
