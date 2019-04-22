# README

[![Netlify Status](https://api.netlify.com/api/v1/badges/15a99f94-b4f9-46a7-bf2c-122bee83114a/deploy-status)](https://app.netlify.com/sites/robb-sh/deploys)
[![License: CC BY-NC-SA 4.0](https://img.shields.io/badge/License-CC%20BY--NC--SA%204.0-success.svg)](https://creativecommons.org/licenses/by-nc-sa/4.0/)

## This code publishes my personal website at <https://robb.sh/>

## Thank you to the following projects for the tools used to build this site

* [GitLab](<https://gitlab.com/>) DevOps lifecycle tools
* [Netlify](<https://www.netlify.com/>) web project automation
* [LetsEncrypt](<https://letsencrypt.org/>) certificate authority
* [Victor Hugo](<https://github.com/netlify-templates/victor-hugo/>) deployment templates
* [Hugo](<https://gohugo.io/>) static site generator
* [Hermit](https://github.com/Track3/hermit) smooth and clean Hugo theme
* [Webpack](<https://webpack.js.org/>) asset pipeline
* [PostCSS](<https://postcss.org/>) JavaScript CSS transform tool
* [Babel](<https://babeljs.io/>) JavaScript compiler
* [RenovateBot](<https://renovatebot.com/>) automated dependency updates

## File structure

```text
|--site                // Hugo builds everything here
|  |--archetypes       // Defaults for content types
|  |--content          // Pages and collections
|  |--data             // YAML data files with any data for use in examples
|  |--layouts          // Layout templates
|  |  |--partials      // Includes
|  |  |--index.html    // The index page
|  |--resources        // Generated assets
|  |--static           // Files here end up in the public folder
|  |--themes           // Themes added here as Git submodules
|--src                 // Files that will pass through the asset pipeline
|  |--css              // Webpack will bundle imported CSS separately
|  |--fonts            // Fonts get moved to a flattened directory inside /dist
|  |--index.js         // index.js is the webpack entry for your CSS & JS assets
```

For more information, see [Hugo's Directory Structure Explained](https://www.jakewiesler.com/blog/hugo-directory-structure/).