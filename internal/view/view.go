package view

import (
	"context"
	"io"
	"net/http"
	"path/filepath"
	"snake-scape/internal/component"
	"snake-scape/internal/middleware"
	"snake-scape/internal/template"

	"github.com/a-h/templ"
)

func ServeFavicon(w http.ResponseWriter, r *http.Request) {
	filePath := "favicon.ico"
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func Home(ctx *middleware.CustomContext, w http.ResponseWriter, r *http.Request) {
	template.Base(
		"Snakescape - Build RuneScape bots with Python",
		[]templ.Component{
			component.TextAndTitle(
			"Giving You the Tools",
				component.Paragraph(
					"This guide will walk you through how to get started building you own bots for RuneScape using python.",
				),
				component.Paragraph(
					"This guide focuses on OldSchool RuneScape, but the skills learned here can be applied to any form of GUI automation.",
				),
				templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					_, err := io.WriteString(w, `
						<p class='text-sm'>All code examples used in this site can be found <a class='text-blue-800 underline' href='https://github.com/phillip-england/snakescape'>in this repository</a> for public use.</p>
					`)
					return err
				}),
			),
			component.TextAndTitle(
				"OldSchool RuneScape",
				component.Paragraph(
					"OldSchool RuneScape is a popular MMORPG that has been around since 2013. If you ended up here because you play the game and want to build bots, you can skip to the next section.",
				),
				templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					_, err := io.WriteString(w, `
						<p class='text-sm'>If you are here because you want to learn more about GUI automation, you can sign-up for an account <a class='text-blue-800 underline' href='https://oldschool.runescape.com/'>here</a> and follow along.</p>
					`)
					return err
				}),
			),
			component.TextAndTitle(
				"Assumptions Made",
				component.Paragraph(
					"This guide is not an introuduction to programming, and assumes you have a basic understanding of python and programming principles.",
				),
				component.Paragraph(
					"However, if you are a novice programmer, this guide will definitely help you to learn more about programming skills, specifically problem solving. A lot of the challenges we face in building bots will have novel solutions, and will expose you to new ways of thinking about general programming problems.",
				),
				component.Paragraph(
					"To get the most out of this guide, you will need a solid grasp of the following concepts:",
				),
				component.UnorderedList(
					"terminal commands",
					"virtual environments",
					"pip and package management",
					"loops, conditionals, functions, and lists",
				),
				templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					_, err := io.WriteString(w, `
						<p class='text-sm'>
							If you are new to programming, I would recommend starting with <a class='text-blue-800 underline' href='https://www.codecademy.com/learn/learn-python-3'>Codecademy's Python 3 course</a>. It will teach you everything you need to get going and follow along with this guide.
						</p>
					`)
					return err
				}),
				component.Paragraph("Now that we have that out of the way, let's get started!"),
				component.LinkButton("Start Building Bots!", "/getting-started"),
			),
		},
	).Render(ctx, w)
}