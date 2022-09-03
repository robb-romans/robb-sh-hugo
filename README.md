# README

[![Netlify Status](https://api.netlify.com/api/v1/badges/15a99f94-b4f9-46a7-bf2c-122bee83114a/deploy-status)](https://app.netlify.com/sites/robb-sh/deploys)
[![License: CC BY-NC-SA 4.0](https://img.shields.io/badge/License-CC%20BY--NC--SA%204.0-success.svg)](https://creativecommons.org/licenses/by-nc-sa/4.0/)
[![Accessibility](https://img.shields.io/badge/accessibility-passing-success?style=flat&logo=html5&logoColor=white)](https://wave.webaim.org/report#/https://robb.sh/)
[![PageSpeed](https://img.shields.io/badge/PageSpeed-98%20%2F%20100-success?style=flat&logo=google&logoColor=white)](https://developers.google.com/speed/pagespeed/insights/?url=https%3A%2F%2Frobb.sh%2F&tab=desktop)
[![Twitter](https://img.shields.io/twitter/follow/RobbRomans.svg?style=social")](https://twitter.com/robbromans)
[![Open in Visual Studio Code](https://img.shields.io/static/v1?logo=visualstudiocode&label=&message=Open%20in%20Visual%20Studio%20Code&labelColor=2c2c32&color=007acc&logoColor=007acc)](https://open.vscode.dev/robb-romans/robb-sh-hugo)

## This code publishes the personal website of Robb Romans at <https://robb.sh/>

## Thank you to the following projects for the tools used to build this site

- [GitHub](https://github.com/) software development platform
- [Netlify](https://www.netlify.com/) web project automation
- [LetsEncrypt](https://letsencrypt.org/) certificate authority
- [Victor Hugo](https://github.com/netlify-templates/victor-hugo/) OG deployment template
- [Hugo](https://gohugo.io/) static site generator
- [Hermit](https://github.com/Track3/hermit) smooth and clean Hugo theme
- [GoatCounter](https://www.goatcounter.com/) open source web analytics platform

## File structure

```text
|--site                // Hugo builds this folder
|  |--archetypes       // Defaults for types of new content
|  |--content          // Pages and collections
|  |--data             // YAML data files with data for use in examples
|  |--layouts          // Layout templates
|  |  |--partials      // Includes
|  |  |--index.html    // The index page
|  |--resources        // Generated assets
|  |--static           // Files here end up in the site/public folder (?)
|  |--themes           // Add themes here as Git submodules
```

For more information, see [Hugo's Directory Structure Explained](https://www.jakewiesler.com/blog/hugo-directory-structure/).
