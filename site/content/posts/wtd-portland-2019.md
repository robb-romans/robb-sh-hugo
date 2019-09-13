---
title: "Highlights from the 2019 Write the Docs Conference in Portland"
date: 2019-07-02T10:56:59-05:00
draft: false
featuredImg: "/images/posts/2019/07/wtd-1.jpg"
tags:
  - docs
---

## Introduction

I was fortunate to be able to attend Write the Docs Portland this year for the second time. I always get a ton of value from the conference sessions and the fellowship with other writers. Here are some notes and links that I hope you find helpful.

### Common themes

- sense of community and overcoming isolation
- learning how others solve problems
- industry best practices
- latest research and solutions
- the surprising value of lightning talks

## Writing Day notes

### Writing Day tables

On Writing Day, there were a multitude of tables set up to discuss various topics. Here are the tables that I spent some time at.

- Tensorflow version 2
- Microsoft
- Netlify CMS docs; making forked repo CMS work
- Cockroachdb
- Mike Lewis - GitLab docs as code; GitLab Pages
- Dan Allen - asciidoctor open source
- Marty Avidon - Wikimedia
- ST Cohen - IA
- Anita - drawing scenes

### Antora / AsciiDoc tables

I spent a lot of time here. OpenDevise and Antora are doing some amazing things. Here are some links to get you started.

- <https://github.com/DocOps/wtd2019/blob/master/README.adoc>
- <https://github.com/tomjoht/documentation-theme-jekyll>
- <https://github.com/DocOps/liquidoc-cmf>
- <https://linuxacademy.com/devops/training/course/name/dev-ops-monitoring-deep-dive>

<br/><br/>

## Presentation notes

### [Loving the CLI](https://www.writethedocs.org/videos/portland/2019/how-i-learned-to-stop-worrying-and-love-the-command-line-mike-jang/)

Mike Jang

{{< youtube YfM220sNVp8 >}}

<br/><br/>


### [Documenting for open source software](https://www.writethedocs.org/videos/portland/2019/documenting-for-open-source-shannon-crabill/)

Shannon Crabill

{{< youtube i22duorDVt0 >}}

<https://shannoncrabil.com/writethedocs>

- incomplete or outdated docs are a pervasive problem, observed by 93% of respondents <https://opensourcesurvey.com>
- avoid assuming the technical proficiency of your readers
- put the effort into your README file
- include a code of conduct
- solve for duplicate pull requests - create pull request and issue templates
- use GH tags intelligently
- have a good changelog
- leverage / link to other helpful docs
- anyone can contribute to open source

<br/><br/>

### [Writer? editor? teacher?](https://www.writethedocs.org/videos/portland/2019/writer-editor-teacher-kathleen-juell/)

Kathleen Juell at Digital Ocean

{{< youtube JDrb-meCtHQ >}}

- create docs with the right level of access
- empower users to ownership of ideas and methodologies
- empower users to create cool stuff with technology

What can we learn from related fields?

Statement of teaching philosophy

- your conception of teaching and learning
- a description of how you teach
- a justification for why you teach that way

Teaching in your intro (e.g. tutorial)

- provide users with necessary context and definitions
- frame the why
- explain the how

provide context and background

- define
- justify
- outline

complex content: see the book "Understanding by Design" re: K-12

- explain general principles where possible
- reference different contexts and scenarios where possible
- gauge your audience
- say why

Editorial templates

- packaged, reproducible suggestions
- prewriting
  - what is this tutorial about?
  - why should the reader learn this?
  - what will they have accomplished?
- writing
  - discussing large code blocks and commands
  - tell the reader what code will do, first
  - after, explain some of its key concepts

Peer editing as teaching

- explain your thinking
- treat grammar like a high-level topic
- ask questions

Write like a teacher

- clarify your why
- have an understanding goal
- explain your own thinking to yourself

<br/><br/>

### [How to edit other people's content without pissing them off](https://www.writethedocs.org/videos/portland/2019/how-to-edit-other-peoples-content-without-pissing-them-off-ingrid-towey/)

Ingrid Towey at Red Hat

{{< youtube 7iWUSetbaos >}}

We're all on the same side

- customer side
- don't argue
- some things are true for all customers
- people form a lasting opinion of your site in 50 milliseconds (1/20 sec)
- you have at most 10 seconds before users leave your website
- you have up to 60 secs if your site is already trusted
- they read 28% of page content
- net: customers scan, not read

Edit, not edict

- not a tennis match
- no passion equals the passion to alter someone else's content
- make your edits reversible
- she edits for 124 writers

Explain why

Resources

- Flesch-Kincaid Grade Levels
- readable.io
- Hemingway editor
- Microsoft Word
- Acrolinx
- IBM style guide
- MicroSoft style guide
- Google style guide
- minimalism
- DITA guidelines for modular docs
- usability testing
- nielsen norman group articles

Get help

- peer review
- group editing

Be kind

- Toastmasters
- positive, negative, positive feedback sandwich
- have a meeting leader
- it's all about the relationships
- assume positive intent

<br/><br/>

### [Making friends of the docs](https://www.writethedocs.org/videos/portland/2019/any-friend-of-the-docs-is-a-friend-of-mine-cultivating-a-community-of-documentation-advocates-heather-stenson/)

Heather Stenson

{{< youtube lZhIQNlGY4s >}}

How to make friends of the docs

- team leads are a great resource. face to face time at least quarterly.
- share product feedback - good or bad!
- tie documentation to the bottom line
- bribery (stickers, shirts, food, happy hours)
- use the processes that are already in place (find and leverage)
- create a docs-specific global channel
- create a docs issue tracker
- create 'friends-of-the-docs' badge and be able to communicate with all at once

Obstacles

