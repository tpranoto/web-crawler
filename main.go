package main

func main() {
	// url := "http://tour.golang.org/welcome/1"
	// fmt.Printf("HTML code of %s ...\n", url)
	// resp, err := http.Get(url)
	// // handle the error if there is one
	// if err != nil {
	// 	panic(err)
	// }
	// // do this now so it won't be forgotten
	// defer resp.Body.Close()
	// // reads html as a slice of bytes
	// html, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// // show the HTML code as a string %s
	// fmt.Printf("%s\n", html)
}

/*

<!doctype html>
<html lang="en" ng-app="tour">

<head>
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-11222381-5"></script>
<script>
window.dataLayer = window.dataLayer || [];
function gtag(){dataLayer.push(arguments);}
gtag("js", new Date());
gtag("config", "UA-11222381-5");
gtag("config", "UA-49880327-6");
</script>
    <meta charset="utf-8">
    <title>A Tour of Go</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="mobile-web-app-capable" content="yes">
    <link rel="shortcut icon" sizes="196x196" href="/favicon.ico">
    <link rel="stylesheet" href="/static/css/app.css" />
    <link rel="stylesheet" href="/static/lib/codemirror/lib/codemirror.css">
    <link href='//fonts.googleapis.com/css?family=Inconsolata' rel='stylesheet' type='text/css'>
</head>

<body>
    <div class="bar top-bar">
        <a class="left logo" href="/list">A Tour of Go</a>
        <div table-of-contents-button=".toc"></div>
        <div feedback-button></div>
    </div>

    <div table-of-contents></div>

    <div ng-view ng-cloak class="ng-cloak"></div>

    <script src="/script.js"></script>
    <script>
    window.transport = HTTPTransport();
    window.socketAddr = "";

    function highlight(selector) {
        var speed = 50;
        var obj = $(selector).stop(true, true)
        for (var i = 0; i < 5; i++) {
            obj.addClass("highlight", speed)
            obj.delay(speed)
            obj.removeClass("highlight", speed)
        }
    }

    function highlightAndClick(selector) {
        highlight(selector);
        setTimeout(function() {
            $(selector)[0].click()
        }, 750);
    }

    function click(selector) {
        $(selector)[0].click();
    }
    </script>
</body>

</html>

*/
