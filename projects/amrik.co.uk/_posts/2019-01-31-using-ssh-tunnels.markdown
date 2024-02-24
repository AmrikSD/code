---
layout: post
title:  "Port forwarding with SSH"
subtitle: "aka, how to access your own media server at work or university"
date:  2019-01-31 11:34:01
categories: [tutorials]
---
# Motivation

Sometimes pesky firewalls are in the way of you accessing a server, SSH tunnelling can come to our aid in this case and let us access them regardless of what some sysadmin wants...

Apparently SSH tunnels are also used to add encryption to legacy applications too, but I'm pretty sure I shouldn't be encouraging that, and regardless it's certainly not interesting to me at least, instead we're going to talk about getting around those pesky firewalls instead of janky bolted on encryption, specifically we're going to be talking about my own scenario.

## My problem

I love to use plex, I've been using it for a few months now and I love the idea of being able to store media on my own server and distribute it around the house or even stream it to my phone like a pseudo Spotify or Netflix. It's quickly become my main source of media consumption outside of YouTube and I use it daily, there are very few complaints to be had. If you're not aware what I'm talking about, just search plex - if you have your own server that's not doing much or even a spare old computer I'm sure you could do much worse with it than putting plex on it and using it to kickstart your media centre!

So with that being said, my university however, isn't a fan of plex. In fact, their network blocks connections to the port used by plex completely - I'm not sure if they have a personal vendetta against plex or they simply block all uncommon ports, but in any case it's annoying, I just want to get some work done with some music.

# The Solution

I simply use an SSH tunnel, of course this solution only works if you can ssh into the machine that holds your media, but I'm going to assume you can figure that part out.

{% highlight bash %}
foo@bar:~$ ssh -L localport:localhost:serverport user@server.com
{% endhighlight %}

This simply opens a connection to your server as usual, however, it then pushes every request from localhost:localport to server.com:serverport, in my particular case I want to use plex from my laptop, so I'll type

{% highlight bash %}
foo@bar:~$ ssh -L 80:localhost:32400 user@myserver.com
{% endhighlight %}

Now I just hop to my browser, navigate to localhost and just that easy, I'm greeted with plex and I can get back to being productive again.

# What does that actually do though?

What we did was SSH port forwarding, or SSH tunnelling as most people call it, specifically we did something called Local Forwarding which forwards a port from our client (Local) machine to a server machine (Remote). We have essentially created an SSH client that will listen for connections on a specific port and when it receives a request it will tunnel the connection to an SSH server, which could be on a different port and potentially even an entirely different machine than the SSH server.

If you're still confused, feel free to take a look at [ssh.com](https://www.ssh.com/ssh/tunneling/example#sec-What-Is-SSH-Port-Forwarding-aka-SSH-Tunneling) on which there is a written article all about SSH tunnelling that goes into more depth than I feel is necessary here.
