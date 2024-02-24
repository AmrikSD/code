---
layout: post
title: "Amrik.de - a personal blog"
subtitle: "yep, this very website."
date: 2021-03-03 17:37:23
categories: [projects, meta]
---

# What is Amrik.de?

Amrik.de is my space on the internet to share things I think are interesting or useful to others.

If I stumble upon or [create](https://amrik.de/categories/projects/) something useful to me, I think writing about it is a good way to share it so that others can find some usefulness out of it too.
It's not all selfless though, I believe that if I take the time to write about something I'll be forced develop at least a slightly deeper understanding of it too, which is always a good thing.

In the rest of this post, I'll try to explain how the website works, how I make posts, and how everything is set up for me to focus most on writing and least on maintaining the website.

# How It Works

## The Infrastructure

![Amrik.DE](https://amrik.de/assets/posts/amrik-dot-de/Architecture.png)

Amrik.de is a pretty simple website, there is no backend needed as all I need to do is show some text and images on the screen and have it look pretty. That being said, there are various technologies being leveraged under the hood, most of which being provided by AWS.

The diagram above provides an overview of all the different parts that go into Amrik.de, I'll cover them each in detail below. If you're less interested in the specifics, feel free to skip to the other sections.

### Route 53

[Route53](https://aws.amazon.com/route53/) is a Domain Name System service.

In the most high level view, all it is doing is helping to translate the domain name ["Amrik.de"](https://amrik.de/meta/2019/09/10/why-dot-de.html) into an IP addresses. Specifically for this project, when someone types "Amrik.de" into their browser Route53 will point them to a CloudFront distribution.

### CloudFront

[CloudFront](https://aws.amazon.com/cloudfront/) is a "Content Delivery Network". For this website it's main function is to deliver data (HTML, CSS, JS) to users from an S3 bucket.

Additionally, by using CloudFront's large edge networking locations, the website can be cached relatively closely to users meaning that the website will load faster than if they were to get the content from S3 directly.

CloudFront was also chosen because of some limitations with S3, specifically that as of now S3 doesn't support HTTPS. I'm of the opinion that even when hosting a project such as a blog with nothing at all sensitive, there is still no reason to not use HTTPS. It's completely free to get certificates these days and if _nothing_ else, it definitely makes your website appear more professional and trustworthy, at least in my eyes.

### S3

[S3](https://aws.amazon.com/s3/) is a service provided by AWS that stores objects. After the content of Amrik.de is generated into HTML etc., it is pushed to S3 and from here it can be accessed only from CloudFront.

### GitHub

GitHub isn't actually part of the website, but it does play a huge part in how the website is maintained. All of the code for the website is stored in [this repository](https://github.com/AmrikSD/Amrik.de) on GitHub. The history of all changes to the source code of the website, posts, and media are all stored on Git. To automate changes to the website, every time I push to the main branch of this repository I use [Github Actions](https://github.com/features/actions) in order to make sure that the changes are pushed to the live version of the website, without needing to do _anything_ manually.

# How I Make Posts

The way I make posts is pretty simple, I don't have to meticulously create each page and apply styling, instead I am leveraging a site generator called [Jekyll](https://jekyllrb.com/). It should be noted that I'm not particularly a massive fan or advocate of Jekyll, I actually think at this point is it quite antiquated and there have been various other site generators released, I just found Jekyll and for me it's been doing the job.

To begin writing a post, I just create a very simple markdown file and populate some metadata at the top such as the title/date/category of the post. The rest of the post is literally just markdown, there is actually no HTML/CSS/JS knowledge necessary in order to make posts.

After the post is finished, in markdown - I simply run `jekyll build`, Jekyll creates a folder called `_site` with all the HTML, CSS & Images that are needed to make the website look as it does.

In reality, I don't actually even build the website locally, I have GitHub actions run everything from building to pushing to s3 etc, I can now focus on writing posts and not deploying them, it's 100% automated.

# How Much It Costs

After the domain name, $0.60 per month.

This can change, of course, but as of now the only part of this website that is incurring _any_ cost is the [Route 53](https://aws.amazon.com/route53/pricing/) Hosted Zone. This is $0.5 per month, with tax added the total is $0.6.

Of course I'm aware that after I expend any free tiers this cost will rise, but even once this happens this website is _very_ cheap to run all things considered the availability and performance given by CloudFront particularly.

# Do I Recommend this Approach?

Yes. After hosting locally for several years, and having to manually deploy posts and deal with any networking changes manually, having all of this work outsourced to GitHub Actions and AWS has meant I am free to enjoy writing more instead of maintaining technologies website.

I think this was a good project to get my feet wet with AWS or "The Cloud!" in general.

I had a lot of concerns around accidental costs and running up a huge bill by accident with AWS but after setting up billing alerts and being careful & seeking advice as to which services should be used for the project, I can say the experience and end product was well worth it.

If you are particularly worried about costs, I would recommend looking at the [The GCP Free Tier](https://cloud.google.com/free/docs/gcp-free-tier). As I understand it, you are given some credit to use on their service - after you use up this credit, instead of being charged unexpectedly, your services are instead turned off - I think for those who are working on personal projects or just looking to learn about these cloud providers, it may be the way to go.
