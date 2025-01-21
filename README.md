# BG_Img_Server

This repo has code for the server of Background_Image_changer
It's written using [Golang](https://go.dev/dl/).


### Prerequisites
- [client](https://github.com/praveenmahasena/bgImgClient) client repo
- [Golang](https://go.dev/dl/) programming language

### Installation

```
git clone git@github.com:praveenmahasena/bgImgServer.git
cd bgImgServer
make
```

This app sends image binary to client side via TCP connection and by receving the image bytes client side writes it into a file and then rendors it as background image using [feh](https://feh.finalrewind.org/).
