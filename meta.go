package meta

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"strings"

	"gopkg.in/yaml.v3"
)

type TaskLevel int32

const (
	LevelCheckin TaskLevel = iota + 1
	LevelEasy
	LevelMedium
	LevelHard
)

type Author struct {
	Name    string `yaml:"name,omitempty"`
	Contact string `yaml:"contact,omitempty"`
}

type TaskType int32

const (
	TypeWeb  TaskType = iota + 1 // 浏览器访问
	TypeNc                       // NC 访问
	TypeFile                     // 纯附件
	TypeExt                      // 外部题目
)

type Task struct {
	Name          string    `yaml:"name,omitempty"`      // 镜像名称
	Type          string    `yaml:"type,omitempty"`      // 镜像类型
	TypeCode      TaskType  `yaml:"type_code,omitempty"` // 题目分类
	Category      string    `yaml:"category,omitempty"`  // 题目分类
	Description   string    `yaml:"description,omitempty"`
	Level         string    `yaml:"level,omitempty"`
	LevelCode     TaskLevel `yaml:"level_code,omitempty"`
	Flag          string    `yaml:"flag,omitempty"`
	AttachmentURL string    `yaml:"attachment_url,omitempty"`
	Hints         []string  `yaml:"hints,omitempty"`
}

type Challenge struct {
	Name  string   `yaml:"name,omitempty"`
	Refer string   `yaml:"refer,omitempty"`
	Tags  []string `yaml:"tags,omitempty"`
}

type Skill struct {
	ID   string `yaml:"id,omitempty"`
	Pid  string `yaml:"pid,omitempty"`
	Tid  string `yaml:"tid,omitempty"`
	Node int    `yaml:"node,omitempty"`
}

type Meta struct {
	Author    Author    `yaml:"author,omitempty"`
	Task      Task      `yaml:"task,omitempty"`
	Challenge Challenge `yaml:"challenge,omitempty"`
	Skill     Skill     `yaml:"skill,omitempty"`
}

func (m *Meta) Format() *Meta {
	// Level
	if m.Task.LevelCode != 0 && m.Task.Level == "" {
		switch m.Task.LevelCode {
		case LevelCheckin:
			m.Task.Level = "签到"
		case LevelEasy:
			m.Task.Level = "简单"
		case LevelMedium:
			m.Task.Level = "中等"
		case LevelHard:
			m.Task.Level = "困难"
		default:
		}
		m.Task.LevelCode = 0
	} else {
		switch strings.ToLower(m.Task.Level) {
		case "签到", "checkin":
			m.Task.LevelCode = LevelCheckin
		case "简单", "easy":
			m.Task.LevelCode = LevelEasy
		case "中等", "中级", "medium":
			m.Task.LevelCode = LevelMedium
		case "困难", "高级", "hard":
			m.Task.LevelCode = LevelHard
		default:
			m.Task.LevelCode = 0
		}
		m.Task.Level = ""
	}
	return m
}

func (m *Meta) parseFormat() *Meta {
	// Type 默认 web
	if m.Task.Type != "" {
		m.Task.Type = strings.ToLower(m.Task.Type)
	} else {
		m.Task.Type = "web"
	}
	// Level 默认 简单easy
	switch strings.ToLower(m.Task.Level) {
	case "签到", "checkin":
		m.Task.LevelCode = LevelCheckin
	case "简单", "easy":
		m.Task.LevelCode = LevelEasy
	case "中等", "中级", "medium":
		m.Task.LevelCode = LevelMedium
	case "困难", "高级", "hard":
		m.Task.LevelCode = LevelHard
	default:
		m.Task.Level = "easy"
		m.Task.LevelCode = LevelEasy
	}
	return m
}

func New(name, contact string) *Meta { return &Meta{Author: Author{Name: name, Contact: contact}} }
func Empty() *Meta                   { return &Meta{} }
func Default() *Meta                 { return New("陌竹", "mozhu233@outlook.com") }
func NewSkill(id, pid, tid string, node int, image, name string, level TaskLevel) *Meta {
	return New("陌竹", "mozhu233@outlook.com").NewSkill(id, pid, tid, node, image, name, level)
}

func (m *Meta) NewSkill(id, pid, tid string, node int, image, name string, level TaskLevel) *Meta {
	n := m.R()
	n.Skill.ID = id
	n.Skill.Pid = pid
	n.Skill.Tid = tid
	n.Skill.Node = node
	n.Challenge.Name = name
	n.Task.Name = image
	n.Task.LevelCode = level
	return n
}

func (m *Meta) R() *Meta {
	n := *m
	return &n
}

func Template() string {
	m := Default()
	m.Task.Name = "challenge_game_2023_web_abc"
	m.Task.Type = "web"
	m.Task.Category = "Web"
	m.Task.Description = "这是一个模板"
	m.Task.Level = "easy"
	m.Task.Flag = "ctftrain{this_is_a_test_flag}"
	m.Task.Hints = []string{"这是一个模板", "没有提示"}

	m.Challenge.Name = "Web题目"
	m.Challenge.Refer = "https://www.ctfhub.com"
	m.Challenge.Tags = []string{"web", "2023"}

	buf, err := yaml.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

func ParseMetaFromFile(fds ...multipart.File) ([]*Meta, error) {
	var data bytes.Buffer
	for _, fd := range fds {
		buf, err := io.ReadAll(fd)
		if err != nil {
			return nil, err
		}
		data.WriteString("---\n")
		data.Write(buf)
	}
	return ParseMetas(data.Bytes())
}

func ParseMetas(data []byte) ([]*Meta, error) {
	var metas []*Meta
	dec := yaml.NewDecoder(bytes.NewReader(data))
	for {
		meta := Meta{}
		err := dec.Decode(&meta)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, err
			}
			break
		}
		metas = append(metas, meta.parseFormat())
	}
	if len(metas) == 0 {
		return metas, errors.New("failed to parse meta data")
	}
	return metas, nil
}
