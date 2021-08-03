[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

# apkreport

Generate a csv report of basic app information from the 
[MobSF](https://github.com/MobSF/Mobile-Security-Framework-MobSF) API.

Sample run

```bash
./apkreport
```

You will have to edit the generated config file (with your API info): `~/.apkreport.yml`
and run again.

Sample output columns:

```csv
package_name, filename, md5, play_prog_title, play_version, play_url, size, play_installs, play_dev	play_dev_site, play_dev_address
```

## Install

In order to install simply cd to this directory and do:

```bash
make
```

If you want to move the binary to `/usr/bin/local`:

```bash
make install
```