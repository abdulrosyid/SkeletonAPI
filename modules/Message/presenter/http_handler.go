package presenter

import (
	"SkeletonAPI/modules/Message/usecase"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

type HTTPMessageHandler struct {
	MessageUseCase usecase.MessageUseCase
}

func NewHTTPHandler(message usecase.MessageUseCase) *HTTPMessageHandler {
	return &HTTPMessageHandler{MessageUseCase: message}
}

func (h *HTTPMessageHandler) Mount(group *echo.Group) {
	group.POST("/message/add", h.AddMessage)
	group.GET("/message/get", h.GetMessage)
	group.GET("/message/client", h.Client)
}

func (h *HTTPMessageHandler) AddMessage(c echo.Context) error {
	message := c.QueryParam("message")
	result := h.MessageUseCase.AddMessage(message)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.Error)
	}

	return c.JSON(http.StatusOK, "add message success")
}

func (h *HTTPMessageHandler) GetMessage(c echo.Context) error {
	result := h.MessageUseCase.GetMessage()
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	return c.JSON(http.StatusOK, result.Result)
}

func (h *HTTPMessageHandler) Client(c echo.Context) error {
	err := homeTemplate.Execute(c.Response(), "ws://localhost:8080/ws")
	if err!= nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return nil
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
    };

	ws = new WebSocket("{{.}}");

	ws.onopen = function() {
      console.log('Connected')
    }

   ws.onmessage = function(evt) {
      var out = document.getElementById('output');
      out.innerHTML += 'Response : ' + evt.data + '<br>';
    }

	ws.onmessage = function(evt) {
      var out = document.getElementById('output');
	  var message = 'Response : ' + evt.data
      out.innerHTML += message + '<br>';
    }

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("Send: " + input.value);
        ws.send(input.value);
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<form>
<p><input id="input" type="text" value="Testing">
<button id="send">Send Message</button>
<p id="output"></p>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
