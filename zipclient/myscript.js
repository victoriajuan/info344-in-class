var nameString = document.getElementById('nameText').value

function changeURL() {
    var stateObj = { foo: "bar" };
    history.pushState(stateObj, "page 2", "hello?name" + nameString);
}

var link = document.getElementById('click');
link.addEventListener('click', changeURL, false);