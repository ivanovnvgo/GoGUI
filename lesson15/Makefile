# If PREFIX isn't provided, we check for $(DESTDIR)/usr/local and use that if it exists.
# Otherwice we fall back to using /usr.

LOCAL != test -d $(DESTDIR)/usr/local && echo -n "/local" || echo -n ""
LOCAL ?= $(shell test -d $(DESTDIR)/usr/local && echo "/local" || echo "")
PREFIX ?= /usr$(LOCAL)

default:
	# Run "sudo make install" to install the application.
	# Run "sudo make uninstall" to uninstall the application.

install:
	install -Dm00644 usr/local/share/applications/lesson15.desktop $(DESTDIR)$(PREFIX)/share/applications/lesson15.desktop
	install -Dm00755 usr/local/bin/lesson15 $(DESTDIR)$(PREFIX)/bin/lesson15
	install -Dm00644 usr/local/share/pixmaps/lesson15.png $(DESTDIR)$(PREFIX)/share/pixmaps/lesson15.png
uninstall:
	-rm $(DESTDIR)$(PREFIX)/share/applications/lesson15.desktop
	-rm $(DESTDIR)$(PREFIX)/bin/lesson15
	-rm $(DESTDIR)$(PREFIX)/share/pixmaps/lesson15.png