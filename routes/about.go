package routes

import (
	"fmt"
	"html/template"

	"github.com/TrevorEdris/about-me/controller"
	"github.com/TrevorEdris/about-me/embedded"

	"github.com/labstack/echo/v4"
)

type (
	About struct {
		controller.Controller
	}

	AboutData struct {
		ShowCacheWarning  bool
		FrontendTabs      []AboutTab
		BackendTabs       []AboutTab
		InterestTabs      []AboutTab
		TechnologyTabs    []AboutTab
		QualificationTabs []AboutTab
		CertificationTabs []AboutTab
	}

	AboutTab struct {
		Title string
		Body  template.HTML
	}
)

func (c *About) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "about"
	page.Title = "About Me"

	// This page will be not cached!
	page.Cache.Enabled = false
	page.Cache.Tags = []string{"page_about", "page:list"}

	awsServices := fmt.Sprintf(`The following is a list of the AWS services I use on a near-daily basis.
<ul>
<li>%s S3</li>
<li>%s Batch</li>
<li>%s CloudWatch</li>
<li>%s DynamoDB</li>
<li>%s RDS</li>
<li>%s ElastiCache</li>
</ul>
<br>
<br>
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/150px-Amazon_Web_Services_Logo.svg.png" alt="AWS" />`,
		embedded.NewImg(embedded.ImgPathS3, "S3").HTML,
		embedded.NewImg(embedded.ImgPathEC2, "Batch").HTML,
		embedded.NewImg(embedded.ImgPathCloudWatch, "CloudWatch").HTML,
		embedded.NewImg(embedded.ImgPathDynamoDB, "DynamoDB").HTML,
		embedded.NewImg(embedded.ImgPathRDS, "RDS").HTML,
		embedded.NewImg(embedded.ImgPathElastiCache, "ElastiCache").HTML)

	// A simple example of how the Data field can contain anything you want to send to the templates
	// even though you wouldn't normally send markup like this
	page.Data = AboutData{
		ShowCacheWarning: false,
		QualificationTabs: []AboutTab{
			{
				Title: "6+ Years of Experience",
				Body:  template.HTML(`I have been programming professionally since <strong>June 2017</strong>, though I also completed <strong>3 successful internships</strong> throughout my time at <strong>Purdue University</strong>, one of which was Remote during a Fall semester alongside classes.`),
			},
			{
				Title: "Education",
				Body:  template.HTML(`I obtained my <strong>Bachelor's of Science</strong> in <strong>Computer Science</strong> from <strong><a href="https://www.purdue.edu/">Purdue University</a></strong> in <strong>May of 2017</strong>.<br><br><br><img src="https://marcom.purdue.edu/app/uploads/2020/01/1_OurNewLogo.png" alt="Purdue logo" />`),
			},
			{
				Title: "Soft Skills",
				Body:  template.HTML(`I take pride in my ability to <strong>communicate technical topics</strong> to non-technical audiences in a way that is <strong>easily understandable</strong>. I am able to <strong>identify and describe</strong> both technical and non-technical <strong>requirements</strong> for projects. I have a strong preference to follow <strong>well-defined standards and practices</strong> if they exist and define them when they do not.`),
			},
		},
		TechnologyTabs: []AboutTab{
			{
				Title: "Go",
				Body:  template.HTML(fmt.Sprintf(`The <a href="https://go.dev/">"Go"</a> programming language. I use this language daily in my career and it is my "Go"-to choice of languages for side projects.<br><br>%s`, embedded.NewImg(embedded.ImgPathGopher, "Gopher").HTML)),
			},
			{
				Title: "AWS",
				Body:  template.HTML(awsServices),
			},
			{
				Title: "Kubernetes",
				Body:  template.HTML(`A majority of the services I develop professionally are deployed to <a href="https://kubernetes.io/">Kubernetes</a>. This includes everything from mappings to deployments and many things in between.`),
			},
			{
				Title: "Docker",
				Body:  template.HTML(`Alongside Kubernetes, <a href="https://www.docker.com/">Docker</a> is my technology of choice for containerization of services. I am experienced with creating Dockerfiles from scratch, using multi-stage builds, and also docker-compose.`),
			},
			{
				Title: "DataDog",
				Body:  template.HTML(`While I'm no expert SRE, I have a wide breadth of knowledge of how to effectively utilize <a href="https://www.datadoghq.com/">Datadog</a> for monitoring and observability.`),
			},
		},
		InterestTabs: []AboutTab{
			{
				Title: "Backend Development",
				Body:  template.HTML(`Aside from this website (most of which is written in Go anyways), all of my experience is with Backend Development. This is where I prefer to focus my time and effort as well.`),
			},
			{
				Title: "RESTful APIs",
				Body:  template.HTML(`I enjoy designing and implementing <a href="https://restfulapi.net/">RESTful APIs</a> to solve business problems for CRUD-based operations to data analysis pipelines.`),
			},
			{
				Title: "Observability",
				Body:  template.HTML(`Increasing the <a href="https://medium.com/linkapi-solutions/what-is-the-difference-between-api-observability-and-api-monitoring-2b29545a3b06">Observability</a> of services is something I make an effort to continuously learn about and improve for both existing and new services.`),
			},
			{
				Title: "Documentation",
				Body:  template.HTML(`Producing and maintaining high-quality documentation, especially for services where documentation is lacking, is a passion of mine. I belive that having good documentation is extremely important for the maintainability of services.`),
			},
			{
				Title: "Space Industry",
				Body:  template.HTML(`I have yet to do anything professional related to the Space industry, however space, astrophysics, telescopes, and all related things are huge interests of mine and I would be delighted to work on a project in that industry. <img src="https://www.nasa.gov/sites/default/files/styles/full_width_feature/public/thumbnails/image/28045752710_6a9cca2c72_k_0.jpg" alt="HST" />`),
			},
		},
	}

	return c.RenderPage(ctx, page)
}
