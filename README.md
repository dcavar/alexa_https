# Alexa HTTPS Server in Go

(C) 2017 by [Damir Cavar], [Rashmi Bidanta], [Prateek Srivastava]

This is an example implementation of an
<a href="https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-a-web-service.html">HTTPS Amazon Alexa Skill</a>
interface (or
<a href="https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-a-web-service.html">Custom Skill</a>)
to process
<a href="https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html">Alexa JSON requests</a>
and generate a response programmed in <a href="https://golang.org/">Go</a>.


## The HTTPS Server

In the *main* function you will find the following lines that basically fire up the HTTPS server:

    http.HandleFunc("/", MyServer)
	err := http.ListenAndServeTLS(":443", "certificate.pem", "privatekey.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

You will have to make sure that you either use existing keys (*certificate.pem* and *privatekey.pem*) or create
self-signed keys by following for example the instructions on the
[Amazon Alexa *Test a Custom Skill*](https://developer.amazon.com/docs/custom-skills/test-a-custom-skill.html#create-a-private-key-and-self-signed-certificate-for-testing)
page.

The implementation of

    func MyServer(w http.ResponseWriter, req *http.Request) {
    ...
    }

contains code to process the HTTP header:

	var header map[string]string
	header = make(map[string]string)
	for _, element := range []string{"Content-Type",
	                                 "Host",
	                                 "From",
	                                 "Content-Language",
	                                 "Content-Encoding",
	                                 "Server",
	                                 "Date",
	                                 "User-Agent"} {
		v := req.Header.Get(element)
		if v != "" {
			header[element] = v
		}
	}

You can extend the header elements by adding your own keywords for HTTP headers. See the
[Wikipedia page on HTTP-header fields](https://en.wikipedia.org/wiki/List_of_HTTP_header_fields)
for more details.

You can read the content of the HTTP-message body using the following code:

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		err = req.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// Process the body of the message here
	}

Assuming that you set up all necessary structs in *AlexaRequest.go*, you can *unmarshal* the JSON data transmitted from
the Amazon Alexa server using the following code:

	jdata := &AlexaRequest{
		Session: &Session{},
		Context: &Context{},
		Request: &Request{},
	}
	err := json.Unmarshal(b, jdata)

We generate a response to the Amazon Alexa server using the following code:

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"version": "1.0", "response": {"outputSpeech": {"type": "PlainText", "text": "This is an example response. Can I help you with anything else?"}, "shouldEndSession": true}}`))

You can now replace the *w.Write* function call with the message that you generate using your particular response
functions.




## The JSON Parser

Via the Alexa HTTPS Custom Client interface you receive a JSON Request object. In the code repo you will find a sample
*AlexaRequest.json*. See above one possible way to process JSON data objects in Go. There are of course many more
alternative ways.



## The XML-RPC Interface

Our implementation is part of a broader system of Microservices that communicate using
<a href="https://en.wikipedia.org/wiki/XML-RPC">XML-RPC</a> or
<a href="https://en.wikipedia.org/wiki/JSON-RPC">JSON-RPC</a>. We create a very simple
<a href="https://en.wikipedia.org/wiki/XML-RPC">XML-RPC</a> request establishing an
<a href="https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol">HTTP</a>-connection to our main Dispatcher module
in an <a href="https://en.wikipedia.org/wiki/Natural_language_processing">NLP</a> pipeline.

Whatever the utterance is that Alexa sends to our server, we extract it and send it to the *callDispatcher* function.

*callDispatcher* transfers the utterance to an HTTP-based XML-RPC server.

Instead of recreating some XML-RPC package for Go, I took the simple solution to generate a simplified XML response and
communicate it to the XML-RPC server using an

	var buffer bytes.Buffer
	buffer.WriteString(`<?xml version="1.0"?><methodCall><methodName>parse</methodName><params><param><value><string>`)
	buffer.WriteString(text)
	buffer.WriteString(`</string></value></param></params></methodCall>`)

We create an HTTP-client:

	client := &http.Client{}

We establish a new request for this client using the HTTP POST method:

	host := "localhost"
	port := "1234"
	req, err := http.NewRequest(http.MethodPost, "http://"+host+":"+port, strings.NewReader(buffer.String()))
	if err != nil {
		log.Fatal(err)
	}

We create an HTTP header:

	req.Header.Add("Content-Type", "text/xml")
	var contLength = string(buffer.Len())
	req.Header.Add("Content-Length", contLength)

And, finally we submit the entire request to the XML-RPC server:

	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}




[Damir Cavar]: http://damir.cavar.me/ "Damir Cavar homepage"
[Rashmi Bidanta]: https://github.com/rbidanta "Rashmi Bidanta GitHub page"
[Prateek Srivastava]: https://github.com/prateek22sri "Prateek Srivastava GitHub page"
