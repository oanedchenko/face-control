# Avatar image checker
The main goal of the project is to facilitate avatar image checking to remind people to set real picture. It's important for instance for corporate environment to let new people to get on board quickly.

_Inspired by https://github.com/zikes/chrisify_ :)

## Current functionality
~~Commandline util checking if the image contains any faces.~~
Commandline util which is slack bot. When it runs it checks slack workspace users and send messages to those users who do not have exactly 1 face detected on his/her avatar.

### Installation
* Install the OpenCV Developer package. On Ubuntu systems that's `sudo apt install libopencv-dev`
* `go get github.com/oanedchenko/face-control`
* `go get github.com/nlopes/slack`
* `cd $GOPATH/src/github.com/oanedchenko/face-control && go build`

### Usage
~~`./face-control < path/to/image.jpg`~~
`SLACK_TOKEN=your-slack-app-token ./face-control`

~The output will be like that:~
~`Face 0 detected: (199,235)-(648,684)`~

## Plans
~I'm going to create a slack bot to check if people use their real picture as an avatar image and to annoy them with remainders to do so.~
To run it in docker container as a demon, to make it work by schedule.
