---
title: "Updating webpack-merge to Version 5"
date: 2020-07-06T21:27:07-05:00
draft: false
aliases:
  - "/posts/2020-07-webpack-merge-v5/"
tags: 
  - webpack
---

This website uses Hugo and Webpack to build content. The dependency
[webpack-merge](https://github.com/survivejs/webpack-merge) package moved from
version 4 to 5.

```diff
diff --git a/package.json b/package.json
index 92409ec..b79763d 100644
--- a/package.json
+++ b/package.json
@@ -1,6 +1,6 @@
 {
     "webpack-dev-server": "^3.11.0",
-    "webpack-merge": "^4.2.2",
+    "webpack-merge": "^5.0.6",
     "whatwg-fetch": "^3.0.0"
   }
```

There are some [breaking
changes](https://github.com/survivejs/webpack-merge/blob/master/CHANGELOG.md#503--2020-07-06)
in webpack-merge version 5. If you encounter new build errors, the simple fix
for me was:

```diff
diff --git a/webpack.dev.js b/webpack.dev.js
index 15d82f2..83b7fc2 100644
--- a/webpack.dev.js
+++ b/webpack.dev.js
@@ -1,4 +1,4 @@
-const merge = require("webpack-merge");
+const {merge} = require("webpack-merge");
 const path = require("path");
 const {CleanWebpackPlugin} = require("clean-webpack-plugin");
 const MiniCssExtractPlugin = require("mini-css-extract-plugin");
```
