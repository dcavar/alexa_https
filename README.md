# Alexa HTTPS Server in Go

(C) 2017 by [Damir Cavar], Rashmi Bidanta, Prateek Srivastava

This is an example implementation of an <a href="https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-a-web-service.html">HTTPS Alexa Skill</a> interface (or <a href="https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-a-web-service.html">Custom Skill</a>) to process <a href="https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html">Alexa JSON requests</a> and generate
a response programmed in <a href="https://golang.org/">Go</a>.


## The HTTPS Server



## The JSON Parser




## The XML-RPC Interface

Our implementation is part of a broader system of Microservices that communicate using
<a href="https://en.wikipedia.org/wiki/XML-RPC">XML-RPC</a> or
<a href="https://en.wikipedia.org/wiki/JSON-RPC">JSON-RPC</a>. We create a very simple
<a href="https://en.wikipedia.org/wiki/XML-RPC">XML-RPC</a> request establishing an
<a href="https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol">HTTP</a>-connection to our main Dispatcher module
in an <a href="https://en.wikipedia.org/wiki/Natural_language_processing">NLP</a> pipeline.




[Damir Cavar]: http://damir.cavar.me/ "Damir Cavar homepage"

