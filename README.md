# schedkit - OCI-packaged schedulers

This repository contains all the packaged schedulers to be used with [sked](https://github.com/schedkit/sked) or [schedctl](https://github.com/schedkit/schedctl).

## Architecture
Every directory name contains a Dockerfile that acts as a build manifest.

The build pipeline takes care of separately building each and every scheduler at every push on `main`, or in a timely fashion.
