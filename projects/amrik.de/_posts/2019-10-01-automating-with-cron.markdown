---
layout: post
title:  "How to automate the boring stuff"
subtitle: "aka, cron 101"
date: 2019-10-01 00:00:00
categories: [tutorials]
---
# The Naive Solution
The solution for automating repetitive tasks that I think a lot of people, myself included, begin with is to write some code that will loop forever, but inside that loop just wait the amount of time that you think is appropriate. This solution seems like it could work, and I suppose it could if you really were set on doing it this way. But this solution brings along with it a plethora of issues, for example having to keep the application running forever, just hogging space on the stack, not to mention it fails to address problems like only executing on specific days or between specific hours, implementing these conditions could get very annoying very quickly, and if you wanted to change when the applications runs, it would be another tedious process yet again.

# A Better Solution
Cron! I can't believe I haven't heard about cron sooner. I've been using Linux systems as part of my University Course for over 3 years now and I've had to buy my own server and run into my own real world problems before I even knew it existed! So, hopefully there's at least one other hobbyist out there who could benefit from this!

## But What is cron?
Cron is a daemon that will execute commands based on a schedule, each user can have their own file that contains their specific schedules and commands for cron to run and these specific files are called crontabs.

## How does it work?
Almost all Linux distros come with cron by default and it's likely already running commands for you without you knowing it, if you don't have cron installed then you're probably a more advanced user of Linux and probably should be getting your advice elsewhere, if you're on windows, don't worry I've got a little bit at the end for you.

Okay so you shouldn't need to install cron but you still need to know how to make crontabs, they look a little bit confusing to begin with, but we'll get through it.

So to begin with you need to edit your crontabs, to do this simply run the command crontabs with the -e flag, as we want to edit ours.

So let's say we run `crontab -e` and add the line `* * * * * cmd` to the text file, what exactly have we done here?

The following diagram should explain the syntax of a crontab.

{% highlight bash %}
# .---------------- minute (0 - 59)
# |  .------------- hour (0 - 23)
# |  |  .---------- day of month (1 - 31)
# |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
# |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7)
# |  |  |  |  |
# *  *  *  *  *   cmd
{% endhighlight %}

So the \*'s are used to define when the a specific command should be ran, in this specific example it's quite boring as we're going to run the script "cmd" every single minute, as the \*'s stand for any value, meaning they'll run at every change of minute.

## A more realistic example

So let's say you're sick of backing up your work and even though you've created a script that makes it somewhat easier, you still wish the script would run automatically every day, preferably when it's likely not to disturb you.

Perfect, we can do exactly that with cron, let's add the line to our crontab `0 2 * * 1-5 backupScript.sh`
This script will now run at 0 minutes past 2, every day of the month, every month, every weekday, according to our diagram above. I'm sure there are many more creative solutions and problems you can solve using Cron, but this is one that I think every single user have running.


# Windows People!
Don't despair Windows people, I use Windows too! In windows there's something else called the "Task Scheduler" or schtasks sometimes, it has the same function as cron I believe. I'm not going to cover it here in depth here, this post has already dragged on long enough, but feel free to have a look at Microsoft's documentation about it [here](https://docs.microsoft.com/en-us/windows/win32/taskschd/schtasks).
