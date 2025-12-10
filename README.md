# bvdeenen-svlogj
Frontend for svlogtail in Void Linux

## the socklog group

The tool is meant to be run as your regular user, so no _sudo_. If you're not yet member of `socklog` you can execute these steps.

    sudo usermod -aG socklog $USER
    newgrp socklog

The `newgrp socklog` is only for necessary as you long as you have not logged out and back in again.