- silos: open up your own silo, first
- become a bridge between silos
- widespread corporate docs obliviousness
  - have docs team be very vocal and visible
  - reach out proactively to other teams
  - every thing you do must demonstrate that "we are a company that values documentation"

Everyone is too busy

- spend time now to save time later
  - support writes a tutorial that saves a bunch of time handling customers opening tickets
- helping -> response -> satisfaction loop

Maintaining this is an ongoing process

- model behaviors that you want to see
- loosen your grip a little
- set boundaries and parameters, guard rails
- help the helpers
- keep things fun and FRIENDly
- swag, give public appreciation
- appreciate people
- stay friends when the team relationships change

Don't:

- be a jerk
- take people for granted

<br/><br/>

### [Defying the Status Quo](https://www.writethedocs.org/videos/portland/2019/defying-the-status-quo-how-a-grassroots-effort-can-transform-an-organization-jodie-putrino/)

Jodie Putrino at F5 Networks

<https://clouddocs.f5.com>

{{< youtube 09z0VDR_bxs >}}

Process was:

- DITA
- reviews from annotated PDF
- waterfall
- ineffective review process

Defiance

> The process is not the thing. It's always worth asking, do we own the process or does the process own us?  --Jeff Bezos

Now:

- simple markup language
- GitHub PRs
- Agile process
- CI/CD, a.k.a. Docs as Code

See her 2017 presentation:
{{< youtube Mzu-c-FoOdw >}}

Influential sessions:

- [2015 WtD: How GitHub uses GitHub to document GitHub](https://youtu.be/s46m8H4BrrE)
- [2015 WtD: Documentation, Disrupted](https://youtu.be/EnB8GtPuauw)

Success factors

- helping others
- fierce advocacy
- don't take it personally
- be patient
- have rabbis

You are the best advocate for your project.

> Status quo is the biggest threat for companies, culture, personal growth - Francois Loc... F5

<br/><br/>

### [Lessons learned in a year of docs-driven development](https://www.writethedocs.org/videos/portland/2019/lessons-learned-in-a-year-of-docs-driven-development-jessica-parsons/)

Jessica Parsons at Netlify

{{< youtube WDYQoZ-QDRM >}}

Write the docs first! Build to match.

You can include user stories or specs, but there's more to it than that.

This is not waterfall all over again.

Iterate.

How were they successful at this?

- strong leadership support
- docs valued in the org
- generally good communicators

Lessons

- good intentions are not enough

Tips on making this happen

- provide scaffolding (term from educational psychology)
- think beyond the written word: diagrams, wire frames, video
- designate a leader to drive success (perhaps rotate)
- make sure your work counts (see Tanya Reilly's talk on Glue work)
- Identify issues when investment is low
- devs adjust plans as they write
- writers spot issues that spark conversations
- leverage team expertise
- add accountability - reviewers and time lines
- notify people about changes
- look at Notion tool (doesn't solve all problems)

Good docs enable:

- reducing impostor syndrome
- parallel execution
  - example: front-end developers against API spec while API devs build the endpoints
- **support learning about features in advance**
- marketing can prepare for launch

Docs-driven doesn't guarantee "user-centered:"

- beware function-focused user stories
- get users involved sooner
- enlist customer-facing teammates

School is still in session.

> Docs or it didn't happen -> docs or it won't happen.

<br/><br/>

### [Show Me the Money: How to Get Your Docs the Love and Support They Deserve](https://www.writethedocs.org/videos/portland/2019/show-me-the-money-how-to-get-your-docs-the-love-and-support-they-deserve-matt-reiner/)

Matt Reiner

{{< youtube 8ZPUOwBIi3g >}}

- SWOT
- map your goals to the long-term plan of the business
- elevator pitch, the what and why
- internal marketing. the seasons when this is important in your company.
- find friends in sales and marketing
- sales controls the perception of "what the product does"
- support is a huge focus these days
- marketing, sales, and support
- competitor analysis

Build the business case

- need metrics and data
- overcome exec's gut feeling that docs are a waste of money
- support deflection rates
- your friends have data that can help you
- be helpful and practical (you don't have to know everything)
- influence others - don't bulldoze
-  assemble your plan. draft. collaborate. present to friends. iterate.

Write the product plan - whole enchilada with details
Financial Plan - get help

<br/><br/>

### [Just Add Data: Make it easier to prioritize your documentation](https://www.writethedocs.org/conf/portland/2019/speakers/#speaker-portland-2019-sarah-moir)

Sarah Moir

{{< youtube 5kTWjB28TDI >}}

Splunk - big data analysis

- prioritizing docs is hard
- use data to build confidence in decisions
- challenge or validate assumptions

<br/><br/>

### [Drawing in technical communication](https://www.writethedocs.org/videos/portland/2019/draw-the-docs-alicja-raszkowska/)

Alicja Raszkowska

{{< youtube Zvys6tUmEzg >}}

See: Mermaid software

<br/><br/>

## Lessons learned at the conference

- The value of our community
  - writers sometimes find themselves isolated
  - connectedness
  - solidarity
  - esprit de corp
  - the richness of face to face conversations
  - the impact of social setting on conversations
- building relationships
  - the collaborative value of un-conferences
- the business value of content
  - expertly presented by Matt Reiner
- the Doc-Ops revolution
  - frictionless interaction with tools, frameworks, and platforms
  - increases productivity
  - make changes often without breaking systems
  - move with speed SRE concept: short lead time. reduce feedback latency.
  - developer experience
  - automated doc testing
- the increasing use of Git and GitHub, GitLab, etc.
- the value of learning the latest documentation best practices

## [A link to Aaron Thayer's detailed conference notes](https://github.com/a-thay/WTD-2019/)
