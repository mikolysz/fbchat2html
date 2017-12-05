# fbchat2html
fbchat2html is a little utility written in golang to convert chat messages from the facebook archive to readable html documents. It expects a json file generated with the help of another utility,  [facebook chat archive parser](https://github.com/ownaginatious/fbchat-archive-parser) as input.
It makes a html file for each conversation it finds and saves it's contents in there. The messages are sorted in chronological order, in contrast to the reverse-chronological order found in the original archive and are split by year, month and day.
Years are headings of level two, months are headings of level three, days are headings of level four and message authors are headings of level five. This helps people using screen-reading software  such as myself navigate the document  more quickly and easily.
## Installation:
1. Install [golang](http://golang.org).
2. Execute the following command:
    go get github.com/miki123211/fbchat2html
3. Install [facebook chat archive parser](https://github.com/ownaginatious/fbchat-archive-parser) as written in it's readme.
4. You're ready to go!
Binaries for windows, linux and mac coming soon.
## Quick start:
1. Generate a facebook archive, you can find instructions on how to do that on facebook's website.
2. Use [facebook chat archive parser](https://github.com/ownaginatious/fbchat-archive-parser) to generate a json file with your message history, How to do that is described in it's documentation.
3. Open your command prompt, go to the directory containing your json file and execute:
    fbchat2html your_json_file.json
By default, it will create a directory named "output" in your current working directory and place your files there. You can change that by setting the --output parameter on the command line.
It will also create a stats.txt file inside that directory containign some useful statistics.
##  Known issues and limitations:
The readme is not  completely done yet, particularly the section on generating the facebook archive.
The documentation and commends probably need reviewing by someone with english as their first language or, at least, someone more proficient in english than me.
It's impossible to change the name of the stats file.
The process of writing the archive probably could be more parallelized to make it a bit faster.
The code definitely needs some refactoring, it's pretty hacky at the time but it works!
Binaries for various platforms need to be generated, preferrably bundled with fbchat_archive_parser. License issues need to be investigated if I decide to do that.
##  Contributing and contact:
Feel free to contribute, submit issues or pull requests, or, if you have the need, contact me privately via email at miki one hundret twenty three thousant two hundret and eleven at gmail dot charlie echo mike.
Also feel free to reach out via [facebook](https://www.facebook.com/profile.php?id=10000809608526) and [twitter](http://twitter.com/miki123211).
## License:
This code is in public domain, do whatever you want with it without worrying about licenses.
