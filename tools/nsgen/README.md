nsgen
=====

This program takes Turtle RDF declarations and turn them into code generated entities, by converting them into JSON-LD and using our own LD machinery.

If you're hacking on this, be prepared to do a lot of `git checkout ../../ld/ns` whenever you generate a blurb of invalid code and the code generator no longer compiles.

You'll find that some RDF out on the internet is written using horrible things like XML Schema and RDFa (Turtle embedded inside HTML, the horrific hybrid nobody asked for). For XML Schema you're basically fucked, but for RDFa, try this: https://www.w3.org/2012/pyRdfa
