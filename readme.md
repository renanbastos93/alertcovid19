# Alert COVID-19
![Go Build](https://github.com/renanbastos93/alertcovid19/workflows/Go%20Build/badge.svg)
![Gosec](https://github.com/renanbastos93/alertcovid19/workflows/Gosec/badge.svg)
![Go Test](https://github.com/renanbastos93/alertcovid19/workflows/Go%20Test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/renanbastos93/alertcovid19)](https://goreportcard.com/report/github.com/renanbastos93/alertcovid19)

Alert COVID-19 is a small multiplatform tool written in Golang to help keep you informed about the current situation of COVID-19 in your region,
while you stay safe at home.

At every hour, it fetches data from an API and displays a native notification on your system if there's any updates
on the number of confirmed cases, deaths or recovered patients.

## Getting started

If you just want to get up and running, follow these steps:

1. Go to the [release page](https://github.com/renanbastos93/alertcovid19/releases)
2. Download the lastest release for your platform (MacOS, Linux or Windows)
3. Extract the contents and run the executable file

### Building the project

If you want to build the project yourself, here's how you can use our Makefile:

```bash
# Build the project
$ make <macos|linux|windows>
# If you're on Windows, you can also run "make zip" to create a zipped executable

# Run the program, updating every 10s and if there's a changes, display a notification
# Defaults to 1h00m00s (1 hour)
$ ./alertcovid19 -t 10s

# If you want to remove the built files
$ make clean

# Or simply...
$ go run .
```
Dependencies for building:
- Go
- Golint
- Gosec


## 👍 Contributing
If you want to say **thank you** and/or support the active development of `AlertCovid19`:

1. Add a [GitHub Star](https://github.com/renanbastos93/alertcovid19) to the project.
2. Tweet about the project [on your Twitter](https://twitter.com/intent/tweet?text=%F0%9F%9A%80%20Alert%20COVID-19%20%E2%80%94%20was%20made%20in%20Golang%20to%20show%20push%20notification%20in%20your%20operating%20system%20with%20updates%20based%20on%20your%20geolocation).
3. Write a review or tutorial on [Medium](https://medium.com/), [Dev.to](https://dev.to/) or personal blog.
4. Help us to translate this `README` to another language.

## Do you wish to contribute?
Open a pull request or issue for we discuss.

## Coffee Supporters
<a href="https://www.buymeacoffee.com/renanbastos93" target="_blank">
  <img src="https://images-na.ssl-images-amazon.com/images/I/41LnWYwUe4L._SX331_BO1,204,203,200_.jpg" alt="Buy Me A Coffee" height="100" > Buy My Coffee
</a>

## Code Contributors
| [<img src="https://avatars1.githubusercontent.com/u/16732610?s=460&v=4" width="115"><br><sub>@wgrr</sub>](https://github.com/wgrr) | [<img src="https://avatars1.githubusercontent.com/u/14180225?s=460&v=4" width="115"><br><sub>@u5surf</sub>](https://github.com/u5surf) | [<img src="https://avatars0.githubusercontent.com/u/20388082?s=460&u=43d4f0f9f66f40170e10ddeb23b5cca41b5afd81&v=4" width="115"><br><sub>@fsmiamoto</sub>](https://github.com/fsmiamoto) | [<img src="https://avatars3.githubusercontent.com/u/2396581?s=460&u=2c624fe4d878b0a25589be49dc47dd9a3e6c0e43&v=4" width="115"><br><sub>@cyruzin</sub>](https://github.com/cyruzin) | [<img src="https://avatars0.githubusercontent.com/u/8202898?s=460&u=668363f7f686077bea518133e28b77d11fd4c242&v=4" width="115"><br><sub>@renanbastos93</sub>](https://github.com/renanbastos93) |
| :---: |  :---: |  :---: |  :---: |  :---: |
