package service

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"期待功能", "赚豆子点数", "我要退出"}

type model struct {
	cursor int
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("😣，您的豆子点数不足，您可以选择其中一项继续操作？\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(按 q 退出，按 ↑/↓ 选择)\n")

	return s.String()
}

func RunTui() {
	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("嗷，报错了~", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		fmt.Printf("\n---\n您选择了 %s!\n", m.choice)
		switch m.cursor {
		case 0:
			fmt.Println("即将实现，请请期待！")
		case 1:
			fmt.Println("扫描小程序码，看广告赚豆子点数。")
			ShowAdCode()
		case 2:
			fmt.Println("谢谢您的使用！")
		}
		os.Exit(0)
	}
}
