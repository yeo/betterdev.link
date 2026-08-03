---
title: Hello, world
description: Better Dev now has a blog — longer-form writing to go with the weekly links.
date: 2026-07-28
draft: false
---

For years Better Dev has been a weekly list of links. Some topics deserve
more than a one-line blurb, so this blog is where the longer-form writing
will live.

## What to expect

Occasional posts on the craft of software engineering:

- Deep dives into tools and techniques from past issues
- Notes on running this newsletter and its infrastructure
- Things I learned that don't fit a link blurb

I like self-hosted, especially focus on small unit of software that we can
operate outself without relying on massive manage service. [> This whole
site is a small Go binary compiling markdown to static HTML, running on a
tiny k8s cluster. <]

## How posts are written

Posts are support a TOCs and a side notes so our reader can follow and link
to each section easiser

### Table of contents

Every `##` and `###` heading gets an anchor id, and a table of contents is
generated automatically at the top of the post when there are at least two
headings.

### Side notes

Wrap a remark in the note markers right where it belongs in the text, and
it becomes a side note in the right margin — like the one in the section
above. On narrow screens it turns into an inline note box instead.

```markdown
Weekly issues keep arriving as usual. [> The blog and the newsletter
share the same build, one binary produces both. <]
```

That's it. Subscribe to the [newsletter](/) if you haven't — the blog
will stay quiet and occasional.
