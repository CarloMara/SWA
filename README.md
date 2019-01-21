# Simple Web Application

This is a simple application to learn a bit about current development trends in web application development. It's by no means bleeding edge or anything fancy, just a simple playground to test things.


# Tech stack

This is really something that I'm figuring out as I go. Currently it's golang+revel for the back end and npm+bootstrap+webpack for the front end.

# How to start

Clone the repo inside GOPATH. Install revel, node and webpack-cli

Launch with _revel run_, or _revel run -p 8080_ to bind the server to a known port. Issue this command from the root of the repo.

To rebuild static assets and js use _webpack-cli_ from the webpack folder. Remember that _node_modules_ is not committed to the repo, so to deploy you need call _npm rebuild_