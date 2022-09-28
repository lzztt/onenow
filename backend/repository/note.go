package repository

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"one.now/backend/entity"
)

const (
	file_pattern = "*-*_*.md"
)

type Metadata struct {
	createTime     int64
	lastUpdateTime int64
	id             uuid.UUID
	file           string
}

func loadMetadata(files []string) []*Metadata {
	m := make(map[uuid.UUID]*Metadata)

	for _, file := range files {
		t := strings.Split(file, "/")
		p := strings.SplitN(t[len(t)-1], "_", 2)
		id := uuid.MustParse(p[0])
		time, err := strconv.ParseInt(p[1][:len(p[1])-3], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if data, ok := m[id]; ok {
			if data.createTime > time {
				data.createTime = time
			}

			if data.lastUpdateTime < time {
				data.lastUpdateTime = time
				data.file = file
			}
		} else {
			m[id] = &Metadata{
				createTime:     time,
				lastUpdateTime: time,
				id:             id,
				file:           file,
			}
		}
	}

	v := make([]*Metadata, 0, len(m))

	for _, val := range m {
		v = append(v, val)
	}

	sort.SliceStable(v, func(i, j int) bool {
		return v[i].createTime < v[j].createTime
	})

	return v
}

func GetNotes(dir string) []*entity.Note {
	m, err := filepath.Glob(filepath.Join(dir, file_pattern))
	if err != nil {
		log.Fatalln(err)
	}
	metadata := loadMetadata(m)

	notes := make([]*entity.Note, len(metadata))
	for i, v := range metadata {
		id := v.id
		data, err := os.ReadFile(v.file)
		if err != nil {
			log.Fatal(err)
		}

		notes[i] = &entity.Note{
			Id:             id,
			Body:           string(data),
			CreateTime:     time.Unix(v.createTime, 0),
			LastUpdateTime: time.Unix(v.lastUpdateTime, 0),
		}
	}

	return notes
}
