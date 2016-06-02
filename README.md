# notifier
A library for sending native desktop notifications from Go. This uses
platform-specific helper libraries as follows:

* OSX: [Terminal Notifier](https://github.com/julienXX/terminal-notifier)
* Windows: [notifu](https://www.paralint.com/projects/notifu/)

Those libraries are embedded directly in Go in this library using [go-bindata](https://github.com/jteeuwen/go-bindata), so there are no external dependencies or expectations for installations on the user's system. These libraries were also chosen for their small size, particularly in the case of notifu, which is far smaller than things like Growl or Snarl.

To generate updated embedded binaries, you can run, for example:

```
go get -u github.com/jteeuwen/go-bindata/...
cd platform/terminal-notifier-1.6.3
go-bindata terminal-notifier.app/...
mv bindata.go ../../osx
```

You then have to manually change the package in `bindata.go` to `osx` instead of `main` in that
case.

This is currently a work in progress and only runs on OSX and Windows and embeds
all binaries for all platforms instead of dynamically only including the
required platform at build time, for example.

The platform directory is only here to serve as a reference for the native binaries
being used.

For documentation purposes, here are the raw options for terminal-notifier:

```
$ ./terminal-notifier.app/Contents/MacOS/terminal-notifier
terminal-notifier (1.6.3) is a command-line tool to send OS X User Notifications.

Usage: terminal-notifier -[message|list|remove] [VALUE|ID|ID] [options]

   Either of these is required (unless message data is piped to the tool):

       -help              Display this help banner.
       -message VALUE     The notification message.
       -remove ID         Removes a notification with the specified ‘group’ ID.
       -list ID           If the specified ‘group’ ID exists show when it was delivered,
                          or use ‘ALL’ as ID to see all notifications.
                          The output is a tab-separated list.

   Optional:

       -title VALUE       The notification title. Defaults to ‘Terminal’.
       -subtitle VALUE    The notification subtitle.
       -sound NAME        The name of a sound to play when the notification appears. The names are listed
                          in Sound Preferences. Use 'default' for the default notification sound.
       -group ID          A string which identifies the group the notifications belong to.
                          Old notifications with the same ID will be removed.
       -activate ID       The bundle identifier of the application to activate when the user clicks the notification.
       -sender ID         The bundle identifier of the application that should be shown as the sender, including its icon.
       -appIcon URL       The URL of a image to display instead of the application icon (Mavericks+ only)
       -contentImage URL  The URL of a image to display attached to the notification (Mavericks+ only)
       -open URL          The URL of a resource to open when the user clicks the notification.
       -execute COMMAND   A shell command to perform when the user clicks the notification.

When the user activates a notification, the results are logged to the system logs.
Use Console.app to view these logs.

Note that in some circumstances the first character of a message has to be escaped in order to be recognized.
An example of this is when using an open bracket, which has to be escaped like so: ‘\[’.
```
