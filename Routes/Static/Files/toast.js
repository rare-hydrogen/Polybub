function PopToast(message, classes) {
  const fragment = document.createDocumentFragment();
  const notif = fragment.appendChild(document.createElement("div"));
  const p = notif.appendChild(document.createElement("p"));
  const button = notif.appendChild(document.createElement("button"));

  notif.setAttribute("class", classes);
  p.textContent = message;
  button.setAttribute("class", "delete");
  button.setAttribute("onclick", "this.parentElement.remove();");

  document.body.appendChild(notif);

  setTimeout(() => {
    notif.remove();
  }, 3600);
}

// If a field throws an error, highlight the field
function HighlightField(isError, name) {
  const errFields = document.getElementsByName(name);
  if (!errFields) {
    return;
  }
  errFields.forEach((element) => {
    if (isError) {
      element.classList.add("is-danger");
    } else {
      element.classList.remove("is-danger");
    }
  });
}

// Wipe all field highlights
function WipeHighlightFields() {
  const allInputs = document.getElementsByTagName("input");

  for (let i = 0; i < allInputs.length; i++) {
    allInputs[i].classList.remove("is-danger");
  }
}

// For each error, highlight the matching field
function SetErrorFields(message) {
  const regex = /\b([A-Za-z]+)(?=\s*:)/gm;

  for (const m of message.matchAll(regex)) {
    HighlightField(true, m[1]);
  }
}

// Handle each type of response
function ResponseToHighlightAndNotif(event) {
  // Only toast for json responses
  const contentType = event.detail.xhr.getResponseHeader("Content-Type");
  if (contentType && contentType.includes("text/html")) {
    return;
  }

  if (event.detail.xhr.responseText[0] == "<") {
    return;
  }
  var verb = event.detail.requestConfig.verb;
  var status = event.detail.xhr.status;
  var message = JSON.parse(event.detail.xhr.responseText).message;

  // 100
  if (status < 200) {
    return;
  }
  // 200
  if (status < 300) {
    if (verb === "get") {
      return;
    }
    WipeHighlightFields();
    PopToast("Success!", "toast notification is-primary");
    return;
  }

  // 300
  if (status < 400) {
    return;
  }

  // 400
  if (status < 500) {
    if (verb === "get") {
      return;
    }
    //const name = message.split(":")[0];
    SetErrorFields(message);
    PopToast(message, "toast notification is-warning");
    return;
  }

  // 500
  if (status < 600) {
    if (verb === "get") {
      return;
    }
    PopToast(message, "toast notification is-danger");
  }
  return;
}

window.addEventListener("DOMContentLoaded", (event) => {
  // Handle htmx responses
  document.body.addEventListener("htmx:afterRequest", (event) => {
    ResponseToHighlightAndNotif(event);
  });

  // On field change, unhighlight that field
  document.body.addEventListener("input", (event) => {
    const name = event.target.name;
    HighlightField(false, name);
  });
});
