---
title: "Adding Hugo Blog Comments Using Utteranc.es"
date: 2019-04-29T18:55:05-05:00
draft: true
tags:
  - hugo
  - serverless
---

There's some debate about whether adding a comment mechanism to a low-traffic blog like this one is worth the headache. In my case, I'm spending the time for three reasons: to learn about the best options available, to provide a dedicated way to interact with readers, and because this is relevant to my [day job](https://docs.rackspace.com/).

The simplest comment system to use in Hugo is [Disqus](https://disqus.com/). Hugo has built-in support for it, and most themes make it easy to display comments at the end of your posts by only setting a variable or two. However, there are some privacy concerns about Disqus, which you can read about [here](https://discourse.gohugo.io/t/alternative-to-disqus-needed-more-than-ever/5516/). If those don't concern you, Disqus is the shortest path to feedback success.

## Alternatives to Disqus

The Hugo docs list some alternatives to Disqus [here](https://gohugo.io/content-management/comments/#comments-alternatives). Since this site is serverless by using [Netlify](https://www.netlify.com/) to build and deploy, I didn't want to complicate my life with a comment system that requires me to maintain infrastructure like a VM or a Kubernetes cluster. That means I could choose to pay for comments as a service, or find a serverless solution. I don't have any problem paying for good software; but again, this is a hobby blog and mostly an experiment for learning. That raises the threshold for return on investment which eliminates most of the for-profit solutions. Some alternatives have no or low buy-in costs, but fees scale up with page hits or comment count. I was hesitant to choose those, since there is a (remote) possibility one of these posts might get its 15 minutes of fame on the web, and I'd faint when the bill came.

## Choosing Utterances

[Utterances](https://utteranc.es/index.html) is a small, open source TypeScript comments widget that stores comments as GitHub issues. Built by Jeremy Danyow, it's a lighter-weight version of the comment system used with https://docs.microsoft.com. Good enough for me. ![https://docs.microsoft.com/en-us/azure/virtual-machines/linux/tutorial-manage-vm](/images/posts/2019/04/utterances-microsoft-docs-feedback.png)

## Implementing Utteranc.es

Adding Utteranc.es was pretty simple.

Run the Utterances [configuration tool](https://utteranc.es/#configuration) to generate the script configuration.
![Configuration tool output](/images/posts/2019/04/utterances-web-setup-results.png)

Create `site/content/layouts/partials/comments.html` containing the output from the tool. For example:

```bash
$ mkdir -p site/content/layouts/partials
$ cat << EOF > site/content/layouts/partials/comments.html
<script src="https://utteranc.es/client.js"
        repo="robb-romans/robb-sh-comments"
        issue-term="pathname"
        label="utteranc-es"
        theme="photon-dark"
        crossorigin="anonymous"
        async>
</script>
EOF
```

Grant the [Utterances app](https://github.com/apps/utterances) access to your GitHub account. I choose only the repo used to hold comments. Every person who wants to comment on an Utterances site will have to do this once.

## What comments look like

### On a post
![Comments on mobile browser](/images/posts/2019/04/utterances-mobile-web-comments.jpg)

### On GitHub
![Comments on GitHub](/images/posts/2019/04/utterances-github-comments.png)

## When you've had enough feedback on a post

To freeze comments associated with a post, open the GitHub issue associated with the post and click the 🔒Lock conversation button on the bottom right. ![🔒 Lock conversation](/images/posts/2019/04/utterance-github-lock-conversation.png)

You can delete a comment in the GitHub issue by clicking the ellipsis icon by that comment and selecting `Delete`.

## Utteranc.es limitations

* Readers must have a GitHub account to comment. Utteranc.es supports no other login or authentication mechanisms. Fortunately, it's easy to get a free account.
* For storing comments, Utteranc.es supports only GitHub. GitLab, BitBucket, and so on are not supported. Since this site runs on GitLab, I created a separate repo in GitHub to hold the comments. Cats and dogs together.
* Utteranc.es [uses third-party cookies](https://github.com/utterance/utterances/issues/123). If you run a privacy-focused browser or browser extensions, this can require whitelisting Utteranc.es servers to see the comments.
  * `[*.]api.utteranc.es`
  * `https://[*.]utteranc.es:443`

## Try it out

Please try out the comment feature below this post and tell me what you think about this article. Suggestions welcome!

## References

* [https://github.com/utterance/utterances](https://github.com/utterance/utterances)
* [https://gitlab.com/robb-romans/robb-sh-victor-hugo/issues/4](https://gitlab.com/robb-romans/robb-sh-victor-hugo/issues/4)
* [https://github.com/robb-romans/robb-sh-comments/issues/1](https://github.com/robb-romans/robb-sh-comments/issues/1)
