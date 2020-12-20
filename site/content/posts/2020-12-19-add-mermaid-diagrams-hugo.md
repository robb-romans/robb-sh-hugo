---
title: "Build and Display Mermaid Diagrams in Hugo"
date: 2020-12-19T12:42:28-06:00
draft: true
tags:
  - hugo
---

You can use plain text image markup formats stored in Git to create diagrams. For some background on
why, check out this talk:

{{< youtube 3i-C7qbRGGQ >}}

<br/><br/>

Here's a quick tutorial on how to add support for building
[mermaid.js](https://mermaid-js.github.io/mermaid/#/) diagrams in Hugo.

First, add a new shortcode in `layouts/shortcodes/mermaid.html`.

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

Next, add some mermaid.js markup to your post. Here's a simple example.

```go
{{</* mermaid align="left" theme="neutral" */>}}
pie
    title Animal Party
    "Dogs" : 50
    "Cats" : 35
    "Goats" : 10
    "Lemurs" : 5
{{</* /mermaid */>}}
```

After running your Hugo build process, that code renders to:

{{<mermaid align="left" theme="neutral">}}
pie
    title Animal Party
    "Dogs" : 50
    "Cats" : 35
    "Goats" : 10
    "Lemurs" : 5
{{</mermaid>}}
