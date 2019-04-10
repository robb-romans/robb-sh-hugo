# README

[![Netlify Status](https://api.netlify.com/api/v1/badges/15a99f94-b4f9-46a7-bf2c-122bee83114a/deploy-status)](https://app.netlify.com/sites/xenodochial-hoover-28928a/deploys)

This code publishes my personal website at <https://robb.sh/>.

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

```bash
|--site                // hugo builds everything in here
|  |--content          // Pages and collections - ask if you need extra pages
|  |--data             // YAML data files with any data for use in examples
|  |--layouts          // This is where all templates go
|  |  |--partials      // This is where includes live
|  |  |--index.html    // The index page
|  |--static           // Files in here ends up in the public folder
|--src                 // Files that will pass through the asset pipeline
|  |--css              // Webpack will bundle imported css seperately
|  |--fonts            // Fonts get moved to a flattened directory inside dist
|  |--index.js         // index.js is the webpack entry for your css & js assets
```
