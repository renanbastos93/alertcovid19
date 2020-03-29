![Go Build](https://github.com/renanbastos93/alertcovid19/workflows/Go%20Build/badge.svg)
![Gosec](https://github.com/renanbastos93/alertcovid19/workflows/Gosec/badge.svg)
![Go Test](https://github.com/renanbastos93/alertcovid19/workflows/Go%20Test/badge.svg)

## Alert COVID-19
This repository has the goal to create software to many platforms to run in one thread to validate each hour if we have new cases confirmed, deaths or recovered by COVID-19 in Brazil. I wish to improve this code to receive by params other countries or all world to show the notification.

We are currently using this [COVID19-BRAZIL-API](https://covid19-brazil-api-docs.now.sh/) API to query COVID-19 quantity data


### Usage
```bash
$ make <macos|linux|windows>  # if it's the window you can to run "make zip"
$ ./alertcovid19 -timer=10    # 10 seconds
$ make clean
```


### Do you wish to contribute?
Open a pull request or issue for we discuss.
