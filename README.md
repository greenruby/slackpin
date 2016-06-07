Slackpin
===========

I just want to read all the links pinned in the #links channel.

    # list pins 
    go run pins.go

    # clean pins
    go run clean_pins.go

Used by the [GreenRuby][1] publication workflow:

- during the week people submit links
- links that are not redundant or qualify are pinned
- at publication step the editor grabs all the links and include them in the newsletter
- then he unpins everything to make it clean for a new week


[1]: http://greenruby.org
