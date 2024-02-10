@echo off

cd ../web/build
ren _app app
fart -q index.html _app app
