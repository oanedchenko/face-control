# Avatar image checker
The main goal of the project is to facilitate avatar image checking to remind people to set real picture. It's important for instance for corporate environment to let new people to get on board quickly.

_Inspired by https://github.com/zikes/chrisify_ :)

## Current functionality
Commandline util checking if the image contains any faces.

### Installation
* Install the OpenCV Developer package. On Ubuntu systems that's `sudo apt install libopencv-dev`
* `go get github.com/oanedchenko/face-control`
* `cd $GOPATH/src/github.com/oanedchenko/face-control && go build`

### Usage
`./face-control < path/to/image.jpg`

The output will be like that:
`Face 0 detected: (199,235)-(648,684)`

## Plans
I'm going to create a slack bot to check if people use their real picture as an avatar image and to annoy them with remainders to do so.
