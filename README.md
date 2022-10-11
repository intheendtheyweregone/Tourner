# Tourner[^1]
A Go-powered ShareX "round-robin"[^2] uploader

[^1]: "Tourner" is a French Word that means to turn or rotate.
[^2]: This is a simple explanation, but this isn't really round-robin anything, and could be more simply described as a rotation system..

<br>
<img src="https://user-images.githubusercontent.com/114883905/195212649-a452f58e-cef9-4270-8428-ab8e593c3700.png" height="150" />

## How it works?
Tourner proxies requests from a local web server to various LoliSafe (now ChibiSafe) instances, selecting, at random, an instance each time a request to upload a file is made to the "/upload" endpint.

Currenly Tourner only supports LoliSafe, which I believe is referred to as version 3.x of ChibiSafe, but as I continue to work on this project 4.X will be supported, as well as other file uploaders in general.


## It's Purpose?
I really didn't like the idea of keeping all of my funny meme screenshots on one shitty server in Tunisia, and liked the idea of having screenshots spread out across various different hosts as using only one can definitely make it easier for one instance owner to track everything that you upload, with this they'll get bits and pieces, but likely not enough to piece together a full picture. In the future I also plan to add the ability to specify a proxy/proxies that can be applied to further improve your privacy, as ShareX on Windows does not currently support SOCKS Proxies in the first place.


## Install

Download the latest executable release or grab the source code and run ``` go build ``` with proper settings for your operating system.

Now that you have run the executable you should get a message that basically says that the server is starting, now you should open ShareX and head to your ```Custom Destinations``` settings.

![image](https://user-images.githubusercontent.com/114883905/195211814-626c2dc1-8403-42f5-9bdc-6c1cf5ef8a14.png)

You should then copy these settings (adjust them if there are any issues, but as the time of writing this is how it's done) 

Set your image and optionally file uploader to whatever you named the Uploader and then press test

![image](https://user-images.githubusercontent.com/114883905/195211980-052e2692-0e9f-4731-ae40-c1f7bce203a3.png)

IF everything worked correctly then you should see a link from one of the LoliSafe Instances like this:

![image](https://user-images.githubusercontent.com/114883905/195212084-1f528374-c715-48c6-9f5f-6f866dfa6b1f.png)

And that's it! You've successfully setup Tourner and can now use it normally.


## What's next?

I plan on adding a host of different features, as well as cleaning up the code any fixing other issues and bubblewrapping edge cases as I become more familiar with the Go Language, this Repository is purely for my own tracking of my progress, but if anyone comes across this and would like to use it or change it, feel free.

