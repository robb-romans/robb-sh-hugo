---
title: "Build and Display Mermaid Diagrams in Hugo"
slug: how-to-use-mermaid-diagrams-in-hugo
date: 2020-12-19T12:42:28-06:00
draft: false
tags:
  - hugo
---

## Intro

You can use plain text markup formats like
[mermaid.js](https://mermaid-js.github.io/mermaid/#/) to dynamically create
graphical diagrams within your Markdown files, instead of using a separate
diagram tool to create binary image files. For some background on why you might
want to do this, check out this talk:

{{< youtube 3i-C7qbRGGQ >}}

<br><br>

## Add the shortcode

Here's a quick tutorial on how to add support for building mermaid.js diagrams
in Hugo.

First, add a new shortcode in `layouts/shortcodes/mermaid.html`. Note that the
package version is not hardcoded in this code, so the CDN always fetches the
latest version. Also note that you can pass parameters to mermaid.js in the
shortcode. Here we pass the theme and alignment, or set a default if the
parameters are not passed.

```javascript
<script async type="application/javascript" src="https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js">
  var config = {
    startOnLoad:true,
    theme:'{{ if .Get "theme" }}{{ .Get "theme" }}{{ else }}dark{{ end }}',
    align:'{{ if .Get "align" }}{{ .Get "align" }}{{ else }}center{{ end }}'
  };
  mermaid.initialize(config);
</script>

<div class="mermaid">
  {{.Inner}}
</div>
```

## Create a diagram

Next, add some mermaid.js markup to your post. Here's a simple example.

```go
{{</* mermaid align="left" theme="neutral" */>}}
pie
    title French Words I Know
    "Merde" : 50
    "Oui" : 35
    "Alors" : 10
    "Non" : 5
{{</* /mermaid */>}}
```

After running your Hugo build process, that code renders in your page to:

{{<mermaid align="left" theme="neutral">}}
pie
    title French Words I Know
    "Merde" : 50
    "Oui" : 35
    "Alors" : 10
    "Non" : 5
{{</mermaid>}}
