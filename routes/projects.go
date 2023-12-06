package routes

import (
	"html/template"

	"github.com/TrevorEdris/about-me/controller"
	"github.com/TrevorEdris/about-me/embedded"

	"github.com/labstack/echo/v4"
)

type (
	Projects struct {
		controller.Controller
	}

	Project struct {
		Title  string
		Status string
		Body   string
		Image  template.HTML
		Link   template.HTML
		Card   template.HTML
		Notes  []string
	}
)

func (c *Projects) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "projects"
	page.Metatags.Description = "Welcome to the projects page."
	page.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software"}
	page.Pager = controller.NewPager(ctx, 3)
	page.Data = c.fetchProjects(&page.Pager)

	return c.RenderPage(ctx, page)
}

// fetchProjects returns a list of Project, with paging enabled.
func (c *Projects) fetchProjects(pager *controller.Pager) []Project {
	projects := []Project{
		{
			Title:  "About Me (This website)",
			Body:   "A website with a full Go backend, utilizing gohtml templates, hosted via AWS App Runner.",
			Status: "Operational",
			Image:  template.HTML(embedded.NewImg(embedded.ImgPathMyFace, "Trevor Edris").HTML),
			Link:   template.HTML(embedded.NewLink("https://github.com/TrevorEdris/about-me", "Github Repo").HTML),
			Card:   template.HTML(embedded.NewGHRepoCard("TrevorEdris", "about-me").HTML),
			Notes: []string{
				"DNS proxy to the domain https://trevoredris.com is complicated due to Route53 not natively supporting DNS A records for the AWS App Runner service",
				"Project is fully runnable locally, including a mock SMTP server, using a live-reloading service to detect changes to source code files",
			},
		},
		{
			Title:  "Go CSV",
			Body:   "A Go project to generate a CSV full of random, but controllable data",
			Status: "Operational",
			Image:  template.HTML(embedded.NewImg(embedded.ImgPathMyFace, "Trevor Edris").HTML),
			Link:   template.HTML(embedded.NewLink("https://github.com/TrevorEdris/go-csv", "Github Repo").HTML),
			Card:   template.HTML(embedded.NewGHRepoCard("TrevorEdris", "go-csv").HTML),
			Notes: []string{
				"Address reported issues",
				"Host as a service online",
				"More seamless mocks",
			},
		},
		{
			Title:  "Banner",
			Body:   "A simple Go project to generate a banner of text",
			Status: "Operational",
			Image:  template.HTML(embedded.NewImg(embedded.ImgPathMyFace, "Trevor Edris").HTML),
			Link:   template.HTML(embedded.NewLink("https://github.com/TrevorEdris/banner", "Github Repo").HTML),
			Card:   template.HTML(embedded.NewGHRepoCard("TrevorEdris", "banner").HTML),
			Notes: []string{
				"Address reported issues",
			},
		},
		{
			Title:  "Youtube Dependency Graph",
			Body:   "An API with (eventually) a front-end to display a dependency graph of Youtube videos based on the video links within the description.",
			Status: "Locally Deployable",
			Image:  template.HTML(embedded.NewImg(embedded.ImgPathYTLogo, "Youtube Logo").HTML),
			Link:   template.HTML(embedded.NewLink("https://github.com/TrevorEdris/youtube-dependency-graph", "Github Repo").HTML),
			Card:   template.HTML(embedded.NewGHRepoCard("TrevorEdris", "youtube-dependency-graph").HTML),
			Notes: []string{
				"Front-end will be created using a simple React app as well as the D3.js library to visualize the graph",
				"Graph data structure adheres to the JSON Graph Format v2",
				"(Planned) Persist graph into storage to avoid querying full structure",
			},
		},
	}
	pager.SetItems(len(projects))

	return projects[pager.GetOffset() : pager.GetOffset()+pager.ItemsPerPage]
}
